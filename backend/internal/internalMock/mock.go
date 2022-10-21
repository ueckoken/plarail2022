package main

import (
	"fmt"
	"log"
	"net"
	"ueckoken/plarail2022-internal/internal"
	"ueckoken/plarail2022-internal/pkg/serveGrpc"

	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcPrometheus.UnaryServerInterceptor),
	)
	e := internal.GetEnv()
	grpcPrometheus.Register(s)
	serveGrpc.ServeMetrics(fmt.Sprintf(":%d", e.ExternalSideServer.MetricsPort))
	log.Printf("gRPC bind address :%d\n", e.ExternalSideServer.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", e.ExternalSideServer.Port))
	if err != nil {
		log.Fatalf("listen failed")
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
