[![Go Report Card](https://goreportcard.com/badge/github.com/presbrey/gobsips)](https://goreportcard.com/report/github.com/presbrey/gobsips)
[![codecov](https://codecov.io/gh/presbrey/gobsips/branch/main/graph/badge.svg)](https://codecov.io/gh/presbrey/gobsips)
![Go Test](https://github.com/presbrey/gobsips/workflows/Go%20Test/badge.svg)
[![GoDoc](https://godoc.org/github.com/presbrey/gobsips?status.svg)](https://godoc.org/github.com/presbrey/gobsips)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# GoBSIPS (Go Binary Systemd-Integrated Proxy Server)

GoBSIPS is a simple, single-file, static go binary that installs itself into Linux systemd and runs as a modern system service. It provides a SOCKS5 proxy server with authentication enabled by default.

## Features

- Single-file, static binary. SCP it to your servers and run it
- Easy permanent self-installation as a systemd service
- SOCKS5 proxy server with authentication
- Configurable through a simple configuration file

## Installation

Download a binary from GitHub here: https://github.com/presbrey/gobsips/releases/latest

Then to install, copy the binary to the server and run it. We'll use Linux/AMD64 in this example:
```
scp gobsips_linux_amd64 root@remote-server:/usr/local/bin/gobsips
ssh root@remote-server '/usr/local/bin/gobsips && cat /etc/sysconfig/gobsips'
```

`gobsips` will automatically:
1. Create a default configuration file at `/etc/sysconfig/gobsips`
2. Install a systemd service that runs `/usr/local/bin/gobsips`
3. Enable and start the service, and show you the initial configuration

## Configuration

The configuration file is located at `/etc/sysconfig/gobsips`. You can modify the following parameters:

- `USERNAME`: The username for SOCKS5 authentication
- `PASSWORD`: The password for SOCKS5 authentication
- `LISTEN_HOST`: The IP address to listen on (default: 0.0.0.0)
- `LISTEN_PORT`: The port to listen on (default: 1080)

After modifying the configuration, restart the service:

```
sudo systemctl restart gobsips
```

## Usage

Once installed and running, you can use GoBSIPS as a SOCKS5 proxy server. Configure your applications to use the proxy with the following details:

- Proxy Type: SOCKS5
- Host: [Your server's IP address]
- Port: 1080 (or the port you configured)
- Username and Password: As configured in the config file

If you run multiple servers and write golang software, this library may work well for you: https://github.com/presbrey/go-multiproxy

## Building from Source

To build GoBSIPS from source, you need Go installed on your system. Then run:

```
go build -o gobsips *.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
