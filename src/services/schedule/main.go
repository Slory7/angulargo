package main

import (
	"context"
	"errors"
	"math"

	"github.com/slory7/angulargo/src/services"
	trending "github.com/slory7/angulargo/src/services/trending/proto"

	"time"

	"github.com/slory7/angulargo/src/infrastructure/app"
	"github.com/slory7/angulargo/src/infrastructure/business/contracts"
	"github.com/slory7/angulargo/src/infrastructure/config"

	"github.com/nuveo/log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

func main() {
	log.Println("Schedule is running...")

	glbConfig = config.GetConfig(app.GetEnvironment(), &Config{}).(*Config)

	service := micro.NewService()
	service.Init()
	client := service.Client()

	gatherFunc := func() error {
		return gatherTrending(client)
	}

	log.Println("Run gatherTrending first")
	retryFunc(gatherFunc, "gatherTrending", 8)

	log.Printf("Start gatherTrending schedule, wait for %d seconds...\n", glbConfig.IntervalSeconds)
	tick := time.NewTicker(time.Second * time.Duration(glbConfig.IntervalSeconds))
	for range tick.C {
		log.Printf("Run gatherTrending scheduled, interval: %d seconds\n", glbConfig.IntervalSeconds)
		go retryFunc(gatherFunc, "gatherTrending", 8)
	}
}

func retryFunc(do func() error, logname string, count int) {
	for i := 0; i < count; i++ {
		sec := 0
		if i > 0 {
			sec = int(math.Pow(2, float64(i)))
		}
		log.Printf("Try %v, count: %d-%d, wait: %d seconds\n", logname, count, i+1, sec)
		time.Sleep(time.Second * time.Duration(sec))
		if err := do(); err == nil {
			log.Printf("Try %v done\n", logname)
			break
		} else if i == count-1 {
			log.Printf("Try %v failed\n", logname)
		}
	}
}

func gatherTrending(c client.Client) error {
	ctx := context.TODO()
	// ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	// defer cancel()
	ctx = services.GetContextWithTrace(ctx, "schedule.v1")
	services.PrintTrace(ctx, "gatherTrending")

	trendingClient := trending.NewTrendingService(services.ServiceNameTrending, c)
	req := &trending.Empty{}
	timeoutOpt := client.WithRequestTimeout(15 * time.Second)
	result, err := trendingClient.GetAndSaveGithubTrending(ctx, req, timeoutOpt)
	if err != nil {
		if errors.Is(err, contracts.BizErr) || contracts.IsLikeBizError(err) {
			log.Println(err)
			return nil
		}
		log.Errorf("get trending error:%v\n", err)
		return err
	}
	log.Println(result.Title)
	return nil
}
