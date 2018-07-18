root_dir = $(shell git rev-parse --show-toplevel)

proto_idl = idl/service.proto
flat_idl = idl/service.fbs

all: golanggrpc/service.pb.go \
	 gofastgrpc/service.pb.go \
	 gogofastgrpc/service.pb.go \
	 gogofastergrpc/service.pb.go \
	 gogoslickgrpc/service.pb.go \
	 flatbuffer/service_generated.go

golanggrpc:
	mkdir -p golanggrpc

golanggrpc/service.pb.go: golanggrpc $(proto_idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --go_out=plugins=grpc:../golanggrpc/. -I. service.proto

gofastgrpc:
	mkdir -p gofastgrpc

gofastgrpc/service.pb.go: gofastgrpc $(proto_idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --gofast_out=plugins=grpc:../gofastgrpc/. -I. service.proto

gogofastgrpc:
	mkdir -p gogofastgrpc

gogofastgrpc/service.pb.go: gogofastgrpc $(proto_idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --gofast_out=plugins=grpc:../gogofastgrpc/. -I. service.proto

gogofastergrpc:
	mkdir -p gogofastergrpc

gogofastergrpc/service.pb.go: gogofastergrpc $(proto_idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --gofast_out=plugins=grpc:../gogofastergrpc/. -I. service.proto

gogoslickgrpc:
	mkdir -p gogoslickgrpc

gogoslickgrpc/service.pb.go: gogoslickgrpc $(proto_idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --gofast_out=plugins=grpc:../gogoslickgrpc/. -I. service.proto

flatbuffer:
	mkdir -p flatbuffer

flatbuffer/service_generated.go: flatbuffer $(flat_idl)
	docker run -v $(root_dir):$(root_dir) -w $(root_dir) neomantra/flatbuffers flatc --gen-onefile --go -o flatbuffer/ idl/service.fbs
