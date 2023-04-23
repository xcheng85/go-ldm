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

go test -timeout 30s -run ^TestServer$ github.com/xcheng85/go-ldm/server/internal/server
```

## Day2: Error handling GRPC
google grpc status pkg
google grpc codes pkg
google errdetails pkg


## Day3: Inversion of Container (IOC)
LDM interface with multiple implementations
### Logs
ordered data
write-ahead logs (redis, message queue, version control)
data loss/system failure

Raft: consensus algorithm, leader, follower, state machine with log as input
Redux: log changes, action, state


tile(256 mb) ---> block (fitting in memory, 8gb) ---> volumentric dataset (500gb)

helper: index 

## Tile

tile: id and offset
block: active block 

each block: binary block + index file.
index file: index of each tile in a block file and a block file

how to deal with data in disk that is bigger than the available memory ?
make use of virtual memory and the concept of memory-mapped files.

memory mapping for index file
https://ghvsted.com/blog/exploring-mmap-in-go/
https://brunocalza.me/discovering-and-exploring-mmap-using-go/
package: mmap-go


## Day4: cli
cobra
viper

## Day5: secure of GRPC
1. TLS: in-flight, man in the mid
    1. 
2. Authenticate
    1. client-server
        1. username-password
        2. oauth2.0 token
    2. server-server
        1. mutual TLS
            1. CloudFlare: cfssl, cfssljson
3. Authorization
    1. ACL: access control list
        1. Casbin pkg
        2. client will create root/non-root pem certificate 

## Day6: Observability of GRPC
1. Logging: 
    1. no Open-Telemetr loggin spec
    2. uber zap
    3. https://github.com/grpc-ecosystem/go-grpc-middleware/blob/main/interceptors/logging/examples/zap/example_test.go
2. Metrics: 
    1. OpenTelemetry have non metrics support
    2. OpenCencus 

## Day7: Server to Server Service Discovery
Redis master and slave

1. Extra standalone service: Consul, ZooKeeper and Etcd
2. Serf package
    1. https://github.com/hashicorp/serf
    2. event handler on nodeJoin and nodeLeave, nodeFail

Replication: 
1. Pull
2. Push

Pattern:
1. Short job
    1. Run package: For CLI
2. Long running job 
    1. Agent package: For Server  
    2. Pattern used in Consul
<!--  -->

## Day8: Consensus algorithms
1. Raft
    1. used in Etcd, Consul and Kafka
    2. leader election
    3. Leader push to replication
    4. Cock roachDB Multi-Raft, each raft deal with a range

## Day9: Multiplex: run multiple services on the same port
package: cmux

## Day10: Load balancing
1. server proxy: ELB, GOOGLE LB, AZURE LB



GRPC resolver
default round-robin 
1. read and write
2. global, cdn

custom resolver and picker

resolver needs to know each server's address and leader

server expose a rpc endpoint to tell client all the servers

google.golang.org/grpc/resolver


grpc client connection will be passed to Resolver

resolver discorver servers and update the client connections

resolver will have a different client connection to grpc server to fetch the list of servers


pickers: picker a server from server lists

we can implement load balaencing algorithem throw pickers and custom routing logic



## Day11: Kubernetes Deployment
Statefulset

PersistentVolumeClaim
local disk

InitContainer:
busybox image
volumeMounts
based on the 


Container Probe and GRPC Health check

google grpc health package
grpc_health_v1
grpc_health_probe executable

we run a standalone heath grpc server


headless service: no loadbalancer for internal ip
each pod have own dns record


## Day12: Expose grpc to public internet

automatic lb for each pod (headless already)