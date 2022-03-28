package main

type TLSConfig struct {
	UseTLS bool `json:"useTLS" cfg:"useTLS" cfgDefault:false`
	//CAServerName string `json:"caServerName" cfg:"caServerName"`
	CA   string `json:"ca" cfg:"ca"`
	Cert string `json:"cert" cfg:"cert"`
	Key  string `json:"key" cfg:"key"`
}
type Config struct {
	GrpcServerAddress string    `json:"grpcServerAddress" cfg:"grpcServerAddress" cfgDefault:":8010"`
	TLS               TLSConfig `json:"tls" cfg:"tls"`
}
