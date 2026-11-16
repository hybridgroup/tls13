# tls13

A TLS 1.3 client for microcontrollers, written in Go and built with TinyGo. It uses [lcrypto](https://github.com/soypat/lcrypto) for heapless cryptography and the record layer from [lneto](https://github.com/soypat/lneto).

Buffers are allocated when a connection opens and reused by later connections. The handshake buffers are released once the handshake is done.

## Features

* Server verification with any `lcrypto.Verifier`, such as the `x509.Verifier` from lcrypto
* Client certificates in answer to a CertificateRequest
* AES-128-GCM with SHA-256
* X25519 and P-256 key exchange
* SNI and a single ALPN protocol
* `MaxRecord` to cap the read buffer, such as 4096 bytes on small devices

Not supported yet: PSK, session resumption, 0-RTT, HelloRetryRequest and KeyUpdate. Session tickets from the server are read and dropped.

## Usage

See [example_test.go](example_test.go). Configure a `Conn` once, then call `Open` and `Handshake` for each new TCP connection. A `Conn` handles one connection at a time.

## Tests

```
go test ./...
```

The tests run the client against the Go standard library TLS server.
