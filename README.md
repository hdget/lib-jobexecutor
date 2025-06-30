# How to generate protobuf files
```
protoc --go_out=./protobuf \
--go_opt=module=github.com/hdget/lib-jobexecutor/protobuf \
--go-grpc_out=./protobuf proto/*.proto \
--go-grpc_opt=module=github.com/hdget/lib-jobexecutor/protobuf
```