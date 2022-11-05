package main

import (
	"context"
	"fmt"
	atspb "json2grpc/spec"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

type myServer struct {
	atspb.UnimplementedAtsServer
}

func (s *myServer) SendStatus(ctx context.Context, req *atspb.SendStatusRequest) (*atspb.SendStatusResponse, error) {
	fmt.Println(req.GetSensor().String())
	return &atspb.SendStatusResponse{}, nil
}

func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	port := 8888
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	atspb.RegisterAtsServer(s, NewMyServer())
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
