# GoBSIPS (Go Binary Systemd-Integrated Proxy Server)

GoBSIPS is a simple, single-file, static go binary that installs itself into Linux systemd and runs as a modern system service. It provides a SOCKS5 proxy server with basic authentication.

## Features

- Single-file, static binary. SCP it to your servers and run it
- Easy permanent self-installation as a systemd service
- SOCKS5 proxy server with authentication
- Configurable through a simple configuration file

## Installation

To install GoBSIPS, run the following command:

```
sudo mv gobsips /usr/local/bin
sudo /usr/local/bin/gobsips
```

This will:
1. Create a default configuration file at `/etc/sysconfig/gobsips`
2. Install a systemd service that runs to `/usr/local/bin/gobsips`
3. Enable and start the service

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

## Building from Source

To build GoBSIPS from source, you need Go installed on your system. Then run:

```
go build -o gobsips *.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
