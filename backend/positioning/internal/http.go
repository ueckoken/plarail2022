package internal

import (
	"log"
	"net/http"
	"sync"
	"time"
	"ueckoken/plarail2022-positioning/pkg/gormHandler"
	"ueckoken/plarail2022-positioning/pkg/positionReceiver"
	"ueckoken/plarail2022-positioning/pkg/trainState"
)

type PositionReceiver struct {
	db      gormHandler.SQLHandler
	status  ApplicationStatus
	clients *ClientSet
}

func NewPositionReceiver(db gormHandler.SQLHandler, status ApplicationStatus) PositionReceiver {
	cls := ClientSet{}
	return PositionReceiver{db: db, status: status, clients: &cls}
}

func (pos PositionReceiver) StartPositionReceiver() {
	c := make(chan trainState.State)
	p := positionReceiver.NewPositionReceiverHandler(c)
	n := make(chan ClientNotifier)
	h := ClientHandler{ClientNotification: n}

	go pos.RegisterClient(n)
	go pos.HandleChange(c)
	go pos.UnregisterClient()

	mux := http.NewServeMux()
	mux.Handle("/registerPosition", p)
	mux.Handle("/subscribePosition", h)
	log.Print("listening on :8080")
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

type ClientSet struct {
	mtx     sync.Mutex
	clients []Client
}

type Client struct {
	notifier ClientNotifier
}

func (pos *PositionReceiver) RegisterClient(cn chan ClientNotifier) {
	for n := range cn {
		pos.clients.mtx.Lock()
		pos.clients.clients = append(pos.clients.clients, Client{n})
		pos.clients.mtx.Unlock()
	}
}

func (pos *PositionReceiver) HandleChange(cn chan trainState.State) {
	for c := range cn {
		pos.db.Store(c)
		if !pos.status.HallSensorSpec.CanPredict(c.HallSensorName) {
			continue
		}
		// this should be sorted from old to new
		data := pos.db.FetchFromTrainID(c.TrainID)
		var duration []time.Duration
		for i, d := range data.States {
			if d.HallSensorName == c.HallSensorName {
				n, err := pos.status.HallSensorSpec.Nexts(c.HallSensorName)
				if err != nil {
					log.Println(err)
					continue
				}
				if len(n) != 1 {
					continue
				}
				// can calculate duration
				if n[0].GetName() == data.States[i+1].HallSensorName {
					du := data.States[i+1].FetchedTimeStump.Sub(data.States[i].FetchedTimeStump)
					duration = append(duration, du)
				}
			}
		}
		var sum float64
		var count int
		for _, t := range duration {
			sum += t.Seconds()
		}
		avg := sum / float64(count)
		dat := trainState.PositionAndSpeed{State: c, Speed: avg}
		for _, client := range pos.clients.clients {
			select {
			case client.notifier.Notifier <- dat:
			default:
				log.Println("buffer is full...")
			}
		}
	}
}

func (pos *PositionReceiver) UnregisterClient() {
	for {
		pos.clients.mtx.Lock()
		var deletionList []int
		for i, c := range pos.clients.clients {
			select {
			case <-c.notifier.Unregister:
				deletionList = append(deletionList, i)
			default:
				continue
			}
		}
		pos.clients.deleteClient(deletionList)
		pos.clients.mtx.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func (cl *ClientSet) deleteClient(deletion []int) {
	var tmp []Client
	for i, c := range cl.clients {
		if !contain(deletion, i) {
			tmp = append(tmp, c)
		}
	}
	cl.clients = tmp
}

func contain(list []int, data int) bool {
	for _, l := range list {
		if l == data {
			return true
		}
	}
	return false
}
