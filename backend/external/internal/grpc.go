package internal

import (
	"context"
	"fmt"
	"net"
	"ueckoken/plarail2022-external/pkg/envStore"
	"ueckoken/plarail2022-external/pkg/synccontroller"
	"ueckoken/plarail2022-external/spec"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcStateHandler is a handler for gRPC.
type GrpcStateHandler struct {
	logger *zap.Logger
	env    *envStore.Env
	spec.UnimplementedControlServer
	stateOutput chan<- synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]
	stateInput  <-chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]
}

// NewGrpcHandler creates gRPC handler.
func NewGrpcHandler(logger *zap.Logger, env *envStore.Env, stateOutput chan<- synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State], stateInput <-chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]) *GrpcStateHandler {
	return &GrpcStateHandler{logger: logger, env: env, stateOutput: stateOutput, stateInput: stateInput}
}

// Command2Internal handles requests from ATS.
func (g GrpcStateHandler) Command2Internal(_ context.Context, req *spec.RequestSync) (*spec.ResponseSync, error) {
	s := synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]{
		Key:   req.GetStation().GetStationId(),
		Value: spec.Command2InternalRequest_State(req.GetState()),
	}
	g.stateOutput <- s
	return &spec.ResponseSync{Response: spec.ResponseSync_SUCCESS}, nil
}

// / handleInput transmits changes received in channel to ATS.
func (g GrpcStateHandler) handleInput(ctx context.Context) {
	con, err := grpc.DialContext(ctx, g.env.ClientSideServer.ATSAddress.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		g.logger.Error("failed to connect ATS", zap.Error(err))
	}
	defer con.Close()
	for d := range g.stateInput {
		client := spec.NewControlClient(con)
		req := &spec.RequestSync{
			Station: &spec.Stations{StationId: d.Key},
			State:   spec.RequestSync_State(d.Value),
		}
		res, err := client.Command2Internal(ctx, req)
		if err != nil {
			g.logger.Error("failed to send data to ATS", zap.Any("payload", req), zap.Error(err))
		}
		if res.GetResponse() != spec.ResponseSync_SUCCESS {
			g.logger.Error("ATS response seems to be unsuccessfull", zap.Any("payload", res.GetResponse()))
		}
	}
}

type GrpcBlockHandler struct {
	env    *envStore.Env
	logger *zap.Logger
	spec.UnimplementedBlockStateSyncServer
	stateOutput chan<- synccontroller.KV[spec.Blocks_BlockId, spec.NotifyStateRequest_State]
	stateInput  <-chan synccontroller.KV[spec.Blocks_BlockId, spec.NotifyStateRequest_State]
}

// NewGrpcBlockHandler creates gRPC handler.
func NewGrpcBlockHandler(logger *zap.Logger, env *envStore.Env, stateOutput chan<- synccontroller.KV[spec.Blocks_BlockId, spec.NotifyStateRequest_State], stateInput <-chan synccontroller.KV[spec.Blocks_BlockId, spec.NotifyStateRequest_State]) *GrpcBlockHandler {
	return &GrpcBlockHandler{logger: logger, env: env, stateOutput: stateOutput, stateInput: stateInput}
}

// NotifyState handles requests from ATS.
func (g GrpcBlockHandler) NotifyState(_ context.Context, req *spec.NotifyStateRequest) (*spec.NotifyStateResponse, error) {
	g.stateOutput <- synccontroller.KV[spec.Blocks_BlockId, spec.NotifyStateRequest_State]{Key: req.GetBlock().GetBlockId(), Value: req.GetState()}
	return &spec.NotifyStateResponse{Response: spec.NotifyStateResponse_SUCCESS}, nil
}

// handleInput transmits changes received in channel to ATS.
func (g GrpcBlockHandler) handleInput(ctx context.Context) {
	con, err := grpc.DialContext(ctx, g.env.ClientSideServer.ATSAddress.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		g.logger.Error("failed to connect ATS", zap.Error(err))
	}
	defer con.Close()
	for d := range g.stateInput {
		client := spec.NewBlockStateSyncClient(con)
		req := &spec.NotifyStateRequest{
			Block: &spec.Blocks{BlockId: d.Key},
			State: d.Value,
		}
		res, err := client.NotifyState(ctx, req)
		if err != nil {
			g.logger.Error("failed to send data to ATS", zap.Any("payload", req), zap.Error(err))
		}
		if res.GetResponse() != spec.NotifyStateResponse_SUCCESS {
			g.logger.Error("ATS response seems to be unsuccessfull", zap.Any("payload", res.GetResponse()))
		}
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
	spec.RegisterControlServer(s, handler)
	spec.RegisterBlockStateSyncServer(s, blockhandler)
	if err := s.Serve(lis); err != nil {
		logger.Panic("failed to server", zap.Error(err))
	}
}
