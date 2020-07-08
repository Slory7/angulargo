package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/slory7/angulargo/src/services"
	api "github.com/slory7/angulargo/src/services/api/proto"
	trending "github.com/slory7/angulargo/src/services/trending/proto"

	"github.com/slory7/copier"

	"github.com/nuveo/log"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"golang.org/x/net/trace"
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
	md, _ := metadata.FromContext(ctx)
	traceID := md["X-Request-Id"]
	if len(traceID) == 0 {
		traceID = md["Traceid"]
	}
	if len(traceID) == 0 {
		traceID = uuid.New().String()
	}
	if len(md["Fromname"]) == 0 {
		md["Fromname"] = "api.v1"
	}
	md["Traceid"] = traceID

	ctx = metadata.NewContext(ctx, md)

	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("fromName %s", md["Fromname"])
		tr.LazyPrintf("traceID %s", traceID)
	}
	log.Printf("fromName %s\n", md["Fromname"])
	log.Printf("traceID %s\n", traceID)

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
