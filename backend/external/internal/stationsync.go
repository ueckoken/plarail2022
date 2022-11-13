package internal

import (
	"log"
	"time"
	"ueckoken/plarail2022-external/pkg/synccontroller"
	"ueckoken/plarail2022-external/spec"
)

// StartStationSync starts sync controller for station state.
func StartStationSync(syncInput chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State], syncOutput chan<- synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]) {
	s := synccontroller.NewSyncController[spec.Stations_StationId, spec.Command2InternalRequest_State](syncInput, syncOutput)

	go s.Run()
	initStationSync(NewRule(), syncInput)
}

// Init initalize state
func initStationSync(r *InitRule, initializer chan<- synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]) {
	for _, sta := range r.Stations {
		id, ok := spec.Stations_StationId_value[sta.Name]
		if !ok {
			log.Println("unknown station name", sta.Name)
			continue
		}
		state, ok := spec.RequestSync_State_value[sta.State]
		if !ok {
			log.Println("unknown state", sta.State)
		}
		initializer <- synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]{Key: spec.Stations_StationId(id), Value: spec.Command2InternalRequest_State(state)}
		time.Sleep(500 * time.Millisecond)
	}
}
