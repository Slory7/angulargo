package main

import (
	"context"
	"time"

	"github.com/slory7/angulargo/src/infrastructure/app"
	"github.com/slory7/angulargo/src/infrastructure/business/contracts"
	"github.com/slory7/angulargo/src/services"
	gather "github.com/slory7/angulargo/src/services/gather/proto"
	m "github.com/slory7/angulargo/src/services/trending/datamodels"
	trending "github.com/slory7/angulargo/src/services/trending/proto"
	"github.com/slory7/angulargo/src/services/trending/services/githubtrending"

	"github.com/nuveo/log"

	"github.com/slory7/copier"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"

	"golang.org/x/net/trace"
)

func StartRpc() {
	service := micro.NewService(
		micro.Name(services.ServiceNameTrending),
	)

	service.Init()

	err := trending.RegisterTrendingHandler(service.Server(), &TrendingSrv{service.Client()})
	if err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type TrendingSrv struct {
	Client client.Client
}

func (s *TrendingSrv) GetGithubTrending(ctx context.Context, req *trending.Request, rsp *trending.GithubTrendingInfo) error {
	srv := app.Instance.GetIoCInstanceMust((*githubtrending.IGithubTrendingService)(nil)).(githubtrending.IGithubTrendingService)
	info, b, err := srv.GetTrendingInfo(req.Title)
	if err != nil {
		return err
	}
	if !b {
		return contracts.NewBizError("Trending info does not exist.", contracts.NotFound)
	}
	copier.Copy(rsp, info)
	//rsp.TrendingDate, _ = ptypes.TimestampProto(info.TrendingDate)
	return nil
}

func (s *TrendingSrv) FetchGithubTrending(ctx context.Context, req *trending.Empty, rsp *trending.GithubTrendingInfo) error {
	data, err := s.getGithubTrendingInternal(ctx)
	if err != nil {
		return err
	}
	copier.Copy(rsp, data)
	return nil
}

func (s *TrendingSrv) GetAndSaveGithubTrending(ctx context.Context, req *trending.Request, rsp *trending.GithubTrendingInfo) error {
	srv := app.Instance.GetIoCInstanceMust((*githubtrending.IGithubTrendingService)(nil)).(githubtrending.IGithubTrendingService)

	t := time.Now()
	title := t.Format("Monday, 2 January 2006")
	exists, err := srv.IsTitleExists(title)
	if err != nil {
		return err
	}
	if exists {
		return contracts.NewBizError("Same title is already exist: "+title, contracts.Conflict)
	}

	data, err := s.getGithubTrendingInternal(ctx)
	if err != nil {
		return err
	}

	exists, err = srv.SaveToDB(&data)
	if err != nil {
		return err
	}
	if exists {
		return contracts.NewBizError("Same title is already exist: "+data.Title, contracts.Conflict)
	}
	copier.Copy(rsp, data)
	return nil
}

func (s *TrendingSrv) getGithubTrendingInternal(ctx context.Context) (data m.GitTrendingAll, err error) {
	md, _ := metadata.FromContext(ctx)
	traceID := md["Traceid"]

	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("fromName %s", md["Fromname"])
		tr.LazyPrintf("traceID %s", traceID)
	}
	log.Printf("fromName %s\n", md["Fromname"])
	log.Printf("traceID %s\n", traceID)

	gatherClient := gather.NewGatherService(services.ServiceNameGather, s.Client)
	rpcReq := &gather.Request{BaseUrl: glbConfig.TrendingURL, Method: "GET", TimeOut: 5}
	result, err := gatherClient.GetHttpContent(ctx, rpcReq)
	if err != nil {
		log.Errorf("get %s error:%v\n", rpcReq.BaseUrl, err)
		return data, err
	}

	serv, err := app.Instance.GetIoCInstance((*githubtrending.IGithubTrendingDocService)(nil))
	if err != nil {
		return data, err
	}
	docService := serv.(githubtrending.IGithubTrendingDocService)
	repos, err := docService.ParseDoc(result.Content)
	if err != nil {
		return data, err
	}
	for _, rep := range repos {
		if err := app.Instance.Validator.Struct(rep); err != nil {
			return data, err
		}
	}
	trending := m.GitTrendingAll{}
	t := time.Now()
	trending.Title = t.Format("Monday, 2 January 2006")
	trending.TrendingDate = t
	trending.GitRepos = repos

	return trending, nil
}
