package internal

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ueckoken/plarail2022/backend/external/pkg/envStore"
	"github.com/ueckoken/plarail2022/backend/external/pkg/httphandler"
	"github.com/ueckoken/plarail2022/backend/external/pkg/synccontroller"
	"github.com/ueckoken/plarail2022/backend/external/pkg/websockethandler"
	"github.com/ueckoken/plarail2022/backend/external/spec"

	"go.uber.org/zap"
	"net/http/pprof"
)

// Run runs external server.
func Run(logger *zap.Logger) {
	ctx := context.Background()
	synccontrollerInput := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum], 64)
	synccontrollerOutput := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum], 64)
	grpcHandlerOutput := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum], 64)
	main2autooperation := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum], 64)
	main2internal := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum], 64)
	httpInputKV := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum], 64)
	httpInput := make(chan *spec.PointAndState, 64)
	httpOutput := make(chan *spec.PointAndState, 64)

	namespace := "internal"

	synccontrollerOutputTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "synccontroller_output_total",
			Help:      "",
		},
		[]string{},
	)

	httpInputKVTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "httpinputkv_total",
			Help:      "",
		},
		[]string{},
	)

	httpOutputTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "httpoutput_total",
			Help:      "",
		},
		[]string{},
	)

	grpcHandlerOutputTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "grpchandleroutput_total",
			Help:      "",
		},
		[]string{},
	)

	prometheus.MustRegister(synccontrollerOutputTotal)
	prometheus.MustRegister(httpInputKVTotal)
	prometheus.MustRegister(httpOutputTotal)
	prometheus.MustRegister(grpcHandlerOutputTotal)

	go func() {
		for c := range synccontrollerOutput {
			synccontrollerOutputTotal.With(prometheus.Labels{}).Inc()
			main2autooperation <- c
			httpInputKV <- c
			main2internal <- c
		}
	}()

	go func() {
		for c := range httpInputKV {
			httpInputKVTotal.With(prometheus.Labels{}).Inc()
			httpInput <- &spec.PointAndState{Station: &spec.Station{StationId: c.Key}, State: c.Value}
		}
	}()

	go func() {
		for c := range httpOutput {
			httpOutputTotal.With(prometheus.Labels{}).Inc()
			synccontrollerInput <- synccontroller.KV[spec.StationId, spec.PointStateEnum]{Key: c.GetStation().GetStationId(), Value: c.GetState()}
		}
	}()

	go func() {
		for c := range grpcHandlerOutput {
			grpcHandlerOutputTotal.With(prometheus.Labels{}).Inc()
			synccontrollerInput <- c
		}
	}()

	envVal := envStore.GetEnv()

	httpServer := httphandler.NewHTTPServer(
		logger.Named("http-state"),
		httpOutput,
		httpInput,
		envVal,
		"httppointserver",
	)

	StartStationSync(logger.Named("station-sync"), synccontrollerInput, synccontrollerOutput)
	grpcHandler := NewGrpcHandler(logger.Named("grpc-handler"), envVal, main2autooperation, grpcHandlerOutput)
	internalHandler := NewGrpcHandlerForInternal(logger.Named("grpc-internal"), envVal, main2internal)
	go internalHandler.Run(ctx)

	httpBlockInput := make(chan *spec.BlockAndState, 64)
	httpBlockOutput := make(chan *spec.BlockAndState, 64)
	blocksyncInput := make(chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum], 64)
	blocksyncOutput := make(chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum], 64)
	grpcBlockHandlerInput := make(chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum], 64)
	grpcBlockHandlerOutput := make(chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum], 64)

	httpBlockServer := httphandler.NewHTTPServer(
		logger.Named("http-block"),
		httpBlockOutput,
		httpBlockInput,
		envVal,
		"httpblockserver",
	)

	grpcBlockHandlerOutputTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "grpcblockhandleroutput_total",
			Help:      "",
		},
		[]string{},
	)

	blocksyncOutputTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "blocksyncoutput_total",
			Help:      "",
		},
		[]string{},
	)
	prometheus.MustRegister(grpcBlockHandlerOutputTotal)
	prometheus.MustRegister(blocksyncOutputTotal)

	go func() {
		for c := range grpcBlockHandlerOutput {
			grpcBlockHandlerOutputTotal.With(prometheus.Labels{}).Inc()
			blocksyncInput <- synccontroller.KV[spec.BlockId, spec.BlockStateEnum]{Key: c.Key, Value: c.Value}
		}
	}()

	go func() {
		for c := range blocksyncOutput {
			blocksyncOutputTotal.With(prometheus.Labels{}).Inc()
			grpcBlockHandlerInput <- c
			httpBlockInput <- &spec.BlockAndState{BlockId: c.Key, State: c.Value}
		}
	}()

	startBlockSync(logger.Named("blocksync"), blocksyncInput, blocksyncOutput)
	grpcBlockHandl := NewGrpcBlockHandler(logger.Named("grpc-block-handler"), envVal, grpcBlockHandlerInput, grpcBlockHandlerOutput)

	go GRPCListenAndServe(ctx, logger, uint(envVal.ClientSideServer.GrpcPort), grpcHandler, grpcBlockHandl)
	r := mux.NewRouter()
	r.HandleFunc("/", websockethandler.HandleStatic)
	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
	r.HandleFunc("/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	httpBlockServer.RegisterServer(r, "/blockws")
	httpServer.RegisterServer(r, "/pointws")

	srv := &http.Server{
		Handler:           r,
		Addr:              fmt.Sprintf("0.0.0.0:%d", envVal.ClientSideServer.Port),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	logger.Info("start listening")
	if err := srv.ListenAndServe(); err != nil {
		logger.Panic("failed to serve", zap.Error(err))
	}
}
