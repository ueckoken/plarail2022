package main

import (
	"ueckoken/plarail2022-external/internal"
	"ueckoken/plarail2022-external/pkg/envStore"
	"ueckoken/plarail2022-external/pkg/syncController"

	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "plarailexternal"

func main() {
	clientHandler2syncController := make(chan syncController.StationState, 16)
	syncController2clientHandler := make(chan syncController.StationState, 64)
	initEspStatus2syncController := make(chan syncController.StationState)

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
		ClientHandler2syncController: clientHandler2syncController,
		SyncController2clientHandler: syncController2clientHandler,
		Environment:                  envVal,
		NumberOfClientConnection:     clientConn,
		TotalClientConnection:        clientConnTotal,
		TotalCLientCommands:          controlCommandTotal,
		Clients:                      &internal.ClientsCollection{},
	}
	syncController := syncController.SyncController{
		ClientHandler2syncController: clientHandler2syncController,
		SyncController2clientHandler: syncController2clientHandler,
		Environment:                  envVal,
		InitServoRoute:               initEspStatus2syncController,
	}
	go httpServer.StartServer()
	syncController.StartSyncController()
}
