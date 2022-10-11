package main

import (
	"fmt"
	"log"
	"net"
	"ueckoken/plarail2021-soft-internal/internal"
	"ueckoken/plarail2021-soft-internal/pkg/grpcMock"
	"ueckoken/plarail2021-soft-internal/pkg/serveGrpc"
	pb "ueckoken/plarail2021-soft-internal/spec"

	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcPrometheus.UnaryServerInterceptor),
	)
	e := internal.GetEnv()
	m := grpcMock.GrpcMock{}
	pb.RegisterControlServer(s, &m)
	grpcPrometheus.Register(s)
	serveGrpc.ServeMetrics(fmt.Sprintf(":%d", e.ExternalSideServer.MetricsPort))
	fmt.Printf("gRPC bind address :%d\n", e.ExternalSideServer.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", e.ExternalSideServer.Port))
	if err != nil {
		log.Fatalf("listen failed")
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
