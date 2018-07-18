root_dir = $(shell git rev-parse --show-toplevel)

idl = idl/service.proto

all: golanggrpc/service.pb.go gofastgrpc/service.pb.go gogofastgrpc/service.pb.go gogofastergrpc/service.pb.go gogoslickgrpc/service.pb.go

golanggrpc:
	mkdir -p golanggrpc

golanggrpc/service.pb.go: golanggrpc $(idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --go_out=plugins=grpc:../golanggrpc/. -I. service.proto

gofastgrpc:
	mkdir -p gofastgrpc

gofastgrpc/service.pb.go: gofastgrpc $(idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --gofast_out=plugins=grpc:../gofastgrpc/. -I. service.proto

gogofastgrpc:
	mkdir -p gogofastgrpc

gogofastgrpc/service.pb.go: gogofastgrpc $(idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --gofast_out=plugins=grpc:../gogofastgrpc/. -I. service.proto

gogofastergrpc:
	mkdir -p gogofastergrpc

gogofastergrpc/service.pb.go: gogofastergrpc $(idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --gofast_out=plugins=grpc:../gogofastergrpc/. -I. service.proto

gogoslickgrpc:
	mkdir -p gogoslickgrpc

gogoslickgrpc/service.pb.go: gogoslickgrpc $(idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --gofast_out=plugins=grpc:../gogoslickgrpc/. -I. service.proto

