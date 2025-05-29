package integration

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

// findProjectRoot finds the project root by looking for go.mod file
func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up the directory tree looking for go.mod
	for {
		goModPath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached the root directory
			break
		}
		dir = parent
	}

	return "", errors.New("could not find project root (go.mod not found)")
}

// Bootstrap starts the application using 'make run' for integration tests
func Bootstrap(ctx context.Context) (*exec.Cmd, error) {
	// Find the project root by looking for go.mod
	workDir, err := findProjectRoot()
	if err != nil {
		return nil, fmt.Errorf("failed to find project root: %w", err)
	}

	log.Printf("Bootstrap: Starting application in directory: %s", workDir)

	// Check if Makefile exists
	makefilePath := filepath.Join(workDir, "Makefile")
	if _, err := os.Stat(makefilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("makefile not found at %s", makefilePath)
	}

	cmd := exec.CommandContext(ctx, "make", "run")
	cmd.Dir = workDir

	// Ensure output is printed to stdout/stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	log.Printf("Bootstrap: Executing command: make run in %s", workDir)

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start application with 'make run': %w", err)
	}

	log.Printf("Bootstrap: Application started with PID: %d", cmd.Process.Pid)

	// Give the application some time to start
	log.Printf("Bootstrap: Waiting for application to start...")
	time.Sleep(5 * time.Second)

	log.Printf("Bootstrap: Application should be ready")
	return cmd, nil
}

// Shutdown stops the application process and its children
func Shutdown(cmd *exec.Cmd) error {
	if cmd == nil || cmd.Process == nil {
		return nil
	}

	log.Printf("Shutdown: Terminating process with PID: %d", cmd.Process.Pid)

	// Send SIGTERM to the process group to terminate all child processes
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err == nil {
		log.Printf("Shutdown: Sending SIGTERM to process group: %d", pgid)
		syscall.Kill(-pgid, syscall.SIGTERM)
	} else {
		log.Printf("Shutdown: Sending SIGTERM to process: %d", cmd.Process.Pid)
		cmd.Process.Signal(syscall.SIGTERM)
	}

	// Wait for the process to exit with timeout
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(5 * time.Second):
		log.Printf("Shutdown: Graceful shutdown timeout, force killing...")
		// Force kill if graceful shutdown takes too long
		if pgid, err := syscall.Getpgid(cmd.Process.Pid); err == nil {
			syscall.Kill(-pgid, syscall.SIGKILL)
		} else {
			cmd.Process.Kill()
		}
		<-done
	case err := <-done:
		log.Printf("Shutdown: Process terminated")
		return err
	}

	return nil
}
