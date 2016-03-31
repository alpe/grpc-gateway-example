# grpc demo
very much inspired by CoreOS Blog post: https://coreos.com/blog/gRPC-protobufs-swagger.html 
The implementation started with a simple http://grpc.io demo and then got [grpcgateway] integrated to talk REST to the grpc endpoint.

### Generate structs from protobuf
```
cd ./protos
make
```
### dependencies
https://github.com/gengo/grpc-gateway
