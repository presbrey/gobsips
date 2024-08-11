package main

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary config file
	tempFile, err := os.CreateTemp("", "test_config")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write test config
	testConfig := `USERNAME=testuser
PASSWORD=testpass
LISTEN_HOST=127.0.0.1
LISTEN_PORT=8080
`
	if _, err := tempFile.Write([]byte(testConfig)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// Set the configPath to our temp file
	oldConfigPath := configPath
	configPath = tempFile.Name()
	defer func() { configPath = oldConfigPath }()

	// Test loadConfig
	config := loadConfig()

	if config.Username != "testuser" {
		t.Errorf("Expected Username 'testuser', got '%s'", config.Username)
	}
	if config.Password != "testpass" {
		t.Errorf("Expected Password 'testpass', got '%s'", config.Password)
	}
	if config.ListenHost != "127.0.0.1" {
		t.Errorf("Expected ListenHost '127.0.0.1', got '%s'", config.ListenHost)
	}
	if config.ListenPort != "8080" {
		t.Errorf("Expected ListenPort '8080', got '%s'", config.ListenPort)
	}
}

func TestSaveConfig(t *testing.T) {
	// Create a temporary file for the config
	tempFile, err := os.CreateTemp("", "test_config")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	// Set the configPath to our temp file
	oldConfigPath := configPath
	configPath = tempFile.Name()
	defer func() { configPath = oldConfigPath }()

	// Test saveConfig
	testConfig := Config{
		Username:   "testuser",
		Password:   "testpass",
		ListenHost: "127.0.0.1",
		ListenPort: "8080",
	}

	if err := saveConfig(testConfig); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// Read the saved config
	data, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read saved config: %v", err)
	}

	expectedContent := `USERNAME=testuser
PASSWORD=testpass
LISTEN_HOST=127.0.0.1
LISTEN_PORT=8080
`
	if string(data) != expectedContent {
		t.Errorf("Saved config does not match expected content.\nExpected:\n%s\nGot:\n%s", expectedContent, string(data))
	}
}
