package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"github.com/gengo/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "github.com/alpe/grpc-gateway-example/protos"
)

var (
	echoEndpoint = flag.String("echo_endpoint", ":50051", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	http.ListenAndServe(":8080", mux)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
