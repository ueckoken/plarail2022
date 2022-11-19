package websockethandler

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type ClientHandler[T proto.Message] struct {
	logger            *zap.Logger
	upgrader          websocket.Upgrader
	commandFromClient chan<- T
	channelClient     chan ClientChannel[T]
}

// NewClientHandler creates client handler.
func NewClientHandler[T proto.Message](logger *zap.Logger, clientHandlerOutput chan<- T, channelClient chan ClientChannel[T]) ClientHandler[T] {
	return ClientHandler[T]{
		logger: logger,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			}},
		commandFromClient: clientHandlerOutput,
		channelClient:     channelClient,
	}
}

type ClientChannel[T proto.Message] struct {
	SyncToClient chan T
	Done         chan struct{}
}

func (m ClientHandler[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.RemoteAddr)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		m.logger.Error("failed to upgrade", zap.Error(err))
		return
	}
	defer c.Close()
	ctx, cancel := context.WithCancel(context.Background())
	var cSync = make(chan T)
	var cDone = make(chan struct{})
	var cChannel = ClientChannel[T]{cSync, cDone}
	m.channelClient <- cChannel
	c.SetPongHandler(func(string) error {
		return c.SetReadDeadline(time.Now().Add(20 * time.Second))
	})
	c.SetCloseHandler(func(code int, text string) error {
		m.logger.Info("connection closed", zap.Int("code", code), zap.String("text", text))
		cancel()
		return nil
	})
	go handleClientCommand(ctx, m.logger.Named("client-handler"), c, m.commandFromClient)
	go handleClientPing(ctx, m.logger.Named("ping-handler"), c)
	for cChan := range cChannel.SyncToClient {
		err := c.WriteJSON(cChan)
		if err != nil {
			m.logger.Info("failed to send data to client, closing connection...", zap.Error(err))
			cDone <- struct{}{}
			cancel()
			break
		}
	}
}

func handleClientPing(ctx context.Context, logger *zap.Logger, c *websocket.Conn) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := c.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(1*time.Second)); err != nil {
				logger.Error("failed to send ping", zap.Error(err))
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func handleClientCommand[T proto.Message](ctx context.Context, logger *zap.Logger, c *websocket.Conn, ch chan<- T) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			r, err := readClientData[T](c)
			if err != nil {
				logger.Error("failed to read client packet", zap.Error(err))
				return
			}
			ch <- *r
		}
	}
}

func readClientData[T proto.Message](c *websocket.Conn) (*T, error) {
	_, msg, err := c.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("websocket read failed: %e", err)
	}
	var ud T
	err = json.Unmarshal(msg, &ud)
	if err != nil {
		return nil, fmt.Errorf("bad json format: %e", err)
	}

	return &ud, nil
}

//go:embed embed/index.html
var IndexHTML []byte

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(IndexHTML)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
