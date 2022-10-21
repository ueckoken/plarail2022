package main

import (
	"log"
	"ueckoken/plarail2022-internal/internal"
	"ueckoken/plarail2022-internal/pkg/esp32healthcheck"
	"ueckoken/plarail2022-internal/pkg/serveGrpc"
	"ueckoken/plarail2022-internal/pkg/station2espIp"

	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "plarailinternal"

func main() {
	env := internal.GetEnv()
	stations, err := station2espIp.NewStations()
	if err != nil {
		log.Fatalln(err)
	}
	esp32HealthCheck := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "esp32_health_seconds",
			Help:      "Esp32 health, if up, then 1",
		},
		[]string{"esp32Addr"},
	)

	pingHandler := esp32healthcheck.PingHandler{
		Stations:         *stations,
		Esp32HealthCheck: esp32HealthCheck,
	}
	grpcServer := serveGrpc.GrpcServer{
		Stations:    *stations,
		Environment: env,
		PingHandler: pingHandler,
	}
	grpcServer.StartServer()
}
