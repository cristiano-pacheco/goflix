package integration

import (
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestHealthcheck(t *testing.T) {
	// Get port from environment or use default
	testPort := os.Getenv("HTTP_PORT")
	if testPort == "" {
		testPort = "9000"
	}

	// Create and start the bootstrap
	bootstrap := NewBootstrap(testPort)

	// Start the application
	if err := bootstrap.Start(); err != nil {
		t.Fatalf("Failed to start application: %v", err)
	}

	// Ensure cleanup
	defer func() {
		if err := bootstrap.Stop(); err != nil {
			t.Logf("Error stopping application: %v", err)
		}
	}()

	// Create HTTP client
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Test the healthcheck endpoint
	t.Run("GET /healthcheck returns 200 OK", func(t *testing.T) {
		url := bootstrap.GetBaseURL() + "/healthcheck"

		resp, err := client.Get(url)
		if err != nil {
			t.Fatalf("Failed to make request to %s: %v", url, err)
		}
		defer resp.Body.Close()

		// Check status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		// Check response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		expectedBody := "OK"
		if string(body) != expectedBody {
			t.Errorf("Expected response body %q, got %q", expectedBody, string(body))
		}

		// Check content type (optional)
		contentType := resp.Header.Get("Content-Type")
		if contentType != "" && contentType != "text/plain; charset=utf-8" {
			t.Logf("Content-Type: %s", contentType)
		}
	})

	t.Run("Multiple requests to /healthcheck should work", func(t *testing.T) {
		url := bootstrap.GetBaseURL() + "/healthcheck"

		// Make multiple requests to ensure the endpoint is stable
		for i := 0; i < 5; i++ {
			resp, err := client.Get(url)
			if err != nil {
				t.Fatalf("Request %d failed: %v", i+1, err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Request %d: Expected status code %d, got %d", i+1, http.StatusOK, resp.StatusCode)
			}

			resp.Body.Close()
		}
	})

	t.Run("Invalid method to /healthcheck should return 405", func(t *testing.T) {
		url := bootstrap.GetBaseURL() + "/healthcheck"

		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			t.Fatalf("Failed to create POST request: %v", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to make POST request: %v", err)
		}
		defer resp.Body.Close()

		// POST requests should return 405 Method Not Allowed
		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d for POST request, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
		}
	})
}
