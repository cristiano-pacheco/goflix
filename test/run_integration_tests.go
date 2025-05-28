package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var (
		port    = flag.String("port", "9000", "Port to run the application on")
		verbose = flag.Bool("v", false, "Verbose output")
		help    = flag.Bool("h", false, "Show help")
	)
	flag.Parse()

	if *help {
		fmt.Println("Integration Test Runner")
		fmt.Println("Usage: go run run_integration_tests.go [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		return
	}

	// Get the project root directory (one level up from this file)
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	projectRoot := filepath.Join(currentDir, "..")

	// Change to project root
	if err := os.Chdir(projectRoot); err != nil {
		log.Fatalf("Failed to change to project root: %v", err)
	}

	fmt.Printf("Running integration tests on port %s...\n", *port)

	// Set environment variables
	os.Setenv("HTTP_PORT", *port)
	os.Setenv("ENVIRONMENT", "test")

	// Build the test command
	args := []string{"test"}
	if *verbose {
		args = append(args, "-v")
	}
	args = append(args, "./test/integration/...")

	// Run the tests
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		log.Fatalf("Failed to run tests: %v", err)
	}

	fmt.Println("Integration tests completed successfully!")
}
