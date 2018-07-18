root_dir = $(shell git rev-parse --show-toplevel)
gen_dir = gen
proto_idl = idl/service.proto
flat_idl = idl/service.fbs

all: $(gen_dir)/golanggrpc/service.pb.go \
	 $(gen_dir)/gofastgrpc/service.pb.go \
	 $(gen_dir)/gogofastgrpc/service.pb.go \
	 $(gen_dir)/gogofastergrpc/service.pb.go \
	 $(gen_dir)/gogoslickgrpc/service.pb.go \
	 $(gen_dir)/flatbuffer/service_generated.go

clean:
	rm -rf gen

$(gen_dir)/golanggrpc:
	mkdir -p $(gen_dir)/golanggrpc

$(gen_dir)/golanggrpc/service.pb.go: $(gen_dir)/golanggrpc $(proto_idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --go_out=plugins=grpc:../$(gen_dir)/golanggrpc/. -I. service.proto

$(gen_dir)/gofastgrpc:
	mkdir -p $(gen_dir)/gofastgrpc

$(gen_dir)/gofastgrpc/service.pb.go: $(gen_dir)/gofastgrpc $(proto_idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --gofast_out=plugins=grpc:../$(gen_dir)/gofastgrpc/. -I. service.proto

$(gen_dir)/gogofastgrpc:
	mkdir -p $(gen_dir)/gogofastgrpc

$(gen_dir)/gogofastgrpc/service.pb.go: $(gen_dir)/gogofastgrpc $(proto_idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --gofast_out=plugins=grpc:../$(gen_dir)/gogofastgrpc/. -I. service.proto

$(gen_dir)/gogofastergrpc:
	mkdir -p $(gen_dir)/gogofastergrpc

$(gen_dir)/gogofastergrpc/service.pb.go: $(gen_dir)/gogofastergrpc $(proto_idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --gofast_out=plugins=grpc:../$(gen_dir)/gogofastergrpc/. -I. service.proto

$(gen_dir)/gogoslickgrpc:
	mkdir -p $(gen_dir)/gogoslickgrpc

$(gen_dir)/gogoslickgrpc/service.pb.go: $(gen_dir)/gogoslickgrpc $(proto_idl)
	docker run --rm -v $(root_dir):$(root_dir) -w $(root_dir)/idl znly/protoc --gofast_out=plugins=grpc:../$(gen_dir)/gogoslickgrpc/. -I. service.proto

$(gen_dir)/flatbuffer:
	mkdir -p $(gen_dir)/flatbuffer

$(gen_dir)/flatbuffer/service_generated.go: $(gen_dir)/flatbuffer $(flat_idl)
	docker run -v $(root_dir):$(root_dir) -w $(root_dir) neomantra/flatbuffers flatc --gen-onefile --grpc --go -o $(gen_dir)/flatbuffer/ idl/service.fbs
