package internal

import (
	"time"
	"ueckoken/plarail2022-external/pkg/synccontroller"
	"ueckoken/plarail2022-external/spec"

	"go.uber.org/zap"
)

// StartStationSync starts sync controller for station state.
func StartStationSync(logger *zap.Logger, syncInput chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State], syncOutput chan<- synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]) {
	s := synccontroller.NewSyncController(logger, syncInput, syncOutput)

	go s.Run()
	initStationSync(logger.Named("initialization"), NewRule(), syncInput)
}

// InitStationSync initalize state.
func initStationSync(logger *zap.Logger, r *InitRule, initializer chan<- synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]) {
	for _, sta := range r.Stations {
		id, ok := spec.Stations_StationId_value[sta.Name]
		if !ok {
			logger.Error("unknown station name", zap.String("station", sta.Name))
			continue
		}
		state, ok := spec.RequestSync_State_value[sta.State]
		if !ok {
			logger.Error("unknown state", zap.String("state", sta.State))
			continue
		}
		initializer <- synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]{Key: spec.Stations_StationId(id), Value: spec.Command2InternalRequest_State(state)}
		time.Sleep(500 * time.Millisecond)
	}
}
