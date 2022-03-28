package main

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nuveo/log"
	"github.com/slory7/angulargo/src/infrastructure/app"
	"github.com/slory7/angulargo/src/infrastructure/config"
	"github.com/slory7/angulargo/src/infrastructure/framework/security/gtls"
	"github.com/slory7/angulargo/src/proto/api/gw/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func run() error {
	cfg := config.GetConfigFull[Config](app.GetEnvironment(), true)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()

	var c credentials.TransportCredentials
	if cfg.GrpcClientTLS.UseTLS {
		tlsClient := gtls.Client{
			ServerName: cfg.GrpcClientTLS.CAServerName,
			CaFile:     cfg.GrpcClientTLS.CA,
			CertFile:   cfg.GrpcClientTLS.Cert,
			KeyFile:    cfg.GrpcClientTLS.Key,
		}
		var err error
		c, err = tlsClient.GetCredentialsByCA()
		if err != nil {
			return err
		}
	} else {
		c = insecure.NewCredentials()
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(c)}

	err := api.RegisterApiHandlerFromEndpoint(ctx, mux, cfg.GrpcServerAddress, opts)
	if err != nil {
		return err
	}

	cors(mux, cfg.AllowedOrigins, cfg.AllowMethods)

	addr := cfg.Addr

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	if cfg.TLS.UseTLS {
		log.Printf("Listening https restful proxy at %s\n", addr)
		return http.ListenAndServeTLS(addr, cfg.TLS.Cert, cfg.TLS.Key, mux)
	}
	log.Printf("Listening http restful proxy at %s\n", addr)
	return http.ListenAndServe(addr, mux)
}

func cors(next http.Handler, allowedOrigins string, allowMethods string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
		if allowMethods != "" {
			w.Header().Set("Access-Control-Allow-Methods", allowMethods)
		}
		// w.Header().Set(allow_headers, headers)
		// w.Header().Set(allow_credentials, credentials)
		// w.Header().Set(expose_headers, headers)

		// If this was preflight options request let's write empty ok response and return
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Write(nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
