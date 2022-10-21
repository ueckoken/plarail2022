package clientHandler

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"ueckoken/plarail2022-external/pkg/servo"
	"ueckoken/plarail2022-external/pkg/syncController"
	pb "ueckoken/plarail2022-external/spec"

	"github.com/gorilla/websocket"
)

type ClientHandler struct {
	Upgrader          websocket.Upgrader
	ClientCommand     chan syncController.StationState
	ClientChannelSend chan ClientChannel
}

type ClientChannel struct {
	ClientSync chan syncController.StationState
	Done       chan struct{}
}
type clientSendData struct {
	StationName string `json:"station_name"`
	State       string `json:"state"`
}

func (m ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.RemoteAddr)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c, err := m.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()
	ctx, cancel := context.WithCancel(context.Background())
	var cSync = make(chan syncController.StationState, 64)
	var cDone = make(chan struct{})
	var cChannel = ClientChannel{cSync, cDone}
	m.ClientChannelSend <- cChannel
	c.SetPongHandler(func(string) error {
		return c.SetReadDeadline(time.Now().Add(20 * time.Second))
	})
	c.SetCloseHandler(func(code int, text string) error {
		log.Println("connection closed")
		cancel()
		return nil
	})
	go handleClientCommand(ctx, c, &m)
	go handleClientPing(ctx, c)
	for cChan := range cChannel.ClientSync {
		dat := clientSendData{StationName: pb.Stations_StationId_name[cChan.StationID], State: pb.RequestSync_State_name[cChan.State]}
		err := c.WriteJSON(dat)
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

func handleClientCommand(ctx context.Context, c *websocket.Conn, m *ClientHandler) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			r, err := unpackClientSendData(c)
			if err != nil {
				log.Println(err)
				return
			}
			m.ClientCommand <- *r
		}
	}
}

func unpackClientSendData(c *websocket.Conn) (*syncController.StationState, error) {
	_, msg, err := c.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("websocket read failed: %e", err)
	}
	var ud clientSendData
	err = json.Unmarshal(msg, &ud)
	if err != nil {
		return nil, fmt.Errorf("bad json format: %e", err)
	}

	station, ok := pb.Stations_StationId_value[ud.StationName]
	if !ok {
		return nil, fmt.Errorf("bad station format: %s", ud.StationName)
	}

	state, ok := pb.RequestSync_State_value[ud.State]
	if !ok {
		return nil, fmt.Errorf("bad state format: %s", ud.State)
	}
	log.Printf("Received: StationID:%d, State:%d\n", station, state)
	return &syncController.StationState{
		StationState: servo.StationState{
			StationID: station,
			State:     state,
		},
	}, nil
}

//go:embed embed/index.html
var IndexHTML []byte

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(IndexHTML)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
