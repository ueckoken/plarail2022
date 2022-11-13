package internal

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"ueckoken/plarail2022-external/pkg/envStore"
	"ueckoken/plarail2022-external/pkg/websockethandler"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/protobuf/proto"
)

type HTTPServer[T proto.Message] struct {
	StateOutput              chan<- T
	StateInput               <-chan T
	Environment              *envStore.Env
	NumberOfClientConnection *prometheus.GaugeVec
	TotalClientConnection    *prometheus.CounterVec
	TotalCLientCommands      *prometheus.CounterVec
	Clients                  *ClientsCollection[T]
}

type ClientsCollection[T proto.Message] struct {
	Clients []websockethandler.ClientChannel[T]
	mtx     sync.Mutex
}

func (h *HTTPServer[T]) StartServer() {
	clientChannelSend := make(chan websockethandler.ClientChannel[T])
	handlerInput := make(chan T)
	go h.registerClient(clientChannelSend)
	go h.handleChanges()
	go h.unregisterClient()
	r := mux.NewRouter()
	prometheus.MustRegister(h.NumberOfClientConnection)
	prometheus.MustRegister(h.TotalClientConnection)
	prometheus.MustRegister(h.TotalCLientCommands)
	r.HandleFunc("/", websockethandler.HandleStatic)
	r.Handle("/ws", websockethandler.NewClientHandler(handlerInput, clientChannelSend))
	r.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{
		Handler:           r,
		Addr:              fmt.Sprintf("0.0.0.0:%d", h.Environment.ClientSideServer.Port),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func (h *HTTPServer[T]) handleChanges() {
	for d := range h.StateInput {
		h.Clients.mtx.Lock()
		h.TotalCLientCommands.With(prometheus.Labels{}).Inc()
		for _, c := range h.Clients.Clients {
			select {
			case c.SyncToClient <- d:
			default:
				log.Println("buffer is full when send...")
				continue
			}
		}
		h.Clients.mtx.Unlock()
	}
	time.Sleep(1 * time.Second)
}

func (h *HTTPServer[T]) registerClient(cn <-chan websockethandler.ClientChannel[T]) {
	for n := range cn {
		func(h *HTTPServer[T], n websockethandler.ClientChannel[T]) {
			h.Clients.mtx.Lock()
			defer h.Clients.mtx.Unlock()
			h.TotalClientConnection.With(prometheus.Labels{}).Inc()
			h.Clients.Clients = append(h.Clients.Clients, n)
		}(h, n)
	}
}

func (h *HTTPServer[T]) unregisterClient() {
	for {
		h.Clients.mtx.Lock()
		var deletionList []int
		for i, c := range h.Clients.Clients {
			select {
			case <-c.Done:
				deletionList = append(deletionList, i)
			default:
				continue
			}
		}
		h.Clients.deleteClient(deletionList)
		h.Clients.mtx.Unlock()
		h.NumberOfClientConnection.With(prometheus.Labels{}).Set(float64(len(h.Clients.Clients)))
		time.Sleep(1 * time.Second)
	}
}

func (cl *ClientsCollection[T]) deleteClient(deletion []int) {
	var tmp []websockethandler.ClientChannel[T]
	for i, c := range cl.Clients {
		if !contain(deletion, i) {
			tmp = append(tmp, c)
		}
	}
	cl.Clients = tmp
}

func contain(list []int, data int) bool {
	for _, l := range list {
		if l == data {
			return true
		}
	}
	return false
}
