package internal

import (
	"github.com/ueckoken/plarail2022/backend/external/pkg/synccontroller"
	"github.com/ueckoken/plarail2022/backend/external/spec"
	"time"

	"go.uber.org/zap"
)

// StartStationSync starts sync controller for station state.
func StartStationSync(logger *zap.Logger, syncInput chan synccontroller.KV[spec.StationId, spec.PointStateEnum], syncOutput chan<- synccontroller.KV[spec.StationId, spec.PointStateEnum]) {
	s := synccontroller.NewSyncController(logger, syncInput, syncOutput)

	go s.Run()
	initStationSync(logger.Named("initialization"), NewRule(), syncInput)
}

// InitStationSync initalize state.
func initStationSync(logger *zap.Logger, r *InitRule, initializer chan<- synccontroller.KV[spec.StationId, spec.PointStateEnum]) {
	for _, sta := range r.Stations {
		id, ok := spec.StationId_value[sta.Name]
		if !ok {
			logger.Error("unknown station name", zap.String("station", sta.Name))
			continue
		}
		state, ok := spec.PointStateEnum_value[sta.State]
		if !ok {
			logger.Error("unknown state", zap.String("state", sta.State))
			continue
		}
		select {
		case initializer <- synccontroller.KV[spec.StationId, spec.PointStateEnum]{Key: spec.StationId(id), Value: spec.PointStateEnum(state)}:
		default:
			logger.Error("failed to initialize sync controller, buffer full")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
