package main

import (
	"ueckoken/plarail2022-external/internal"
	"ueckoken/plarail2022-external/pkg/envStore"
	"ueckoken/plarail2022-external/pkg/syncController"

	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "plarailexternal"

func main() {
	syncControllerInput := make(chan syncController.StationState)
	grpcHandlerInput := make(chan syncController.StationState)
	main2grpcHandler := make(chan syncController.StationState)
	syncControllerOutput := make(chan syncController.StationState)
	httpInput := make(chan syncController.StationState)

	go func() {
		for c := range syncControllerOutput {
			main2grpcHandler <- c
			httpInput <- c
		}
	}()

	envVal := envStore.GetEnv()

	clientConn := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "clients_connections_seconds",
			Help:      "Number of connections handling websocket",
		},
		[]string{},
	)

	clientConnTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "clients_connections_total",
			Help:      "Total client connection",
		},
		[]string{},
	)

	controlCommandTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "client_commands_total",
			Help:      "Total client commands",
		},
		[]string{},
	)

	httpServer := internal.HTTPServer{
		StateOutput:              syncControllerInput,
		StateInput:               httpInput,
		Environment:              envVal,
		NumberOfClientConnection: clientConn,
		TotalClientConnection:    clientConnTotal,
		TotalCLientCommands:      controlCommandTotal,
		Clients:                  &internal.ClientsCollection{},
	}
	syncController := syncController.SyncController{
		StateInput:  syncControllerInput,
		StateOutput: syncControllerOutput,
		Environment: envVal,
	}
	grpcHandler := internal.NewGrpcHandler(envVal, syncControllerInput, grpcHandlerInput)

	go internal.GRPCListenAndServe(uint(envVal.ClientSideServer.GrpcPort), grpcHandler)
	go httpServer.StartServer()
	syncController.StartSyncController()
}
