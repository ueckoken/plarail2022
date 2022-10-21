package hallsensor

import (
	_ "embed"
	"fmt"

	"github.com/ghodss/yaml"
)

//go:embed embed/esp32.yaml
var esp32AndPins []byte

type Esp32Pins struct {
	Esp32 []struct {
		Address string `yaml:"address"`
		Sensors []struct {
			Name string `yaml:"name"`
			Pin  int    `yaml:"pin"`
		} `yaml:"sensors"`
	} `yaml:"esp32"`
}

func NewEsp32PinSetting() (Esp32Pins, error) {
	var y Esp32Pins
	if err := yaml.Unmarshal(esp32AndPins, &y); err != nil {
		return Esp32Pins{}, err
	}
	return y, nil
}
func (esp *Esp32Pins) Search(addr string, pin int) (name string, err error) {
	for _, e := range esp.Esp32 {
		if addr != e.Address {
			continue
		}
		for _, s := range e.Sensors {
			if pin == s.Pin {
				return s.Name, nil
			}
		}
	}
	return "", fmt.Errorf("not found such addr: %s, pin:%d", addr, pin)
}
