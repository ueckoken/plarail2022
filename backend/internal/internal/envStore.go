package internal

import (
	"fmt"
	"strconv"
	"time"

	"github.com/vrischmann/envconfig"
)

type Port int32

func (p *Port) Unmarshal(s string) error {
	d, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	if !(0 < d && d <= 65535) {
		return fmt.Errorf("Port range failed; Port: %d", d)
	}
	*p = Port(d)
	return nil
}

func (p *Port) String() string {
	return fmt.Sprintf("%d", *p)
}

type Env struct {
	ExternalSideServer struct {
		Port Port `envconfig:"default=54321"`
		// SslCertPath string
		MetricsPort Port `envconfig:"default=9100"`
	}
	NodeConnection struct {
		Timeout time.Duration `envconfig:"default=1s"`
	}
}

func GetEnv() *Env {
	var env Env
	if err := envconfig.Init(&env); err != nil {
		panic(err)
	}
	return &env
}
