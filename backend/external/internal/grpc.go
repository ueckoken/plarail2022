package internal

import (
	"context"
	"fmt"
	"log"
	"net"
	"ueckoken/plarail2022-external/pkg/envStore"
	"ueckoken/plarail2022-external/pkg/synccontroller"
	"ueckoken/plarail2022-external/spec"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcHandler is a handler for gRPC.
type GrpcHandler struct {
	env *envStore.Env
	spec.UnimplementedControlServer
	stateOutput chan<- synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]
	stateInput  <-chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]
}

// NewGrpcHandler creates gRPC handler.
func NewGrpcHandler(env *envStore.Env, stateOutput chan<- synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State], stateInput <-chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]) *GrpcHandler {
	return &GrpcHandler{env: env, stateOutput: stateOutput, stateInput: stateInput}
}

// Command2Internal handles requests and save requests to memory and treats internal server.
func (g GrpcHandler) Command2Internal(ctx context.Context, req *spec.RequestSync) (*spec.ResponseSync, error) {
	s := synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]{
		Key:   req.GetStation().GetStationId(),
		Value: spec.Command2InternalRequest_State(req.GetState()),
	}
	g.stateOutput <- s
	return &spec.ResponseSync{Response: spec.ResponseSync_SUCCESS}, nil
}

func (g GrpcHandler) handleInput(ctx context.Context) {
	con, err := grpc.Dial(g.env.ClientSideServer.ATSAddress.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("failed to connect ATS", err)
	}
	defer con.Close()
	for d := range g.stateInput {
		client := spec.NewControlClient(con)
		res, err := client.Command2Internal(ctx,
			&spec.RequestSync{
				Station: &spec.Stations{StationId: d.Key},
				State:   spec.RequestSync_State(d.Value),
			})
		if err != nil {
			log.Println("failed to send data to ATS", err)
		}
		if res.GetResponse() != spec.ResponseSync_SUCCESS {
			log.Println("ATS response unsuccessfull", res.GetResponse())
		}
	}
}

// GRPCListenAndServe listens and serve.
func GRPCListenAndServe(port uint, handler *GrpcHandler) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("failed to listen port", err)
	}
	go handler.handleInput(context.Background())
	s := grpc.NewServer()
	spec.RegisterControlServer(s, handler)
	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
	}
}
