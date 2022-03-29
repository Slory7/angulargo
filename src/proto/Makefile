build:
	protoc -I. --proto_path=$(GOPATH)\src --go_out=paths=source_relative:. --micro_out=paths=source_relative:. gather/gather.proto 
	protoc -I. --proto_path=$(GOPATH)\src --go_out=paths=source_relative:. --micro_out=paths=source_relative:. trending/trending.proto
	protoc -I. --proto_path=$(GOPATH)\src --go-grpc_out=paths=source_relative:. --micro_out=paths=source_relative:. api/api.proto
protoapi:
	protoc -I. --proto_path=$(GOPATH)\src --go-grpc_out=paths=source_relative:. --micro_out=paths=source_relative:. api/api.proto
gateway:
	protoc -I. --proto_path=$(GOPATH)\src --go-grpc_out=./api/gw --go-grpc_opt paths=source_relative api/api.proto
	protoc -I. --proto_path=$(GOPATH)\src --grpc-gateway_out=./api/gw --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=generate_unbound_methods=true --grpc-gateway_opt=paths=source_relative api/api.proto 
protots:
	protoc -I. --proto_path=$(GOPATH)\src --grpc-web_out=import_style=typescript,mode=grpcweb:../client/ts/proto api/api.proto
	protoc -I. --proto_path=$(GOPATH)\src --grpc-web_out=import_style=typescript,mode=grpcweb:../client/ts/proto trending/trending.proto
dart:
	protoc -I. --proto_path=$(GOPATH)\src --dart_out=grpc:../client/proto  trending/trending.proto
	protoc -I. --proto_path=$(GOPATH)\src --dart_out=grpc:../client/proto  api/api.proto
	