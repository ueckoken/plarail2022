package synccontroller

import (
	"fmt"
	"log"
	"sync"
	"time"

	"go.uber.org/zap"
)

type KV[T, U comparable] struct {
	Key   T
	Value U
}

type stationKVS[T, U comparable] struct {
	values map[T]*U
	mtx    sync.Mutex
}

func newStationKVS[T, U comparable]() *stationKVS[T, U] {
	skvs := stationKVS[T, U]{values: make(map[T]*U, 0)}
	return &skvs
}

func (skvs *stationKVS[T, U]) update(key T, value U) error {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	// err := skvs.validator.Validate(u, skvs.stations)
	var err error
	if err != nil {
		return err
	}
	skvs.values[key] = &value
	return nil
}

func (skvs *stationKVS[T, U]) get(key T) (value *U, err error) {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	dat, ok := skvs.values[key]
	if !ok {
		return nil, fmt.Errorf("not found such key=%+v", key)
	}
	return dat, nil

}

func (skvs *stationKVS[T, U]) retrieve() map[T]*U {
	return skvs.values
}

type SyncController[T, U comparable] struct {
	logger      *zap.Logger
	stateInput  <-chan KV[T, U]
	stateOutput chan<- KV[T, U]
	kvs         *stationKVS[T, U]
}

// NewSyncController creates new sync controller.
func NewSyncController[T, U comparable](logger *zap.Logger, stateInput <-chan KV[T, U], stateOutput chan<- KV[T, U]) *SyncController[T, U] {
	return &SyncController[T, U]{logger: logger, stateInput: stateInput, stateOutput: stateOutput, kvs: newStationKVS[T, U]()}
}

// Run runs sync controller, and this will block forever.
// This receives and saves state from stateInput.
// This emits state when received, or periodically.
func (s *SyncController[T, U]) Run() {
	go s.periodicallySync()
	s.triggeredSync()
}

func (s *SyncController[T, U]) triggeredSync() {
	for c := range s.stateInput {
		err := s.kvs.update(c.Key, c.Value)
		if err != nil {
			log.Println("syncController validator err: ", err)
			continue
		}
		s.stateOutput <- c
	}
}

func (s *SyncController[T, U]) periodicallySync() {
	ch := time.Tick(2 * time.Second)
	s.kvs.mtx.Lock()
	defer s.kvs.mtx.Unlock()
	for range ch {
		d := s.kvs.retrieve()
		for key, value := range d {
			select {
			case s.stateOutput <- KV[T, U]{Key: key, Value: *value}:
			default:
				log.Println("buffer full for:")
			}
		}
	}
}
