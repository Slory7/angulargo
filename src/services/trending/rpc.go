package main

import (
	"context"
	"errors"
	gather "services/gather/proto"
	"services/infrastructure/appstart"
	m "services/trending/datamodels"
	trending "services/trending/proto"
	"services/trending/services/githubtrending"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/nuveo/log"

	"github.com/jinzhu/copier"
	"github.com/micro/go-micro/client"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/trace"
)

func StartRpc() {
	service := micro.NewService(
		micro.Name("angulargo.micro.srv.trending"),
	)

	service.Init()

	trending.RegisterTrendingHandler(service.Server(), &TrendingSrv{service.Client()})

	service.Run()
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
	data, err := s.getGithubTrendingInternal(ctx)
	if err != nil {
		return err
	}

	serv, err2 := appstart.GetIoCInstance((*githubtrending.IGithubTrendingService)(nil))
	if err2 != nil {
		return err2
	}
	tServ := serv.(githubtrending.IGithubTrendingService)
	exists, err3 := tServ.SaveToDB(&data)
	if err3 != nil {
		return err3
	}
	if exists {
		return errors.New("Same title is already exist")
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
	result, err1 := gatherClient.GetHttpContent(ctx, rpcReq)
	if err1 != nil {
		log.Errorf("get %s error:%v\n", rpcReq.BaseUrl, err1)
		return data, err1
	}

	serv, err2 := appstart.GetIoCInstance((*githubtrending.IGithubTrendingDocService)(nil))
	if err2 != nil {
		return data, err2
	}
	docService := serv.(githubtrending.IGithubTrendingDocService)
	repos, err3 := docService.ParseDoc(result.Content)
	if err3 != nil {
		return data, err3
	}
	trending := m.GitTrendingAll{}
	t := time.Now()
	trending.Title = t.Format("Monday, 2 January 2006")
	trending.TrendingDate = t
	trending.GitRepos = repos

	return trending, nil
}
