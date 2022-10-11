package main

import (
	"ueckoken/plarail2021-soft-speed/internal"
	"ueckoken/plarail2021-soft-speed/pkg/train2IP"

	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "plarail2021_speed"

func main() {
	clientConn := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "clients_connections_seconds",
			Help:      "Number of connections handling websocket",
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

	speed := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "speed",
			Help:      "speed",
		},
		[]string{},
	)

	trainInfo := train2IP.GetTable()
	handler := internal.ClientHandler{Clients: &internal.ClientSet{}, TrainName2Id: trainInfo}
	server := internal.SpeedServer{ClientHandler: &handler, NumberOfClientConnection: clientConn, TotalCLientCommands: controlCommandTotal, Speed: speed}
	server.StartSpeed()
}
