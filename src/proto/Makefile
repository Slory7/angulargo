build:
	protoc -I. --proto_path=$(GOPATH)\src --go_out=paths=source_relative:. --micro_out=paths=source_relative:. gather/gather.proto 
	protoc -I. --proto_path=$(GOPATH)\src --go_out=paths=source_relative:. --micro_out=paths=source_relative:. trending/trending.proto
	protoc -I. --proto_path=$(GOPATH)\src --go_out=plugins=grpc,paths=source_relative:. --micro_out=paths=source_relative:. api/api.proto
protoapi:
	protoc -I. --proto_path=$(GOPATH)\src --go_out=plugins=grpc,paths=source_relative:. --micro_out=paths=source_relative:. api/api.proto
dart:
	protoc -I. --proto_path=$(GOPATH)\src --dart_out=grpc:../client/proto  trending/trending.proto
	protoc -I. --proto_path=$(GOPATH)\src --dart_out=grpc:../client/proto  api/api.proto
	