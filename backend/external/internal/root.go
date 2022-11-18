package internal

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ueckoken/plarail2022/backend/external/pkg/envStore"
	"github.com/ueckoken/plarail2022/backend/external/pkg/httphandler"
	"github.com/ueckoken/plarail2022/backend/external/pkg/synccontroller"
	"github.com/ueckoken/plarail2022/backend/external/pkg/websockethandler"
	"github.com/ueckoken/plarail2022/backend/external/spec"

	"go.uber.org/zap"
)

// Run runs external server.
func Run(logger *zap.Logger) {
	ctx := context.Background()
	synccontrollerInput := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum])
	synccontrollerOutput := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum], 32)
	grpcHandlerInput := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum])
	main2autooperation := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum])
	main2internal := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum], 32)
	httpInputKV := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum], 32)
	httpInput := make(chan *spec.PointAndState)
	httpOutput := make(chan *spec.PointAndState)

	go func() {
		for c := range synccontrollerOutput {
			select {
			case main2autooperation <- c:
			default:
				logger.Info("buffer full", zap.String("buffer", "main2grpcHandler"))
			}
			select {
			case httpInputKV <- c:
			default:
				logger.Info("buffer full", zap.String("buffer", "httpInputKV"))
			}
			select {
			case main2internal <- c:
			default:
				logger.Info("buffer full", zap.String("buffer", "internal"))
			}
		}
	}()

	go func() {
		for c := range httpInputKV {
			select {
			case httpInput <- &spec.PointAndState{Station: &spec.Station{StationId: c.Key}, State: c.Value}:
			default:
				logger.Info("buffer full", zap.String("buffer", "httpInput"))
			}
		}
	}()

	go func() {
		for c := range httpOutput {
			select {
			case synccontrollerInput <- synccontroller.KV[spec.StationId, spec.PointStateEnum]{Key: c.GetStation().GetStationId(), Value: c.GetState()}:
			default:
				logger.Info("buffer full", zap.String("buffer", "synccontrollerInput-httpoutput"))
			}
		}
	}()

	go func() {
		for c := range grpcHandlerInput {
			select {
			case synccontrollerInput <- c:
			default:
				logger.Info("buffer full", zap.String("buffer", "synccontrollerInput-grpchandler"))
			}
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
	grpcHandler := NewGrpcHandler(logger.Named("grpc-handler"), envVal, main2autooperation, grpcHandlerInput)
	internalHandler := NewGrpcHandlerForInternal(logger.Named("grpc-internal"), envVal, main2internal)
	go internalHandler.Run(ctx)

	httpBlockInput := make(chan *spec.BlockAndState)
	httpBlockOutput := make(chan *spec.BlockAndState)
	blocksyncInput := make(chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum])
	blocksyncOutput := make(chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum])
	grpcBlockHandlerInput := make(chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum])
	grpcBlockHandlerOutput := make(chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum])

	httpBlockServer := httphandler.NewHTTPServer(
		logger.Named("http-block"),
		httpBlockOutput,
		httpBlockInput,
		envVal,
		"httpblockserver",
	)

	go func() {
		for c := range grpcBlockHandlerOutput {
			select {
			case blocksyncInput <- synccontroller.KV[spec.BlockId, spec.BlockStateEnum]{Key: c.Key, Value: c.Value}:
			default:
				logger.Info("buffer full", zap.String("buffer", "blocksyncInput-grpchandler"))
			}
		}
	}()

	go func() {
		for c := range blocksyncOutput {
			select {
			case grpcBlockHandlerInput <- c:
			default:
				logger.Info("buffer full", zap.String("buffer", "grpcblockhandlerinput-blocksync"))
			}
			select {
			case httpBlockInput <- &spec.BlockAndState{BlockId: c.Key, State: c.Value}:
			default:
				logger.Info("buffer full", zap.String("buffer", "httpblockinput-blocksync"))
			}
		}
	}()

	startBlockSync(logger.Named("blocksync"), blocksyncInput, blocksyncOutput)
	grpcBlockHandl := NewGrpcBlockHandler(logger.Named("grpc-block-handler"), envVal, grpcBlockHandlerOutput, grpcBlockHandlerInput)

	go GRPCListenAndServe(ctx, logger, uint(envVal.ClientSideServer.GrpcPort), grpcHandler, grpcBlockHandl)
	r := mux.NewRouter()
	r.HandleFunc("/", websockethandler.HandleStatic)
	r.Handle("/metrics", promhttp.Handler())
	httpBlockServer.RegisterServer(r, "/blockws")
	httpServer.RegisterServer(r, "/pointws")

	srv := &http.Server{
		Handler:           r,
		Addr:              fmt.Sprintf("0.0.0.0:%d", int(envVal.ClientSideServer)),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	logger.Info("start listening")
	if err := srv.ListenAndServe(); err != nil {
		logger.Panic("failed to serve", zap.Error(err))
	}
}
