package main

import (
	"context"
	"flag"
	"log"

	"github.com/slory7/angulargo/src/proto/api"
	"github.com/slory7/angulargo/src/proto/trending"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8010", "gRPC server endpoint")
)

func main() {
	flag.Parse()
	// tlsClient := gtls.Client{
	// 	ServerName: "localhost",
	// 	CaFile:     "ca.crt",
	// 	CertFile:   "client.pem",
	// 	KeyFile:    "client.key",
	// }

	//c, err := tlsClient.GetCredentialsByCA()
	c := insecure.NewCredentials()
	conn, err := grpc.Dial(*grpcServerEndpoint, grpc.WithTransportCredentials(c))
	if err == nil {
		c := api.NewApiClient(conn)
		info, err := c.GetGithubTrending(context.Background(), &trending.Request{Title: "tt"})
		log.Print(err)
		log.Println(info)
	} else {
		log.Print(err)

	}
	defer conn.Close()

}
