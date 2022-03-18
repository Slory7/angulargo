package srv

import (
	"context"
	"time"

	"github.com/asim/go-micro/v3/client"
	"github.com/nuveo/log"
	"github.com/slory7/angulargo/src/proto/trending"
	"github.com/slory7/angulargo/src/services"
	"github.com/slory7/copier"
)

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
