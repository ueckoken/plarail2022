package internal

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"ueckoken/plarail2022-external/pkg/clientHandler"
	"ueckoken/plarail2022-external/pkg/envStore"
	"ueckoken/plarail2022-external/pkg/syncController"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HTTPServer struct {
	ClientHandler2syncController chan syncController.StationState
	SyncController2clientHandler chan syncController.StationState
	Environment                  *envStore.Env
	NumberOfClientConnection     *prometheus.GaugeVec
	TotalClientConnection        *prometheus.CounterVec
	TotalCLientCommands          *prometheus.CounterVec
	Clients                      *ClientsCollection
}

type ClientsCollection struct {
	Clients []clientHandler.ClientChannel
	mtx     sync.Mutex
}

func (h *HTTPServer) StartServer() {
	clientChannelSend := make(chan clientHandler.ClientChannel)
	go h.registerClient(clientChannelSend)
	go h.handleChanges()
	go h.unregisterClient()
	r := mux.NewRouter()
	prometheus.MustRegister(h.NumberOfClientConnection)
	prometheus.MustRegister(h.TotalClientConnection)
	prometheus.MustRegister(h.TotalCLientCommands)
	r.HandleFunc("/", clientHandler.HandleStatic)
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	r.Handle("/ws", clientHandler.ClientHandler{Upgrader: upgrader, ClientCommand: h.ClientHandler2syncController, ClientChannelSend: clientChannelSend})
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

func (h *HTTPServer) handleChanges() {
	for d := range h.SyncController2clientHandler {
		h.Clients.mtx.Lock()
		h.TotalCLientCommands.With(prometheus.Labels{}).Inc()
		for _, c := range h.Clients.Clients {
			select {
			case c.ClientSync <- d:
			default:
				log.Println("buffer is full when send...")
				continue
			}
		}
		h.Clients.mtx.Unlock()
	}
	time.Sleep(1 * time.Second)
}

func (h *HTTPServer) registerClient(cn chan clientHandler.ClientChannel) {
	for n := range cn {
		func(h *HTTPServer, n clientHandler.ClientChannel) {
			h.Clients.mtx.Lock()
			defer h.Clients.mtx.Unlock()
			h.TotalClientConnection.With(prometheus.Labels{}).Inc()
			h.Clients.Clients = append(h.Clients.Clients, n)
		}(h, n)
	}
}

func (h *HTTPServer) unregisterClient() {
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

func (cl *ClientsCollection) deleteClient(deletion []int) {
	var tmp []clientHandler.ClientChannel
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
