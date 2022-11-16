package servo

import (
	"context"
	"fmt"
	"time"

	"github.com/ueckoken/plarail2022/backend/external/pkg/envStore"
	"github.com/ueckoken/plarail2022/backend/external/spec"
	pb "github.com/ueckoken/plarail2022/backend/external/spec"

	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	UNKNOWN = "UNKNOWN"
	SUCCESS = "SUCCESS"
	FAILED  = "FAILED"
)

type Command2Internal struct {
	station *spec.PointAndState
	env     *envStore.Env
}

// NewCommand2Internal is Constructor of CommandInternal.
// CommandInternal Struct has a method to talk to Internal server with gRPC.
func NewCommand2Internal(state *spec.PointAndState, e *envStore.Env) *Command2Internal {
	return &Command2Internal{station: state, env: e}
}

// sendRaw is making a connection to internal server and talk with internal server.
// This method will return gRPC response and gRPC error val.
// If you want join gRPC response Status Code and gRPC error msg, please use Command2Internal.trapResponseGrpcErr method.
func (c2i *Command2Internal) sendRaw() (*pb.NotifyPointStateResponse, error) {
	// Set up a connection to the server.
	ctx, cancel1 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel1()
	conn, err := grpc.DialContext(
		ctx,
		c2i.env.InternalServer.Addr.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(grpcPrometheus.UnaryClientInterceptor),
	)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewNotificationClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(ctx, c2i.env.InternalServer.TimeoutSec)
	defer cancel()
	r, err := c.NotifyPointState(ctx, c2i.convert2pb())
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c2i *Command2Internal) Send() error {
	return trapResponseGrpcErr(c2i.sendRaw())
}

func (c2i *Command2Internal) convert2pb() *pb.NotifyPointStateRequest {
	return &pb.NotifyPointStateRequest{
		State:   &pb.PointAndState{Station: c2i.station.GetStation(), State: c2i.station.GetState()},
	}
}
func (c2i *Command2Internal) String() string {
	return fmt.Sprintf("%s -> %s", c2i.station.GetStation(), c2i.station.GetState())
}
func trapResponseGrpcErr(rs *pb.NotifyPointStateResponse, grpcErr error) error {
	// From Error will return true in ok if err is occurred by gRPC or nil
	sta, ok := status.FromError(grpcErr)
	if (sta != nil && ok) || rs == nil { // gRPC error occur
		return fmt.Errorf("gRPC Err: `%w`", grpcErr)
	}
	// check Response Status
	switch rs.Response.String() {
	case UNKNOWN:
		return fmt.Errorf("gRPC Err: `%w`; gRPC Response status is `%s`", grpcErr, UNKNOWN)
	case SUCCESS:
		if grpcErr != nil {
			return fmt.Errorf("gRPC Err: `%w`; gRPC Response status is `%s`", grpcErr, SUCCESS)
		}
		return nil
	case FAILED:
		return fmt.Errorf("gRPC Err: `%w`; gRPC Response status is `%s`", grpcErr, FAILED)
	default:
		return fmt.Errorf("gRPC Err: `%w`; Unknown error is occured", grpcErr)
	}
}
