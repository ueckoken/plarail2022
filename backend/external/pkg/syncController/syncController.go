package syncController

import (
	"errors"
	"log"
	"sync"
	"time"
	"ueckoken/plarail2022-external/pkg/envStore"
	"ueckoken/plarail2022-external/pkg/servo"
	"ueckoken/plarail2022-external/pkg/stationNameId"
	"ueckoken/plarail2022-external/spec"
)

type StationState struct {
	servo.StationState
}

type stationKVS struct {
	stations  []StationState
	mtx       sync.Mutex
	validator IValidator
}

func newStationKvs() *stationKVS {
	v := NewRouteValidator()
	skvs := stationKVS{validator: v}
	return &skvs
}
func (skvs *stationKVS) update(u StationState) error {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	// err := skvs.validator.Validate(u, skvs.stations)
	var err error
	if err != nil {
		return err
	}
	log.Printf("validation passed u=`%v`\n", u)
	for i, s := range skvs.stations {
		if s.StationID == u.StationID {
			skvs.stations[i].State = u.State
			return nil
		}
	}
	skvs.stations = append(skvs.stations, u)
	return nil
}

// forceUpdate differs from update for ignore route validation.
func (skvs *stationKVS) forceUpdate(u StationState) {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	for i, s := range skvs.stations {
		if s.StationID == u.StationID {
			skvs.stations[i].State = u.State
			return
		}
	}
	skvs.stations = append(skvs.stations, u)
}
func (skvs *stationKVS) get(stationID int32) (station StationState, err error) {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	for _, s := range skvs.stations {
		if s.StationID == stationID {
			return s, nil
		}
	}
	return StationState{}, errors.New("Not found")
}

func (skvs *stationKVS) retrieve() []StationState {
	return skvs.stations
}

type SyncController struct {
	ClientHandler2syncController chan StationState
	SyncController2clientHandler chan StationState
	InitServoRoute               chan StationState
	Environment                  *envStore.Env
}

func (s *SyncController) StartSyncController() {
	kvs := newStationKvs()

	go s.Init(NewInitializeRule())
	s.initNode(s.Environment, kvs)

	go s.periodicallySync(kvs)
	s.triggeredSync(s.Environment, kvs)
}

func (s *SyncController) initNode(e *envStore.Env, kvs *stationKVS) {
	for c := range s.InitServoRoute {
		kvs.forceUpdate(c)
		c2i := servo.NewCommand2Internal(c.StationState, e)
		log.Println("initNode: ", c2i.String())
		err := c2i.Send()
		if err != nil {
			log.Fatalf("initNode Send err: `%v`\n", err)
			return
		}
	}
}

func (s *SyncController) triggeredSync(e *envStore.Env, kvs *stationKVS) {
	for c := range s.ClientHandler2syncController {
		err := kvs.update(c)
		if err != nil {
			log.Println("syncController validator err: ", err)
			continue
		}
		c2i := servo.NewCommand2Internal(c.StationState, e)
		err = c2i.Send()
		if err != nil {
			log.Println("syncController send err: ", err)
			continue
		}
		s.SyncController2clientHandler <- c
	}
}

func (s *SyncController) periodicallySync(kvs *stationKVS) {
	ch := time.Tick(2 * time.Second)
	for range ch {
		kvs.mtx.Lock()
		k := kvs.retrieve()
		for _, st := range k {
			select {
			case s.SyncController2clientHandler <- st:
			default:
				log.Println("buffer full for:")
			}
		}
		kvs.mtx.Unlock()
	}
}

func (s SyncController) Init(r *InitRule) {
	for _, sta := range r.Stations {
		id, err := stationNameId.Name2Id(sta.Name)
		if err != nil {
			log.Fatalln(err)
		}
		state, ok := spec.RequestSync_State_value[sta.State]
		if !ok {
			log.Fatalln(sta.State, "is incorrect")
		}
		s.InitServoRoute <- StationState{
			struct {
				StationID int32
				State     int32
			}{
				StationID: id,
				State:     state,
			},
		}
		time.Sleep(500 * time.Millisecond)
	}
	close(s.InitServoRoute)
}
