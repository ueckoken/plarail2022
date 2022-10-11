package station2espIp

import (
	_ "embed"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"
	pb "ueckoken/plarail2021-soft-internal/spec"
)

type Stations interface {
	Detail(name string) (*StationDetail, error)
	Enumerate() []Station
}
type stations struct {
	Stations []Station `yaml:"stations"`
}
type Station struct {
	Station StationDetail `yaml:"station"`
}
type StationDetail struct {
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Pin      int    `yaml:"pin"`
	SetAngle bool   `yaml:"set_angle"`
	OnAngle  int    `yaml:"on_angle"`
	OffAngle int    `yaml:"off_angle"`
}

//go:embed embed/station2espIp.yml
var ConfigFile []byte

func NewStations() (*stations, error) {
	t := stations{}
	err := yaml.Unmarshal(ConfigFile, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *stations) Detail(name string) (*StationDetail, error) {
	for _, sta := range s.Stations {
		if name == sta.Station.Name {
			return &sta.Station, nil
		}
	}
	return nil, fmt.Errorf("station %s not found", name)
}
func (d *StationDetail) IsAngleDefined() bool {
	return d.SetAngle
}

func (d *StationDetail) GetAngle(state pb.RequestSync_State) (angle int, err error) {
	if !d.IsAngleDefined() {
		return 0, fmt.Errorf("angles are not defined")
	}
	switch state {
	case pb.RequestSync_ON:
		angle = d.OnAngle
	case pb.RequestSync_OFF:
		angle = d.OffAngle
	default:
		return 0, status.Errorf(codes.InvalidArgument, "state is not ON or OFF\n")
	}
	return angle, nil
}

func (s *stations) Enumerate() []Station {
	return s.Stations
}
