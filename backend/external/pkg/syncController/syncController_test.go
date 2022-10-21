package syncController

import (
	"sync"
	"testing"
	"ueckoken/plarail2022-external/pkg/servo"
)

func (skvs *stationKVS) contain(ss StationState) bool {
	for _, s := range skvs.stations {
		if s == ss {
			return true
		}
	}
	return false
}

type ValidatorMock struct{}

func (vm *ValidatorMock) Validate(_ StationState, _ []StationState) error {
	return nil
}

func TestSyncController_update(t *testing.T) {
	station1 := StationState{servo.StationState{StationID: 1, State: 1}}
	station2 := StationState{servo.StationState{StationID: 2, State: 1}}
	kvs := stationKVS{
		stations:  nil,
		mtx:       sync.Mutex{},
		validator: &ValidatorMock{},
	}
	err := kvs.update(station1)
	if err != nil {
		t.Errorf("validate failed. `%v`", err)
	}
	if !kvs.contain(station1) {
		t.Errorf("append failed")
	}

	// new station append
	err = kvs.update(station2)
	if err != nil {
		t.Errorf("validate failed. `%v`", err)
	}
	if !kvs.contain(station2) {
		t.Errorf("station add failed")
	}
	if len(kvs.stations) != 2 {
		t.Errorf("append failed")
	}
	if !kvs.contain(station1) {
		t.Errorf("stations before update are not keeping")
	}
	if !kvs.contain(station1) && !kvs.contain(station2) {
		t.Errorf("station2 is not append with `update` method")
	}

	// update exist station data
	station1 = StationState{servo.StationState{StationID: 1, State: 0}}
	err = kvs.update(station1)
	if err != nil {
		t.Errorf("validate failed. `%v`", err)
	}
	if len(kvs.stations) != 2 {
		t.Errorf("append failed")
	}
	if !kvs.contain(station1) {
		t.Errorf("not update station data")
	}
}
func TestSyncController_get(t *testing.T) {
	station1 := StationState{servo.StationState{StationID: 1, State: 1}}
	station2 := StationState{servo.StationState{StationID: 2, State: 1}}
	skvs := stationKVS{
		stations: nil,
		mtx:      sync.Mutex{},
	}
	// member is not exist
	station, err := skvs.get(0)
	if station != (StationState{}) {
		t.Errorf("'station' is expect for empty but not empty")
	}

	if err == nil {
		t.Errorf("'err' is expect not nil")
	} else if err.Error() != "Not found" {
		t.Errorf("err.Error() expect 'Not found' but return %e", err)
	}

	skvs = stationKVS{
		stations: []StationState{station1, station2},
		mtx:      sync.Mutex{},
	}
	station, err = skvs.get(1)
	if station != station1 {
		t.Errorf("'station1' is expect but called station%d", station.StationID)
	}
	if err != nil {
		t.Errorf("return err is not nil: %e", err)
	}

	station, err = skvs.get(2)
	if station != station2 {
		t.Errorf("'station2' is expect but called station%d", station.StationID)
	}
	if err != nil {
		t.Errorf("return err is not nil: %e", err)
	}

	// test for call 'get' not exist record
	station, err = skvs.get(3)
	if station != (StationState{}) {
		t.Errorf("'station' is expect for empty but called station%d", station.StationID)
	}
	if err == nil {
		t.Errorf("expect err but return nil")
	} else if err.Error() != "Not found" {
		t.Errorf("err.Error() expect 'Not found' but return %e", err)
	}
}
