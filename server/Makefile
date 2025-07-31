PROTO_DIR := ./proto
OUT_DIR := ../

PROTO_FILES := $(shell find $(PROTO_DIR) -name '*.proto')

generate:
	protoc --proto_path=$(PROTO_DIR) \
	       --go_out=$(OUT_DIR) \
	       --go-grpc_out=$(OUT_DIR) \
	       $(PROTO_FILES)
