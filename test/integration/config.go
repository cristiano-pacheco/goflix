package integration

import (
	"os"
)

// SetupTestEnvironment sets up the necessary environment variables for integration tests
func SetupTestEnvironment() {
	// Set test environment variables
	testEnvVars := map[string]string{
		"ENVIRONMENT":             "test",
		"HTTP_PORT":               "9000",
		"DB_HOST":                 "localhost",
		"DB_PORT":                 "5432",
		"DB_NAME":                 "goflix",
		"DB_USER":                 "postgres",
		"DB_PASSWORD":             "postgres",
		"DB_MAX_OPEN_CONNECTIONS": "10",
		"DB_MAX_IDLE_CONNECTIONS": "5",
		"DB_SSL_MODE":             "false",
		"DB_PREPARE_STMT":         "true",
		"DB_ENABLE_LOGS":          "false",
		"CORS_ALLOWED_ORIGINS":    "*",
		"CORS_ALLOWED_METHODS":    "GET,POST,PUT,DELETE,OPTIONS",
		"CORS_ALLOWED_HEADERS":    "*",
		"CORS_EXPOSED_HEADERS":    "",
		"CORS_ALLOW_CREDENTIALS":  "false",
		"CORS_MAX_AGE":            "3600",
		"JWT_SECRET":              "test-secret-key-for-integration-tests",
		"JWT_EXPIRATION_HOURS":    "24",
		"MAIL_HOST":               "",
		"MAIL_PORT":               "587",
		"MAIL_USERNAME":           "",
		"MAIL_PASSWORD":           "",
		"MAIL_FROM":               "test@example.com",
		"TELEMETRY_ENABLED":       "false",
		"TELEMETRY_ENDPOINT":      "",
		"TELEMETRY_SERVICE_NAME":  "goflix-test",
		"APP_NAME":                "goflix-test",
		"APP_VERSION":             "test",
		"LOG_LEVEL":               "info",
		"LOG_FORMAT":              "json",
		"RABBITMQ_URL":            "",
		"RABBITMQ_ENABLED":        "false",
	}

	for key, value := range testEnvVars {
		if err := os.Setenv(key, value); err != nil {
			// Log error but don't fail, as some variables might be read-only
			continue
		}
	}
}

// CleanupTestEnvironment cleans up test environment variables
func CleanupTestEnvironment() {
	testEnvVars := []string{
		"ENVIRONMENT",
		"HTTP_PORT",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"DB_USER",
		"DB_PASSWORD",
		"DB_MAX_OPEN_CONNECTIONS",
		"DB_MAX_IDLE_CONNECTIONS",
		"DB_SSL_MODE",
		"DB_PREPARE_STMT",
		"DB_ENABLE_LOGS",
		"CORS_ALLOWED_ORIGINS",
		"CORS_ALLOWED_METHODS",
		"CORS_ALLOWED_HEADERS",
		"CORS_EXPOSED_HEADERS",
		"CORS_ALLOW_CREDENTIALS",
		"CORS_MAX_AGE",
		"JWT_SECRET",
		"JWT_EXPIRATION_HOURS",
		"MAIL_HOST",
		"MAIL_PORT",
		"MAIL_USERNAME",
		"MAIL_PASSWORD",
		"MAIL_FROM",
		"TELEMETRY_ENABLED",
		"TELEMETRY_ENDPOINT",
		"TELEMETRY_SERVICE_NAME",
		"APP_NAME",
		"APP_VERSION",
		"LOG_LEVEL",
		"LOG_FORMAT",
		"RABBITMQ_URL",
		"RABBITMQ_ENABLED",
	}

	for _, key := range testEnvVars {
		os.Unsetenv(key)
	}
}
