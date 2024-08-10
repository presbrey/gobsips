package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	serviceName = "gobsips.service"
)

func installDefaultConfig() error {
	config := Config{}
	config.Username = "api"
	config.Password = md5sum(getMachineID())
	config.ListenHost = "0.0.0.0"
	config.ListenPort = "1080"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return saveConfig(config)
	}
	return nil
}

func installSystemdService() error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %v", err)
	}

	serviceContent := fmt.Sprintf(`[Unit]
Description=Go Binary Systemd-Integrated Proxy Server
After=network.target

[Service]
ExecStart=%s -daemon
Restart=always
User=nobody

[Install]
WantedBy=multi-user.target
`, exePath)

	if err := os.WriteFile("/etc/systemd/system/"+serviceName, []byte(serviceContent), 0644); err != nil {
		return err
	}

	// Enable and start the service
	if err := runCommand("systemctl", "daemon-reload"); err != nil {
		return err
	}
	if err := runCommand("systemctl", "enable", serviceName); err != nil {
		return err
	}
	if err := runCommand("systemctl", "start", serviceName); err != nil {
		return err
	}

	return nil
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getMachineID() string {
	if data, err := os.ReadFile("/etc/machine-id"); err == nil {
		return strings.TrimSpace(string(data))
	}
	return ""
}

func md5sum(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}
