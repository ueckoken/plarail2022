package serveGrpc

import (
	"fmt"
	"github.com/ueckoken/plarail2022/backend/internal/pkg/station2espIp"
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
