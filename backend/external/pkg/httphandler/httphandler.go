package httphandler

import (
	"fmt"
	"github.com/ueckoken/plarail2022/backend/external/pkg/envStore"
	"github.com/ueckoken/plarail2022/backend/external/pkg/websockethandler"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// HTTPServer is a server managing websockethandler and websocket clients.
type HTTPServer[T proto.Message] struct {
	logger                   *zap.Logger
	httpOutput               chan<- T
	httpInput                <-chan T
	environment              *envStore.Env
	numberOfClientConnection *prometheus.GaugeVec
	totalClientConnection    *prometheus.CounterVec
	totalCLientCommands      *prometheus.CounterVec
	clients                  *ClientsCollection[T]
}

// NewHTTPServer creates HTTP server.
func NewHTTPServer[T proto.Message](logger *zap.Logger, httpOutput chan<- T, httpInput <-chan T, env *envStore.Env) *HTTPServer[T] {
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
		logger:                   logger,
		httpOutput:               httpOutput,
		httpInput:                httpInput,
		environment:              env,
		numberOfClientConnection: clientConn,
		totalClientConnection:    clientConnTotal,
		totalCLientCommands:      controlCommandTotal,
		clients:                  &ClientsCollection[T]{},
	}
}

// ClientsCollection is a set of clients.
// When client is disconnected, this will remove it and stop sending data.
type ClientsCollection[T proto.Message] struct {
	clients []websockethandler.ClientChannel[T]
	mtx     sync.Mutex
}

// StartServer starts HTTP server and websocket server.
func (h *HTTPServer[T]) StartServer(port int) {
	clientChannelSend := make(chan websockethandler.ClientChannel[T])
	handlerInput := make(chan T)
	go h.registerClient(clientChannelSend)
	go h.broadcastChanges()
	go h.unregisterClients()
	r := mux.NewRouter()
	prometheus.MustRegister(h.numberOfClientConnection)
	prometheus.MustRegister(h.totalClientConnection)
	prometheus.MustRegister(h.totalCLientCommands)
	r.HandleFunc("/", websockethandler.HandleStatic)
	r.Handle("/ws", websockethandler.NewClientHandler(h.logger.Named("ws-handler"), handlerInput, clientChannelSend))
	r.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{
		Handler:           r,
		Addr:              fmt.Sprintf("0.0.0.0:%d", port),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	h.logger.Info("start listening")
	if err := srv.ListenAndServe(); err != nil {
		h.logger.Panic("failed to serve", zap.Error(err))
	}
}

// broadcastChanges receives the change from main channel and broadcast it to clients.
func (h *HTTPServer[T]) broadcastChanges() {
	for d := range h.httpInput {
		h.clients.mtx.Lock()
		h.totalCLientCommands.With(prometheus.Labels{}).Inc()
		for _, c := range h.clients.clients {
			select {
			case c.SyncToClient <- d:
			default:
				h.logger.Info("client buffer is full...")
				continue
			}
		}
		h.clients.mtx.Unlock()
	}
}

// registerClient registers a client that listening changes.
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

// unregisterClients remove clients that seems to be disconnected.
func (h *HTTPServer[T]) unregisterClients() {
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
	}
}

// deleteClient deletes clients specified with arg.
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
