package hallsensor

import (
	_ "embed"
	"fmt"

	"github.com/ghodss/yaml"
)

//go:embed embed/hallsensor.yaml
var HallSensorSetting []byte

type PhysicalSensors struct {
	sensors map[string]PhysicalSensor
}

func (p *PhysicalSensors) Nexts(name string) (ret []PhysicalSensor, err error) {
	s, ok := p.sensors[name]
	if !ok {
		return nil, fmt.Errorf("not found such name: %s", name)
	}
	nexts := s.nexts
	for _, nx := range nexts {
		n, ok := p.sensors[nx]
		if !ok {
			return nil, fmt.Errorf("not found such name when finding nexts: %s", name)
		}
		ret = append(ret, n)
	}
	if len(ret) == 0 {
		return nil, fmt.Errorf("no nexts for: %s", name)
	}
	return ret, nil
}

func (p *PhysicalSensors) CanPredict(name string) bool {
	s, ok := p.sensors[name]
	if !ok {
		return false
	}
	return s.predict
}

func NewPhysicalSensors() (PhysicalSensors, error) {
	var y SensorYaml
	if err := yaml.Unmarshal(HallSensorSetting, &y); err != nil {
		return PhysicalSensors{}, err
	}
	p := make(map[string]PhysicalSensor)
	ps := PhysicalSensors{sensors: p}
	for _, h := range y.Halls {
		ps.sensors[h.Name] = PhysicalSensor{name: h.Name, predict: h.Predict, nexts: h.Nexts}
	}
	return ps, nil
}

type PhysicalSensor struct {
	name    string
	predict bool
	nexts   []string
}

func (p *PhysicalSensor) GetName() string {
	return p.name
}

type SensorYaml struct {
	Halls []struct {
		Name    string   `yaml:"name"`
		Predict bool     `yaml:"predict"`
		Nexts   []string `yaml:"nexts"`
	} `yaml:"halls"`
}
