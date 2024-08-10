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

func init() {
	flag.Parse()
	if *flagInstall && os.Getuid() != 0 {
		log.Fatal("You must install this program as root.")
	}
}

func main() {
	switch {

	case *flagInstall:

		if err := installDefaultConfig(); err != nil {
			log.Fatalf("Failed to write default config to %s: %v", configPath, err)
		}
		if err := installSystemdService(); err != nil {
			log.Fatal("Failed to install systemd service:", err)
		}
		fmt.Println("gobsips installed and started as a systemd service.")
		return

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
			log.Fatal(err)
		}

		// Start server
		addr := fmt.Sprintf("%s:%s", config.ListenHost, config.ListenPort)
		log.Printf("Starting gobsips on %s", addr)
		if err := server.ListenAndServe("tcp", addr); err != nil {
			log.Fatal(err)
		}

	}
}
