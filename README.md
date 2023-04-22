# go-ldm
ldm streaming server

grpc: 
1. private api
2. faster than json by 6 times
3. type-safety; (inversify, class-validation)
4. fast serialization (marshall and unmarshall)
5. versioning in protobuf


## Steps for grpc

```shell
cd client
go mod init github.com/xcheng85/go-ldm/client

cd server
go mod init github.com/xcheng85/go-ldm/server

go work init client server

# language specific runtime for compiling protobuf
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

# go: module github.com/golang/protobuf is deprecated: Use the "google.golang.org/protobuf" module instead.
go mod graph | grep github.com/golang/protobuf
go mod why github.com/golang/protobuf

# in root dir of server
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/v1/ldm.proto

## it will create ldm_grpc.pb.go and ldm.pb.go in folder server/api/v1
```

## Day1: Build internal grpc server
```shell
mkdir -p internal/server
# for unit test
go get github.com/stretchr/testify
```
