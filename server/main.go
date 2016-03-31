package main

import (
	pb "github.com/alpe/grpc-gateway-example/protos"
	"golang.org/x/net/context"
	"log"
	"net"

	"google.golang.org/grpc"
)




//go:generate protoc \
//  -I/usr/local/include \
//  -I protos \
//  -I$GOPATH/src \
//  -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis \
//  --go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:protos \
//   protos/service.proto

// //go:generate protoc -I/usr/local/include -I protos -I$GOPATH/src -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:protos protos/service.proto

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}