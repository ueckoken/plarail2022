package serveGrpc

import (
	"context"
	"fmt"
	"github.com/ueckoken/plarail2022/backend/internal/pkg/msg2Esp"
	"github.com/ueckoken/plarail2022/backend/internal/pkg/station2espIp"
	pb "github.com/ueckoken/plarail2022/backend/internal/spec"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ControlServer struct {
	pb.UnimplementedPointStateNotificationServer
	Stations station2espIp.Stations
	client   *http.Client
}

func (c *ControlServer) NotifyPointState(_ context.Context, req *pb.NotifyPointStateRequest) (*pb.NotifyPointStateResponse, error) {
	sta, err := c.unpackStations(req.GetState())
	if err != nil {
		return nil, err
	}
	angle := 0
	if sta.IsAngleDefined() {
		angle, err = sta.GetAngle(req.GetState().GetState())
		if err != nil {
			return nil, err
		}
	}
	s2n := msg2Esp.NewSend2Node(c.client, sta, c.unpackState(req.GetState().GetState()), angle)
	err = s2n.Send()

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "sender err %s; not connected to Node", err.Error())
	}
	return &pb.NotifyPointStateResponse{}, nil
}

func (c *ControlServer) unpackStations(req *pb.PointAndState) (*station2espIp.StationDetail, error) {
	s, ok := pb.StationId_name[int32(req.GetStation().GetStationId())]
	if !ok {
		return nil, fmt.Errorf("station: %s do not define in proto file", req.String())
	}
	sta, err := c.Stations.Detail(s)
	if err != nil {
		return nil, fmt.Errorf("%w; station `%s` is not defined in yaml file", err, s)
	}
	return sta, nil
}

func (*ControlServer) unpackState(state pb.PointStateEnum) string {
	return state.String()
}
