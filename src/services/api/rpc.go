package main

import (
	"github.com/slory7/angulargo/src/proto/api"
	"github.com/slory7/angulargo/src/services"
	"github.com/slory7/angulargo/src/services/api/srv"

	"github.com/nuveo/log"

	"github.com/asim/go-micro/v3"
)

func StartRpc() {
	service := micro.NewService(
		//micro.Server(grpc.NewServer()),
		//micro.Client(gc.NewClient()),
		micro.Name(services.ServiceNameApi),
	)

	service.Init()

	err := api.RegisterTrendingSrvHandler(service.Server(), &srv.APISrv{service.Client()})

	if err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
