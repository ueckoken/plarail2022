package envStore

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/vrischmann/envconfig"
)

type hostnamePort string
type Port int32

func (t *hostnamePort) Unmarshal(s string) error {
	err := validator.New().Var(s, "hostname_port")
	if err != nil {
		return err
	}
	*t = hostnamePort(s)
	return nil
}
func (t *hostnamePort) String() string {
	return string(*t)
}
func (p *Port) Unmarshal(s string) error {
	d, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	if !(0 < d && d <= 65535) {
		return fmt.Errorf("Port range failed; Port: %d ", d)
	}
	*p = Port(d)
	return nil
}

type Env struct {
	ClientSideServer struct {
		Port Port `envconfig:"default=54321"`
	}
	InternalServer struct {
		Addr        hostnamePort
		TimeoutSec  time.Duration `envconfig:"default=1s"`
		MetricsPort Port          `envconfig:"default=9100"`
		// SslCertPath string
	}
}

func GetEnv() *Env {
	var env Env
	if err := envconfig.Init(&env); err != nil {
		panic(err)
	}
	return &env
}
