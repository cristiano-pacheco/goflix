# Integration Tests

This directory contains integration tests for the GoFlix application. The tests are designed to run against a real instance of the application, providing end-to-end testing capabilities.

## Overview

The integration test setup includes:

- **Bootstrap**: A completely independent application bootstrap that starts the GoFlix application in a separate process
- **Test Configuration**: Environment configuration specifically for testing
- **Health Check Tests**: Tests that verify the `/healthcheck` endpoint returns a 200 status code

## Architecture

### Bootstrap (`bootstrap.go`)
The bootstrap is responsible for:
- Starting the application in a separate process using `go run`
- Setting up test environment variables
- Waiting for the application to become ready
- Providing graceful shutdown capabilities
- Being completely independent from the test code

### Configuration (`config.go`)
Provides test-specific configuration including:
- HTTP port set to 9000
- Test database settings
- Disabled telemetry and external services
- Minimal logging configuration

### Tests (`healthcheck_test.go`)
Contains integration tests for:
- Basic healthcheck endpoint functionality
- Multiple request handling
- HEAD request support

## Running the Tests

### Method 1: Using Make
```bash
# From the project root
make test-integration
```

### Method 2: Using Go Test Directly
```bash
# From the project root
go test -v ./test/integration/...
```

### Method 3: Using the Standalone Test Runner
```bash
# From the test directory
cd test
go run run_integration_tests.go -v

# Or with custom port
go run run_integration_tests.go -port 8080 -v
```

### Method 4: Manual cURL Testing
If you want to manually test the endpoint:

1. Start the application with the test configuration:
```bash
HTTP_PORT=9000 ENVIRONMENT=test go run main.go
```

2. Test the healthcheck endpoint:
```bash
curl --location 'http://localhost:9000/healthcheck'
```

Expected response:
- Status Code: 200 OK
- Body: "OK"

## Test Configuration

The tests use the following configuration:
- **Port**: 9000 (configurable)
- **Environment**: test
- **Database**: goflix_test (PostgreSQL)
- **Telemetry**: Disabled
- **External Services**: Disabled

## Prerequisites

Before running the integration tests, ensure:

1. **Go**: Go 1.21+ is installed
2. **Database**: PostgreSQL is running (if database tests are added)
3. **Port Availability**: Port 9000 (or specified port) is available
4. **Dependencies**: All Go dependencies are installed (`go mod download`)

## Adding New Integration Tests

To add new integration tests:

1. Create a new test file in this directory (e.g., `api_test.go`)
2. Use the existing bootstrap pattern:
```go
func TestNewEndpoint(t *testing.T) {
    bootstrap := NewBootstrap("9000")
    if err := bootstrap.Start(); err != nil {
        t.Fatalf("Failed to start application: %v", err)
    }
    defer bootstrap.Stop()
    
    // Your test logic here
    client := &http.Client{Timeout: 5 * time.Second}
    resp, err := client.Get(bootstrap.GetBaseURL() + "/your-endpoint")
    // ... assertions
}
```

## Troubleshooting

### Port Already in Use
If you get a "port already in use" error:
```bash
# Find and kill the process using port 9000
lsof -ti:9000 | xargs kill -9
```

### Application Fails to Start
Check the application logs for configuration issues. Ensure all required environment variables are set.

### Tests Timeout
The bootstrap waits up to 30 seconds for the application to become ready. If tests timeout:
- Check if the application is starting correctly
- Verify database connectivity (if required)
- Check for port conflicts

## Design Principles

1. **Independence**: The bootstrap runs the application completely independently from the test code
2. **Isolation**: Each test gets a fresh application instance
3. **Realistic**: Tests run against the real application, not mocks
4. **Fast**: Tests start quickly and clean up properly
5. **Configurable**: Port and other settings can be customized

## Future Enhancements

Potential improvements for the integration test suite:
- Database migration and cleanup
- Docker-based test environment
- Parallel test execution with port management
- API endpoint testing beyond healthcheck
- Performance and load testing capabilities 