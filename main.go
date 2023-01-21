package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/mattgonewild/di"
	"google.golang.org/grpc"
)

type server struct {
	di.UnimplementedDataLayerServer
}

func (s *server) CreateImg(_ context.Context, _ *di.CreateImgRequest) (*di.CreateImgResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) ReadImg(_ context.Context, _ *di.ReadImgRequest) (*di.ReadImgResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) UpdateImg(_ context.Context, _ *di.UpdateImgRequest) (*di.UpdateImgResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DeleteImg(_ context.Context, _ *di.DeleteImgRequest) (*di.DeleteImgResponse, error) {
	panic("not implemented") // TODO: Implement
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	di.RegisterDataLayerServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
