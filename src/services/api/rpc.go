package main

import (
	_ "github.com/asim/go-micro/plugins/registry/etcd/v3"
	mgrpc "github.com/asim/go-micro/plugins/server/grpc/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/asim/go-micro/v3/transport"
	"github.com/nuveo/log"
	"github.com/slory7/angulargo/src/infrastructure/app"
	"github.com/slory7/angulargo/src/infrastructure/config"
	"github.com/slory7/angulargo/src/infrastructure/framework/security/gtls"
	"github.com/slory7/angulargo/src/proto/api"
	"github.com/slory7/angulargo/src/services"
	"github.com/slory7/angulargo/src/services/api/srv"
)

func StartRpc() {
	var s server.Server

	cfg := config.GetConfig[Config](app.GetEnvironment())
	if cfg.TLS.UseTLS {
		tlsServer := gtls.Server{
			CaFile:   cfg.TLS.CA,
			CertFile: cfg.TLS.Cert,
			KeyFile:  cfg.TLS.Key,
		}
		t, err := tlsServer.GetTLSConfig()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Use ssl\n")
		s = mgrpc.NewServer(server.TLSConfig(t))
	} else {
		s = mgrpc.NewServer()
	}

	service := micro.NewService(
		micro.Server(s),
		//micro.Client(),
		micro.Name(services.ServiceNameApi),
		micro.Address(cfg.GrpcServerAddress),
		micro.Transport(transport.NewHTTPTransport(
			transport.Secure(false),
		)),
	)

	service.Init()

	err := api.RegisterApiHandler(service.Server(), &srv.APISrv{service.Client()})

	if err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
