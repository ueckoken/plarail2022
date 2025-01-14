package synccontroller

import (
	"github.com/ueckoken/plarail2022/backend/external/spec"
	"strings"
	"testing"
)

func (skvs *stationKVS[T, U]) contain(dat KV[T, U]) bool {
	d, ok := skvs.values[dat.Key]
	if !ok {
		return false
	}
	if *d != dat.Value {
		return false
	}
	return true
}

type ValidatorMock struct{}

func TestSyncController_update(t *testing.T) {
	station1 := KV[spec.StationId, spec.PointStateEnum]{Key: spec.StationId(1), Value: spec.PointStateEnum(1)}
	station2 := KV[spec.StationId, spec.PointStateEnum]{Key: spec.StationId(2), Value: spec.PointStateEnum(1)}
	kvs := newStationKVS[spec.StationId, spec.PointStateEnum]()
	changed := kvs.update(station1)
	if !changed {
		t.Errorf("kvs is not changed")
	}
	if !kvs.contain(station1) {
		t.Errorf("append failed")
	}

	// new station append
	changed = kvs.update(station2)
	if !changed {
		t.Errorf("kvs is not changed")
	}
	if !kvs.contain(station2) {
		t.Errorf("station add failed")
	}
	if len(kvs.values) != 2 {
		t.Errorf("append failed")
	}
	if !kvs.contain(station1) {
		t.Errorf("stations before update are not keeping")
	}
	if !kvs.contain(station1) && !kvs.contain(station2) {
		t.Errorf("station2 is not append with `update` method")
	}

	// update exist station data
	station1 = KV[spec.StationId, spec.PointStateEnum]{Key: spec.StationId(1), Value: spec.PointStateEnum(0)}
	changed = kvs.update(station1)
	if !changed {
		t.Errorf("kvs is not changed")
	}
	if len(kvs.values) != 2 {
		t.Errorf("append failed")
	}
	if !kvs.contain(station1) {
		t.Errorf("not update station data")
	}
}
func TestSyncController_get(t *testing.T) {
	station1 := KV[spec.StationId, spec.PointStateEnum]{Key: spec.StationId(1), Value: spec.PointStateEnum(1)}
	station2 := KV[spec.StationId, spec.PointStateEnum]{Key: spec.StationId(2), Value: spec.PointStateEnum(1)}
	kvs := newStationKVS[spec.StationId, spec.PointStateEnum]()
	// member is not exist
	_, err := kvs.get(0)
	if err == nil {
		t.Errorf("'err' is expect not nil")
	} else if !strings.Contains(err.Error(), "not found") {
		t.Errorf("err.Error() expect 'Not found' but return %e", err)
	}

	kvs = newStationKVS[spec.StationId, spec.PointStateEnum]()
	changed := kvs.update(station1)
	if !changed {
		t.Errorf("kvs is not changed")
	}
	station, err := kvs.get(spec.StationId(1))
	if err != nil {
		t.Errorf("return err is not nil: %e", err)
	}
	if *station != station1.Value {
		t.Errorf("'station1' is expect but called station%d", station)
	}

	changed = kvs.update(station2)
	if !changed {
		t.Errorf("kvs is not changed")
	}
	station, err = kvs.get(2)
	if err != nil {
		t.Errorf("return err is not nil: %e", err)
	}
	if *station != station2.Value {
		t.Errorf("'station2' is expect but called station%d", station)
	}

	// test for call 'get' not exist record
	_, err = kvs.get(3)
	if err == nil {
		t.Errorf("expect err but return nil")
	}
	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("err.Error() expect 'Not found' but return %e", err)
	}
}
