package tls13_test

import (
	"crypto/rand"
	"net"
	"time"

	"github.com/hybridgroup/tls13"
	"github.com/soypat/lcrypto/x509"
)

func Example() {
	var (
		verifier x509.Verifier
		p256     x509.P256Verifier
		schemes  [8]uint16
		conn     tls13.Conn
	)
	err := verifier.Configure(x509.VerifierConfig{
		Roots:      x509.Chain{rootDER},
		Nanotime:   func() int64 { return time.Now().UnixNano() },
		Algorithms: []x509.SignatureVerifier{&p256},
	})
	if err != nil {
		panic(err)
	}
	err = conn.Configure(tls13.Config{
		Rand:         rand.Reader,
		ServerName:   "example.com",
		Verifier:     &verifier,
		Schemes:      verifier.Schemes(schemes[:0]),
		Suites:       []tls13.Suite{tls13.AES128GCMSHA256()},
		KeyExchanges: []tls13.KeyExchange{tls13.X25519()},
		MaxRecord:    4096,
	})
	if err != nil {
		panic(err)
	}
	tcp, err := net.Dial("tcp", "example.com:443")
	if err != nil {
		panic(err)
	}
	if err = conn.Open(tcp); err != nil {
		panic(err)
	}
	if err = conn.Handshake(); err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte("HEAD / HTTP/1.1\r\nHost: example.com\r\nConnection: close\r\n\r\n"))
}

// rootDER is the DER certificate of the CA that issued the server's chain.
var rootDER []byte
