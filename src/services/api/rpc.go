package main

import (
	"context"
	"time"

	"github.com/slory7/angulargo/src/proto/api"
	"github.com/slory7/angulargo/src/proto/trending"
	"github.com/slory7/angulargo/src/services"

	"github.com/slory7/copier"

	"github.com/nuveo/log"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

func StartRpc() {
	service := micro.NewService(
		micro.Name(services.ServiceNameApi),
	)

	service.Init()

	err := api.RegisterTrendingSrvHandler(service.Server(), &APISrv{service.Client()})

	if err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type APISrv struct {
	Client client.Client
}

func (s *APISrv) GetGithubTrending(ctx context.Context, req *trending.Request, rsp *trending.GithubTrendingInfo) error {
	ctx = services.GetContextWithTrace(ctx, "api.v1")
	services.PrintTrace(ctx, "GetGithubTrending")
	if req.Title == "" {
		req.Title = time.Now().Format("Monday, 2 January 2006")
	}
	trendingClient := trending.NewTrendingService(services.ServiceNameTrending, s.Client)
	result, err := trendingClient.GetGithubTrending(ctx, req)
	if err != nil {
		log.Errorf("get trending error:%v\n", err)
		return err
	}
	copier.Copy(rsp, result)
	return nil
}
