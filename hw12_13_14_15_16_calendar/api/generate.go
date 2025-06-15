package api

//go:generate bash -c "rm -rf ./pb && mkdir -p ./pb"
//go:generate protoc --proto_path=../api --go_out=../api/pb --go_opt=paths=source_relative --go-grpc_out=../api/pb --go-grpc_opt=paths=source_relative ../api/event.proto
