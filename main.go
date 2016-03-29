package main

import (
	pb "github.com/alpe/grpc-gateway-example/protos"
	"golang.org/x/net/context"
	"log"
	"net"

	"google.golang.org/grpc"
)




//go:generate protoc -I protos -I$GOPATH/src protos/service.proto --go_out=plugins=grpc:protos


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