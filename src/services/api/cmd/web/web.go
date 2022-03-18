package main

import (
	"net/http"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/web"
	"github.com/nuveo/log"
	"github.com/slory7/angulargo/src/proto/api"
	"github.com/slory7/angulargo/src/services"
	"github.com/slory7/angulargo/src/services/api/srv/rpcsrv"
	"google.golang.org/grpc"
)

func main() {
	s := micro.NewService(
		micro.Name(services.ServiceNameApiWeb),
	)
	service := web.NewService(
		web.MicroService(s),
		//web.TLSConfig()
	)
	service.Init()

	gs := grpc.NewServer()
	api.RegisterTrendingSrvServer(gs, &rpcsrv.APIRPCSrv{Client: client.DefaultClient})

	service.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "application/grpc")
		gs.ServeHTTP(w, r)
	})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
