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

	"github.com/golang/protobuf/ptypes"

	"github.com/nuveo/log"

	"github.com/jinzhu/copier"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"

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
	data, err := s.getGithubTrendingInternal(ctx)
	if err != nil {
		return err
	}
	rsp.Title = data.Title
	rsp.TrendingDate, _ = ptypes.TimestampProto(data.TrendingDate)
	copier.Copy(&rsp.GitRepos, &data.GitRepos)

	return nil
}

func (s *TrendingSrv) GetAndSaveGithubTrending(ctx context.Context, req *trending.Request, rsp *trending.GithubTrendingInfo) error {
	serv, err := app.Instance.GetIoCInstance((*githubtrending.IGithubTrendingService)(nil))
	if err != nil {
		return err
	}
	tServ := serv.(githubtrending.IGithubTrendingService)

	t := time.Now()
	title := t.Format("Monday, 2 January 2006")
	exists, err := tServ.IsTitleExists(title)
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

	exists, err = tServ.SaveToDB(&data)
	if err != nil {
		return err
	}
	if exists {
		return contracts.NewBizError("Same title is already exist: "+data.Title, contracts.Conflict)
	}

	rsp.Title = data.Title
	rsp.TrendingDate, _ = ptypes.TimestampProto(data.TrendingDate)
	copier.Copy(&rsp.GitRepos, &data.GitRepos)

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

	gatherClient := gather.NewGatherService("angulargo.micro.srv.gather", s.Client)
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
