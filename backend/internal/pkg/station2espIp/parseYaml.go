package station2espIp

import (
	_ "embed"
	"fmt"

	pb "github.com/ueckoken/plarail2022/backend/internal/spec"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"
)

type IStations interface {
	Detail(name string) (*StationDetail, error)
	Enumerate() []Station
}
type Stations struct {
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

func NewStations() (*Stations, error) {
	t := Stations{}
	err := yaml.Unmarshal(ConfigFile, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Stations) Detail(name string) (*StationDetail, error) {
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

func (d *StationDetail) GetAngle(state pb.State) (angle int, err error) {
	if !d.IsAngleDefined() {
		return 0, fmt.Errorf("angles are not defined")
	}
	switch state {
	case pb.State_ON:
		angle = d.OnAngle
	case pb.State_OFF:
		angle = d.OffAngle
	default:
		return 0, status.Errorf(codes.InvalidArgument, "state is not ON or OFF\n")
	}
	return angle, nil
}

func (s *Stations) Enumerate() []Station {
	return s.Stations
}
