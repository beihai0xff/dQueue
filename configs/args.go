// Package configs provides config management
// args.go provides some common command line arguments.
package configs

import (
	"flag"
)

var (
	// ConfigPath config file path
	ConfigPath = flag.String("config-filepath", "./config.yaml", "The server config file path")

	// GRPCPort grpc server port
	GRPCPort = flag.Int("grpc-port", DefaultGRPCPort, "The grpc server grpcPort")
	// HTTPPort http server port
	HTTPPort = flag.Int("http-port", DefaultHTTPPort, "The http server grpcPort")

	// CertPath tls cert file path, the file must contain PEM encoded data.
	CertPath = flag.String("tls-cert-path", DefaultCertPath, "The TLS cert file path")
	// KeyPath tls key file path, the file must contain PEM encoded data.
	KeyPath = flag.String("tls-key-path", DefaultKeyPath, "The TLS key file path")
	// NameServerURL name server url
	NameServerURL = flag.String("name-server-url", DefaultNameServerURL, "The name server connection url")
)
