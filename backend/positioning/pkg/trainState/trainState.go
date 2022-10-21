package trainState

import (
	"time"
)

type States struct {
	States []State
}

type State struct {
	TrainID          int
	HallSensorName   string
	FetchedTimeStump time.Time
}

type PositionAndSpeed struct {
	State
	Speed float64
}
