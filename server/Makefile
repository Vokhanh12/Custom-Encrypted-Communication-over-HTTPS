PROTO_DIR := proto
GOOGLE_API_DIR := $(PROTO_DIR)/google/api
GO_OUT_DIR := api
API_OUT_DIR := api

GOOGLE_PROTO_FILES := $(shell find $(GOOGLE_API_DIR) -name '*.proto')
MAIN_PROTO_FILES := $(shell find $(PROTO_DIR) -name '*.proto' ! -path "$(GOOGLE_API_DIR)/*")

INCLUDE_DIRS := -I$(PROTO_DIR) -I$(GOOGLE_API_DIR) -I$(PROTO_DIR)/google

all: clean gen_google_api gen_main

clean:
	find $(GO_OUT_DIR) -name "*.pb.go" -type f -delete
	rm -rf $(API_OUT_DIR)

gen_google_api:
	@for proto_file in $(GOOGLE_PROTO_FILES); do \
		echo "Generating [google/api] $$proto_file..."; \
		protoc \
			$(INCLUDE_DIRS) \
			--go_out=$(GO_OUT_DIR) \
			--go_opt=paths=source_relative \
			$$proto_file; \
	done

gen_main:
	@for proto_file in $(MAIN_PROTO_FILES); do \
		echo "Generating [main] $$proto_file..."; \
		protoc \
			$(INCLUDE_DIRS) \
			--go_out=$(GO_OUT_DIR) \
			--go-grpc_out=$(GO_OUT_DIR) \
			--go_opt=paths=source_relative \
			--go-grpc_opt=paths=source_relative \
			--validate_out=lang=go:$(GO_OUT_DIR) \
			--validate_opt=paths=source_relative \
			$$proto_file || exit 1; \
	done

.PHONY: all clean gen_google_api gen_main
