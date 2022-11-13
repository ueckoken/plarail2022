package httphandler

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
	httpOutput               chan<- T
	httpInput                <-chan T
	environment              *envStore.Env
	numberOfClientConnection *prometheus.GaugeVec
	totalClientConnection    *prometheus.CounterVec
	totalCLientCommands      *prometheus.CounterVec
	clients                  *ClientsCollection[T]
}

func NewHTTPServer[T proto.Message](httpOutput chan<- T, httpInput <-chan T, env *envStore.Env) *HTTPServer[T] {
	const namespace = "plarailexternal"
	clientConn := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "clients_connections_seconds",
			Help:      "Number of connections handling websocket",
		},
		[]string{},
	)

	clientConnTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "clients_connections_total",
			Help:      "Total client connection",
		},
		[]string{},
	)

	controlCommandTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "client_commands_total",
			Help:      "Total client commands",
		},
		[]string{},
	)
	return &HTTPServer[T]{
		httpOutput:               httpOutput,
		httpInput:                httpInput,
		environment:              env,
		numberOfClientConnection: clientConn,
		totalClientConnection:    clientConnTotal,
		totalCLientCommands:      controlCommandTotal,
		clients:                  &ClientsCollection[T]{},
	}
}

type ClientsCollection[T proto.Message] struct {
	clients []websockethandler.ClientChannel[T]
	mtx     sync.Mutex
}

func (h *HTTPServer[T]) StartServer() {
	clientChannelSend := make(chan websockethandler.ClientChannel[T])
	handlerInput := make(chan T)
	go h.registerClient(clientChannelSend)
	go h.handleChanges()
	go h.unregisterClient()
	r := mux.NewRouter()
	prometheus.MustRegister(h.numberOfClientConnection)
	prometheus.MustRegister(h.totalClientConnection)
	prometheus.MustRegister(h.totalCLientCommands)
	r.HandleFunc("/", websockethandler.HandleStatic)
	r.Handle("/ws", websockethandler.NewClientHandler(handlerInput, clientChannelSend))
	r.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{
		Handler:           r,
		Addr:              fmt.Sprintf("0.0.0.0:%d", h.environment.ClientSideServer.Port),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func (h *HTTPServer[T]) handleChanges() {
	for d := range h.httpInput {
		h.clients.mtx.Lock()
		h.totalCLientCommands.With(prometheus.Labels{}).Inc()
		for _, c := range h.clients.clients {
			select {
			case c.SyncToClient <- d:
			default:
				log.Println("buffer is full when send...")
				continue
			}
		}
		h.clients.mtx.Unlock()
	}
	time.Sleep(1 * time.Second)
}

func (h *HTTPServer[T]) registerClient(cn <-chan websockethandler.ClientChannel[T]) {
	for n := range cn {
		func(h *HTTPServer[T], n websockethandler.ClientChannel[T]) {
			h.clients.mtx.Lock()
			defer h.clients.mtx.Unlock()
			h.totalClientConnection.With(prometheus.Labels{}).Inc()
			h.clients.clients = append(h.clients.clients, n)
		}(h, n)
	}
}

func (h *HTTPServer[T]) unregisterClient() {
	for {
		h.clients.mtx.Lock()
		var deletionList []int
		for i, c := range h.clients.clients {
			select {
			case <-c.Done:
				deletionList = append(deletionList, i)
			default:
				continue
			}
		}
		h.clients.deleteClient(deletionList)
		h.clients.mtx.Unlock()
		h.numberOfClientConnection.With(prometheus.Labels{}).Set(float64(len(h.clients.clients)))
		time.Sleep(1 * time.Second)
	}
}

func (cl *ClientsCollection[T]) deleteClient(deletion []int) {
	var tmp []websockethandler.ClientChannel[T]
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
