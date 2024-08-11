package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	configPath = "/etc/sysconfig/gobsips"
)

type Config struct {
	Username   string
	Password   string
	ListenHost string
	ListenPort string
}

func loadConfig() Config {
	config := Config{}
	if data, err := os.ReadFile(configPath); err == nil {
		scanner := bufio.NewScanner(strings.NewReader(string(data)))
		for scanner.Scan() {
			parts := strings.SplitN(scanner.Text(), "=", 2)
			if len(parts) == 2 {
				key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
				switch key {
				case "USERNAME":
					config.Username = value
				case "PASSWORD":
					config.Password = value
				case "LISTEN_HOST":
					config.ListenHost = value
				case "LISTEN_PORT":
					config.ListenPort = value
				}
			}
		}
	}
	return config
}

func saveConfig(config Config) error {
	content := fmt.Sprintf(`USERNAME=%s
PASSWORD=%s
LISTEN_HOST=%s
LISTEN_PORT=%s
`, config.Username, config.Password, config.ListenHost, config.ListenPort)

	return os.WriteFile(configPath, []byte(content), 0644)
}
