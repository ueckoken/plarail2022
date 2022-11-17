package internal

import (
	"context"
	"github.com/ueckoken/plarail2022/backend/external/pkg/envStore"
	"github.com/ueckoken/plarail2022/backend/external/pkg/httphandler"
	"github.com/ueckoken/plarail2022/backend/external/pkg/synccontroller"
	"github.com/ueckoken/plarail2022/backend/external/spec"

	"go.uber.org/zap"
)

// Run runs external server.
func Run(logger *zap.Logger) {
	ctx := context.Background()
	synccontrollerInput := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum])
	synccontrollerOutput := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum])
	grpcHandlerInput := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum])
	main2autooperation := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum])
	main2internal := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum])
	httpInputKV := make(chan synccontroller.KV[spec.StationId, spec.PointStateEnum])
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

	go func(){
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
		logger.Named("sync-controller"),
		httpOutput,
		httpInput,
		envVal,
	)
	StartStationSync(logger.Named("station-sync"), synccontrollerInput, synccontrollerOutput)
	grpcHandler := NewGrpcHandler(logger.Named("grpc-handler"), envVal, main2autooperation, grpcHandlerInput)
	internalHandler := NewGrpcHandlerForInternal(logger.Named("grpc-internal"), envVal, main2internal)
	go internalHandler.Run(ctx)

	client2blocksync := make(chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum])
	blocksync2client := make(chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum])
	startBlockSync(logger.Named("blocksync"), client2blocksync, blocksync2client)
	grpcBlockHandl := NewGrpcBlockHandler(logger.Named("grpc-block-handler"), envVal, client2blocksync, blocksync2client)

	go GRPCListenAndServe(ctx, logger, uint(envVal.ClientSideServer.GrpcPort), grpcHandler, grpcBlockHandl)
	httpServer.StartServer()
}
