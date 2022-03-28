package main

type TLSConfig struct {
	UseTLS       bool   `json:"useTLS" cfg:"useTLS" cfgDefault:false`
	CAServerName string `json:"caServerName" cfg:"caServerName"`
	CA           string `json:"ca" cfg:"ca"`
	Cert         string `json:"cert" cfg:"cert"`
	Key          string `json:"key" cfg:"key"`
}
type Config struct {
	Addr              string    `json:"addr" cfg:"addr" cfgDefault:":8011"`
	TLS               TLSConfig `json:"tls" cfg:"tls"`
	GrpcServerAddress string    `json:"grpcServerAddress" cfg:"grpcServerAddress" cfgDefault:":8010"`
	GrpcClientTLS     TLSConfig `json:"grpcClientTLS" cfg:"grpcClientTLS"`
	AllowedOrigins    string    `json:"allowedOrigins" cfg:"allowedOrigins" cfgDefault:"*"`
	AllowMethods      string    `json:"allowMethods" cfg:"allowMethods"`
}
