package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/armon/go-socks5"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	// Save original flag values and restore them after the test
	origDaemonFlag := *flagDaemon
	origInstallFlag := *flagInstall
	defer func() {
		*flagDaemon = origDaemonFlag
		*flagInstall = origInstallFlag
	}()

	// Test installation path
	*flagInstall = true
	*flagDaemon = false

	// Redirect stdout to capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := do()
	if os.Getuid() != 0 {
		assert.Error(t, err, "you must install this program as root")
	}

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var output []byte
	r.Read(output)

	if os.Getuid() != 0 {
		assert.Equal(t, "", string(output))
	}

	getuid = func() int {
		return 0
	}
	err = do()
	assert.Error(t, err, "failed to write default config to %s: %v", configPath, err)
	configPath = "/tmp/test-config"
	err = do()
	assert.Error(t, err, "failed to install systemd service: %v", err)
	run = func(cmd *exec.Cmd) error {
		return nil
	}
	systemdPath = "/tmp/test-systemd-service"
	err = do()
	assert.NoError(t, err)

	*flagInstall = false
	*flagDaemon = true
	listenAndServe = func(_ *socks5.Server, network, addr string) error {
		return fmt.Errorf("testing without starting server (%v:%v)", network, addr)
	}
	err = do()
	assert.Error(t, err, "failed to start server: testing without starting server (tcp:0.0.0.0:1080)")
	listenAndServe = func(_ *socks5.Server, _, _ string) error {
		return nil
	}
	err = do()
	assert.NoError(t, err)
}

func TestMain(m *testing.M) {
	// This is needed to properly initialize flags for tests
	flag.Parse()
	os.Exit(m.Run())
}
