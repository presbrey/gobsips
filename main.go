package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/armon/go-socks5"
)

var (
	flagDaemon  = flag.Bool("daemon", false, "Run as a daemon")
	flagInstall = flag.Bool("install", true, "Install the daemon (systemd service)")
)

func main() {
	if err := do(); err != nil {
		log.Fatal(err)
	}
}

func do() error {
	flag.Parse()
	if *flagDaemon {
		*flagInstall = false
	}
	if *flagInstall && getuid() != 0 {
		return fmt.Errorf("you must install this program as root")
	}

	switch {

	case *flagInstall:

		if err := installDefaultConfig(); err != nil {
			return fmt.Errorf("failed to write default config to %s: %v", configPath, err)
		}
		if err := installSystemdService(); err != nil {
			return fmt.Errorf("failed to install systemd service: %v", err)
		}
		fmt.Println("gobsips installed and started as a systemd service.")
		return nil

	case *flagDaemon:

		config := loadConfig()

		// Create SOCKS5 configuration
		conf := &socks5.Config{
			Credentials: socks5.StaticCredentials{
				config.Username: config.Password,
			},
		}

		// Create SOCKS5 server
		server, err := socks5.New(conf)
		if err != nil {
			return fmt.Errorf("failed to create SOCKS5 server: %v", err)
		}

		// Start server
		addr := fmt.Sprintf("%s:%s", config.ListenHost, config.ListenPort)
		log.Printf("Starting gobsips on %s", addr)
		if err := listenAndServe(server, "tcp", addr); err != nil {
			return fmt.Errorf("failed to start server: %v", err)
		}

	}

	return nil
}

var (
	getuid = os.Getuid

	listenAndServe = func(server *socks5.Server, network, addr string) error {
		return server.ListenAndServe(network, addr)
	}
)
