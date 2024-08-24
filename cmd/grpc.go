package main

import (
	"crypto/tls"
	"crypto/x509"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func InitialGRpc(target string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	opts = append(opts, grpc.WithTransportCredentials(cred), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	cc, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, err
	}
	return cc, nil
}
