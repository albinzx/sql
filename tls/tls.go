package tls

import (
	"crypto/tls"
	"crypto/x509"
)

func WithCA(ca []byte) *tls.Config {
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(ca)

	tlsCfg := &tls.Config{
		RootCAs:    caCertPool,
		MinVersion: tls.VersionTLS12,
	}

	return tlsCfg
}

func WithServerAndCA(serverName string, ca []byte) *tls.Config {
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(ca)

	tlsCfg := &tls.Config{
		RootCAs:    caCertPool,
		ServerName: serverName,
		MinVersion: tls.VersionTLS12,
	}

	return tlsCfg
}
