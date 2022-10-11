package serveGrpc

import (
	"fmt"
	"reflect"
	"testing"
	"ueckoken/plarail2021-soft-internal/pkg/station2espIp"
	pb "ueckoken/plarail2021-soft-internal/spec"
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
func TestControlServer_unpackStations(t *testing.T) {
	type fields struct {
		UnimplementedControlServer pb.UnimplementedControlServer
		Stations                   station2espIp.Stations
	}
	f := fields{
		Stations: &TestStations{
			Stations: []station2espIp.Station{
				{
					Station: station2espIp.StationDetail{
						Name:    "chofu_b1",
						Address: "TEST_ADDR",
						Pin:     1,
					},
				}, {
					Station: station2espIp.StationDetail{
						Name:    "chofu_b2",
						Address: "TEST_ADDR",
						Pin:     2,
					},
				},
				{
					Station: station2espIp.StationDetail{
						Name:    "TOKYO",
						Address: "TEST_ADDR",
						Pin:     1,
					},
				}, {
					Station: station2espIp.StationDetail{
						Name:    "chofu_b2",
						Address: "TEST_ADDR",
						Pin:     2,
					},
				},
			}}}
	type args struct {
		req *pb.Stations
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *station2espIp.StationDetail
		wantErr bool
	}{
		{
			name:   "station exist",
			fields: f,
			args:   args{req: &pb.Stations{StationId: pb.Stations_chofu_b1}},
			want: &station2espIp.StationDetail{
				Name:    "chofu_b1",
				Address: "TEST_ADDR",
				Pin:     1,
			},
			wantErr: false,
		},
		{
			name:    "station not define in yaml",
			fields:  f,
			args:    args{req: &pb.Stations{StationId: pb.Stations_sasazuka_b1}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ControlServer{
				UnimplementedControlServer: tt.fields.UnimplementedControlServer,
				Stations:                   tt.fields.Stations,
			}
			got, err := c.unpackStations(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpackStations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unpackStations() got = %v, want %v", got, tt.want)
			}
		})
	}
}
