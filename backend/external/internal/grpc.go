package internal

import (
	"context"
	"fmt"
	"net"

	"github.com/ueckoken/plarail2022/backend/external/pkg/envStore"
	"github.com/ueckoken/plarail2022/backend/external/pkg/synccontroller"
	"github.com/ueckoken/plarail2022/backend/external/spec"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcStateHandler is a handler for gRPC.
type GrpcStateHandler struct {
	logger *zap.Logger
	env    *envStore.Env
	spec.UnimplementedStateManagerServer
	stateOutput chan<- synccontroller.KV[spec.StationId, spec.PointStateEnum]
	stateInput  <-chan synccontroller.KV[spec.StationId, spec.PointStateEnum]
}

// NewGrpcHandler creates gRPC handler that handles requests from ATS.
func NewGrpcHandler(logger *zap.Logger, env *envStore.Env, stateInput <-chan synccontroller.KV[spec.StationId, spec.PointStateEnum], stateOutput chan<- synccontroller.KV[spec.StationId, spec.PointStateEnum]) *GrpcStateHandler {
	return &GrpcStateHandler{logger: logger, env: env, stateOutput: stateOutput, stateInput: stateInput}
}

// Command2Internal handles requests from ATS.
func (g GrpcStateHandler) UpdatePointState(_ context.Context, req *spec.UpdatePointStateRequest) (*spec.UpdatePointStateResponse, error) {
	s := synccontroller.KV[spec.StationId, spec.PointStateEnum]{
		Key:   req.GetState().GetStation().GetStationId(),
		Value: req.GetState().GetState(),
	}
	g.stateOutput <- s
	return &spec.UpdatePointStateResponse{}, nil
}

// / handleInput transmits changes received in channel to ATS.
func (g GrpcStateHandler) handleInput(ctx context.Context) {
	con, err := grpc.DialContext(ctx, g.env.ClientSideServer.ATSAddress.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		g.logger.Error("failed to connect to ATS", zap.Error(err))
	}
	defer con.Close()
	for d := range g.stateInput {
		go func(d synccontroller.KV[spec.StationId, spec.PointStateEnum]) {
			client := spec.NewPointStateNotificationClient(con)
			req := &spec.NotifyPointStateRequest{
				State: &spec.PointAndState{
					Station: &spec.Station{StationId: d.Key},
					State:   d.Value,
				},
			}
			_, err := client.NotifyPointState(ctx, req)
			if err != nil {
				g.logger.Error("failed to send data to ATS", zap.Any("payload", req), zap.Error(err))
			}
		}(d)
	}
}

type GrpcBlockHandler struct {
	env    *envStore.Env
	logger *zap.Logger
	spec.UnimplementedBlockStateManagerServer
	stateOutput chan<- synccontroller.KV[spec.BlockId, spec.BlockStateEnum]
	stateInput  <-chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum]
}

// NewGrpcBlockHandler creates gRPC handler.
func NewGrpcBlockHandler(logger *zap.Logger, env *envStore.Env, stateInput <-chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum], stateOutput chan<- synccontroller.KV[spec.BlockId, spec.BlockStateEnum]) *GrpcBlockHandler {
	return &GrpcBlockHandler{logger: logger, env: env, stateOutput: stateOutput, stateInput: stateInput}
}

// UpdateBlockState handles requests from ATS.
func (g GrpcBlockHandler) UpdateBlockState(_ context.Context, req *spec.UpdateBlockStateRequest) (*spec.UpdateBlockStateResponse, error) {
	g.stateOutput <- synccontroller.KV[spec.BlockId, spec.BlockStateEnum]{Key: req.GetState().GetBlockId(), Value: req.GetState().GetState()}
	return &spec.UpdateBlockStateResponse{}, nil
}

// handleInput transmits changes received in channel to ATS.
func (g GrpcBlockHandler) handleInput(ctx context.Context) {
	con, err := grpc.DialContext(ctx, g.env.ClientSideServer.ATSAddress.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		g.logger.Error("failed to connect to ATS", zap.Error(err))
	}
	defer con.Close()
	for d := range g.stateInput {
		go func(d synccontroller.KV[spec.BlockId, spec.BlockStateEnum]) {
			client := spec.NewBlockStateNotificationClient(con)
			req := &spec.NotifyBlockStateRequest{
				State: &spec.BlockAndState{BlockId: d.Key, State: d.Value},
			}
			_, err := client.NotifyBlockState(ctx, req)
			if err != nil {
				g.logger.Error("failed to send data to ATS", zap.Any("payload", req), zap.Error(err))
			}
		}(d)
	}
}

// GRPCListenAndServe listens and serve.
func GRPCListenAndServe(ctx context.Context, logger *zap.Logger, port uint, handler *GrpcStateHandler, blockhandler *GrpcBlockHandler) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Panic("failed to listen", zap.Error(err))
	}
	go handler.handleInput(ctx)
	go blockhandler.handleInput(ctx)
	s := grpc.NewServer()
	spec.RegisterStateManagerServer(s, handler)
	spec.RegisterBlockStateManagerServer(s, blockhandler)
	if err := s.Serve(lis); err != nil {
		logger.Panic("failed to server", zap.Error(err))
	}
}

// GrpcStateHandler is a handler for gRPC.
type GrpcStateHandlerForInternal struct {
	logger     *zap.Logger
	env        *envStore.Env
	stateInput <-chan synccontroller.KV[spec.StationId, spec.PointStateEnum]
}

// / handleInput transmits changes received in channel to ATS.
func (g GrpcStateHandlerForInternal) handleInput(ctx context.Context) {
	con, err := grpc.DialContext(ctx, g.env.InternalServer.Addr.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		g.logger.Error("failed to connect to internal", zap.Error(err))
	}
	defer con.Close()
	client := spec.NewPointStateNotificationClient(con)
	for d := range g.stateInput {
		go func(d synccontroller.KV[spec.StationId, spec.PointStateEnum]) {
			req := &spec.NotifyPointStateRequest{
				State: &spec.PointAndState{
					Station: &spec.Station{StationId: d.Key},
					State:   d.Value,
				},
			}
			_, err := client.NotifyPointState(ctx, req)
			if err != nil {
				g.logger.Error("failed to send data to ATS", zap.Any("payload", req), zap.Error(err))
			}
		}(d)
	}
}

func NewGrpcHandlerForInternal(logger *zap.Logger, env *envStore.Env, stateInput <-chan synccontroller.KV[spec.StationId, spec.PointStateEnum]) *GrpcStateHandlerForInternal {
	return &GrpcStateHandlerForInternal{logger, env, stateInput}
}

func (g GrpcStateHandlerForInternal) Run(ctx context.Context) {
	g.handleInput(ctx)
}
