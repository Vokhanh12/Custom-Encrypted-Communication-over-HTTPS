protoc \
  --proto_path=./proto \
  --go_out=../ \
  --go-grpc_out=../ \
  ./proto/user/user.proto
