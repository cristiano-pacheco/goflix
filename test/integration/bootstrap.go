package integration

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

// Bootstrap represents the application bootstrap for integration tests
type Bootstrap struct {
	cmd    *exec.Cmd
	port   string
	cancel context.CancelFunc
}

// NewBootstrap creates a new bootstrap instance
func NewBootstrap(port string) *Bootstrap {
	return &Bootstrap{
		port: port,
	}
}

// Start starts the application in a separate process
func (b *Bootstrap) Start() error {
	// Set up test environment
	SetupTestEnvironment()

	// Set environment variables for the test
	env := os.Environ()
	env = append(env, fmt.Sprintf("HTTP_PORT=%s", b.port))
	env = append(env, "ENVIRONMENT=test")

	// Create context for the command
	ctx, cancel := context.WithCancel(context.Background())
	b.cancel = cancel

	// Start the application
	b.cmd = exec.CommandContext(ctx, "go", "run", "main.go")
	b.cmd.Env = env
	b.cmd.Dir = "../../"

	// Redirect output for debugging if needed
	b.cmd.Stdout = os.Stdout
	b.cmd.Stderr = os.Stderr

	if err := b.cmd.Start(); err != nil {
		return fmt.Errorf("failed to start application: %w", err)
	}

	// Wait for the application to be ready
	if err := b.waitForReady(); err != nil {
		b.Stop()
		return fmt.Errorf("application failed to become ready: %w", err)
	}

	log.Printf("Application started successfully on port %s", b.port)
	return nil
}

// Stop stops the application
func (b *Bootstrap) Stop() error {
	if b.cancel != nil {
		b.cancel()
	}
	
	if b.cmd != nil && b.cmd.Process != nil {
		// Send SIGTERM to gracefully shutdown
		if err := b.cmd.Process.Signal(syscall.SIGTERM); err != nil {
			log.Printf("Failed to send SIGTERM: %v", err)
			// Force kill if graceful shutdown fails
			if killErr := b.cmd.Process.Kill(); killErr != nil {
				log.Printf("Failed to kill process: %v", killErr)
			}
			return err
		}
		
		// Wait for the process to exit
		done := make(chan error, 1)
		go func() {
			done <- b.cmd.Wait()
		}()
		
		select {
		case err := <-done:
			// Ignore "signal: killed" errors as they are expected
			if err != nil && err.Error() != "signal: killed" && err.Error() != "signal: terminated" {
				return err
			}
			return nil
		case <-time.After(5 * time.Second):
			log.Println("Graceful shutdown timeout, force killing")
			if killErr := b.cmd.Process.Kill(); killErr != nil {
				log.Printf("Failed to kill process: %v", killErr)
			}
			// Wait a bit more for the kill to take effect
			select {
			case <-done:
				return nil
			case <-time.After(2 * time.Second):
				return fmt.Errorf("failed to stop application within timeout")
			}
		}
	}
	
	return nil
}

// waitForReady waits for the application to be ready by checking the health endpoint
func (b *Bootstrap) waitForReady() error {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	url := fmt.Sprintf("http://localhost:%s/healthcheck", b.port)

	for i := 0; i < 30; i++ { // Wait up to 30 seconds
		resp, err := client.Get(url)
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				return nil
			}
		}

		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("application did not become ready within 30 seconds")
}

// GetBaseURL returns the base URL for the application
func (b *Bootstrap) GetBaseURL() string {
	return fmt.Sprintf("http://localhost:%s", b.port)
}

// StartWithGracefulShutdown starts the application and sets up graceful shutdown
func (b *Bootstrap) StartWithGracefulShutdown() error {
	if err := b.Start(); err != nil {
		return err
	}

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Received shutdown signal, stopping application...")
		if err := b.Stop(); err != nil {
			log.Printf("Error stopping application: %v", err)
		}
	}()

	return nil
}
