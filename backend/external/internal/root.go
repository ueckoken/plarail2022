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
	synccontrollerInput := make(chan synccontroller.KV[spec.StationId, spec.State])
	synccontrollerOutput := make(chan synccontroller.KV[spec.StationId, spec.State])
	grpcHandlerInput := make(chan synccontroller.KV[spec.StationId, spec.State])
	main2grpcHandler := make(chan synccontroller.KV[spec.StationId, spec.State])
	httpInputKV := make(chan synccontroller.KV[spec.StationId, spec.State])
	httpInput := make(chan *spec.PointAndState)
	httpOutput := make(chan *spec.PointAndState)

	go func() {
		for c := range synccontrollerOutput {
			select {
			case main2grpcHandler <- c:
			default:
				logger.Info("buffer full", zap.String("buffer", "main2grpcHandler"))
			}
			select {
			case httpInputKV <- c:
			default:
				logger.Info("buffer full", zap.String("buffer", "httpInputKV"))
			}
		}
	}()
	go func() {
		for c := range httpInputKV {
			httpInput <- &spec.PointAndState{Station: &spec.Station{StationId: c.Key}, State: c.Value}
		}
	}()

	go func() {
		for c := range httpOutput {
			synccontrollerInput <- synccontroller.KV[spec.StationId, spec.State]{Key: c.GetStation().GetStationId(), Value: c.GetState()}
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
	grpcHandler := NewGrpcHandler(logger.Named("grpc-handler"), envVal, main2grpcHandler, grpcHandlerInput)

	client2blocksync := make(chan synccontroller.KV[spec.Blocks_BlockId, spec.NotifyStateRequest_State])
	blocksync2client := make(chan synccontroller.KV[spec.Blocks_BlockId, spec.NotifyStateRequest_State])
	startBlockSync(logger.Named("blocksync"), client2blocksync, blocksync2client)
	grpcBlockHandl := NewGrpcBlockHandler(logger.Named("grpc-block-handler"), envVal, client2blocksync, blocksync2client)

	go GRPCListenAndServe(ctx, logger, uint(envVal.ClientSideServer.GrpcPort), grpcHandler, grpcBlockHandl)
	httpServer.StartServer()
}
