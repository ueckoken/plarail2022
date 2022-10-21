package serveGrpc

import (
	"fmt"
	"testing"
	"ueckoken/plarail2022-internal/pkg/station2espIp"
	pb "ueckoken/plarail2022-internal/spec"
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
		UnimplementedControlServer pb.UnimplementedControlServer
		Stations                   station2espIp.Stations
	}
	type args struct {
		state pb.RequestSync_State
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "state is on",
			args: args{state: pb.RequestSync_ON},
			want: "ON",
		},
		{
			name: "state is off",
			args: args{state: pb.RequestSync_OFF},
			want: "OFF",
		},
		{
			name: "state is unknown",
			args: args{state: pb.RequestSync_UNKNOWN},
			want: "UNKNOWN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ControlServer{
				UnimplementedControlServer: tt.fields.UnimplementedControlServer,
				Stations:                   tt.fields.Stations,
			}
			if got := c.unpackState(tt.args.state); got != tt.want {
				t.Errorf("unpackState() = %v, want %v", got, tt.want)
			}
		})
	}
}
