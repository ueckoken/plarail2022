package serveGrpc

import (
	"fmt"
	"github.com/ueckoken/plarail2022/backend/internal/pkg/station2espIp"
	pb "github.com/ueckoken/plarail2022/backend/internal/spec"
	"testing"
)

type TestStations struct {
	Stations []station2espIp.Station
}

func (t *TestStations) Detail(name string) (*station2espIp.StationDetail, error) {
	for _, s := range t.Stations {
		if s.Station.Name == name {
			return &s.Station, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (t *TestStations) Enumerate() []station2espIp.Station {
	return t.Stations
}

func TestControlServer_unpackState(t *testing.T) {
	type fields struct {
		UnimplementedNotificationServer pb.UnimplementedPointStateNotificationServer
		Stations                        station2espIp.Stations
	}
	type args struct {
		state pb.PointStateEnum
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "state is on",
			args: args{state: pb.PointStateEnum_POINTSTATE_ON},
			want: "ON",
		},
		{
			name: "state is off",
			args: args{state: pb.PointStateEnum_POINTSTATE_OFF},
			want: "OFF",
		},
		{
			name: "state is unknown",
			args: args{state: pb.PointStateEnum_POINTSTATE_UNKNOWN},
			want: "UNKNOWN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ControlServer{
				UnimplementedPointStateNotificationServer: tt.fields.UnimplementedNotificationServer,
				Stations: tt.fields.Stations,
			}
			if got := c.unpackState(tt.args.state); got != tt.want {
				t.Errorf("unpackState() = %v, want %v", got, tt.want)
			}
		})
	}
}
