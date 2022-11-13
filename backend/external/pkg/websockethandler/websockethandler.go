package websockethandler

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type ClientHandler[T proto.Message] struct {
	upgrader          websocket.Upgrader
	commandFromClient     chan<- T
	channelClient chan ClientChannel[T]
}

func NewClientHandler[T proto.Message](clientHandlerOutput chan <- T, channelClient chan ClientChannel[T]) ClientHandler[T]{
	return ClientHandler[T]{upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			}},
			commandFromClient: clientHandlerOutput,
			channelClient: channelClient,
		}
}

type ClientChannel[T proto.Message] struct {
	SyncToClient chan T
	Done       chan struct{}
}

func (m ClientHandler[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.RemoteAddr)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
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
		log.Println("connection closed")
		cancel()
		return nil
	})
	go handleClientCommand(ctx, c, m.commandFromClient)
	go handleClientPing(ctx, c)
	for cChan := range cChannel.SyncToClient {
		err := c.WriteJSON(cChan)
		if err != nil {
			log.Println("err", err)
			cDone <- struct{}{}
			cancel()
			break
		}
	}
}

func handleClientPing(ctx context.Context, c *websocket.Conn) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := c.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(1*time.Second)); err != nil {
				log.Printf("err occured in clientHandler.handleClientPing, err=%s", err)
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func handleClientCommand[T proto.Message](ctx context.Context, c *websocket.Conn, ch chan <- T) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			r, err := readClientData[T](c)
			if err != nil {
				log.Println(err)
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
