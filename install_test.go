package main

import (
	"os"
	"testing"
)

func TestInstallDefaultConfig(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "test_config")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set the configPath to a file in our temp directory
	oldConfigPath := configPath
	configPath = tempDir + "/config"
	defer func() { configPath = oldConfigPath }()

	// Test installDefaultConfig
	if err := installDefaultConfig(); err != nil {
		t.Fatalf("installDefaultConfig failed: %v", err)
	}

	// Check if the config file was created
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Errorf("Config file was not created")
	}

	// Load the created config and check its contents
	config := loadConfig()

	if config.Username != "api" {
		t.Errorf("Expected Username 'api', got '%s'", config.Username)
	}
	if config.ListenHost != "0.0.0.0" {
		t.Errorf("Expected ListenHost '0.0.0.0', got '%s'", config.ListenHost)
	}
	if config.ListenPort != "1080" {
		t.Errorf("Expected ListenPort '1080', got '%s'", config.ListenPort)
	}
	// We can't check the password as it's based on the machine ID
}

func TestGetMachineID(t *testing.T) {
	// Create a temporary file for the machine-id
	tempFile, err := os.CreateTemp("", "test_machine_id")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write a test machine ID
	testMachineID := "test_machine_id_123"
	if _, err := tempFile.Write([]byte(testMachineID)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// Replace the machine-id file path
	oldMachineIDPath := "/etc/machine-id"
	pathMachineID = tempFile.Name()
	defer func() { pathMachineID = oldMachineIDPath }()

	// Test getMachineID
	result := getMachineID()

	if result != testMachineID {
		t.Errorf("Expected machine ID '%s', got '%s'", testMachineID, result)
	}
}

func TestMD5Sum(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"test", "098f6bcd4621d373cade4e832627b4f6"},
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"gobsips", "27c2ff08761fb2321cbae70a74e35b9c"},
	}

	for _, tc := range testCases {
		result := md5sum(tc.input)
		if result != tc.expected {
			t.Errorf("For input '%s', expected MD5 '%s', got '%s'", tc.input, tc.expected, result)
		}
	}
}
