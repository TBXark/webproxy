# webproxy

## Overview

`webproxy` is a simple HTTP reverse proxy server that routes requests through a SOCKS5 proxy. It is useful for routing traffic through a SOCKS5 proxy to a target server.

## Usage

### Build

To build the project, run:

```sh
go build -o webproxy main.go
```

### Run

To run the proxy server, use the following command:

```sh
./webproxy -target <target_url> -socks5 <socks5_address> -listen <listen_address>
```

- `target`: The target URL to which the requests will be proxied.
- `socks5`: The SOCKS5 proxy address.
- `listen`: The address on which the proxy server will listen (default is `:8080`).

### Example

```sh
./webproxy -target http://example.com -socks5 127.0.0.1:1080 -listen :8080
```

This will start the proxy server on port 8080, routing traffic to `http://example.com` through the SOCKS5 proxy at `127.0.0.1:1080`.

## License

This project is licensed under the MIT License.

