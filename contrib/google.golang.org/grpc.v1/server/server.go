package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/b2wdigital/goignite/v2/core/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"
)

type Ext func(ctx context.Context) []grpc.ServerOption

type Server struct {
	server           *grpc.Server
	serviceRegistrar grpc.ServiceRegistrar
	options          *Options
}

func NewServer(ctx context.Context, exts ...Ext) *Server {
	opt, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return NewServerWithOptions(ctx, opt, exts...)
}

func NewServerWithOptions(ctx context.Context, opt *Options, exts ...Ext) *Server {

	logger := log.FromContext(ctx)

	err := gzip.SetLevel(9)
	if err != nil {
		logger.Fatalf("could not set level: %s", err.Error())
	}

	var s *grpc.Server

	var serverOptions []grpc.ServerOption

	if opt.TLS.Enabled {

		logger.Debug("configuring tls on grpc server")

		var creds credentials.TransportCredentials
		certPool := x509.NewCertPool()

		if opt.TLS.CertFile != "" && opt.TLS.KeyFile != "" {

			logger.Trace("configuring cert and key certificates on grpc server")

			// Load the certificates from disk
			certificate, err := tls.LoadX509KeyPair(opt.TLS.CertFile, opt.TLS.KeyFile)
			if err != nil {
				logger.Fatalf("could not load server key pair: %s", err.Error())
			}

			logger.Trace("cert and key certificates loaded")

			if opt.TLS.CAFile != "" {

				logger.Trace("configuring ca certificate on grpc server")

				ca, err := ioutil.ReadFile(opt.TLS.CAFile)
				if err != nil {
					logger.Fatalf("could not read ca certificate: %s", err.Error())
				}

				// Append the client certificates from the CA
				if ok := certPool.AppendCertsFromPEM(ca); !ok {
					logger.Fatalf("failed to append client certs")
				}

				logger.Trace("ca certificate loaded")

			}

			// Create the TLS credentials
			creds = credentials.NewTLS(&tls.Config{
				ClientAuth:   tls.NoClientCert,
				Certificates: []tls.Certificate{certificate},
				ClientCAs:    certPool,
			})

		} else {

			creds = credentials.NewTLS(&tls.Config{
				ClientAuth:   tls.NoClientCert,
				Certificates: []tls.Certificate{},
				ClientCAs:    certPool,
			})

		}

		serverOptions = append(serverOptions, grpc.Creds(creds))
	}

	for _, ext := range exts {
		serverOptions = append(serverOptions, ext(ctx)...)
	}

	serverOptions = append(serverOptions, grpc.MaxConcurrentStreams(uint32(opt.MaxConcurrentStreams)))
	serverOptions = append(serverOptions, grpc.InitialConnWindowSize(opt.InitialConnWindowSize))
	serverOptions = append(serverOptions, grpc.InitialWindowSize(opt.InitialWindowSize))

	s = grpc.NewServer(serverOptions...)

	return &Server{
		server:  s,
		options: opt,
	}
}

func (s *Server) Server() *grpc.Server {
	return s.server
}

func (s *Server) ServiceRegistrar() grpc.ServiceRegistrar {
	return s.server
}

func (s *Server) Serve(ctx context.Context) {

	logger := log.FromContext(ctx)

	service.RegisterChannelzServiceToServer(s.server)

	// Register reflection service on gRPC server.
	reflection.Register(s.server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.options.Port))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err.Error())
	}

	logger.Infof("grpc server started on port %v", s.options.Port)

	if err := s.server.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err.Error())
	}

}
