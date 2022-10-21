package internal

import (
	"net/http"
	"ueckoken/plarail2022-positioning/pkg/addressChecker"
	"ueckoken/plarail2022-positioning/pkg/trainState"

	"context"
	"log"
	"time"
	pb "ueckoken/plarail2022-positioning/spec"

	"github.com/gorilla/websocket"
)

type ClientHandler struct {
	upgrader           websocket.Upgrader
	ClientNotification chan ClientNotifier
	Checker            addressChecker.AddressChecker
}

type ClientNotifier struct {
	Notifier   chan trainState.PositionAndSpeed
	Unregister chan struct{}
}

type ClientSendData struct {
	TrainName string  `json:"train_name"`
	HallName  string  `json:"hall_name"`
	Duration  float64 `json:"duration"`
	FetchedAt float64 `json:"fetched_at"`
}

func (m ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	ctx, cancel := context.WithCancel(context.Background())
	n := make(chan trainState.PositionAndSpeed)
	unregister := make(chan struct{})
	notifier := ClientNotifier{n, unregister}
	m.ClientNotification <- notifier
	c.SetPongHandler(func(string) error {
		return c.SetReadDeadline(time.Now().Add(20 * time.Second))
	})
	c.SetCloseHandler(func(code int, text string) error {
		log.Println("connection closed")
		cancel()
		return nil
	})
	go handleClientPing(ctx, c)
	for notification := range notifier.Notifier {
		data := ClientSendData{TrainName: pb.SendSpeed_Train_name[int32(notification.TrainID)], HallName: notification.HallSensorName, Duration: notification.Speed, FetchedAt: float64(notification.FetchedTimeStump.UnixMilli()) / 1000}
		err := c.WriteJSON(data)
		if err != nil {
			notifier.Unregister <- struct{}{}
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
				log.Println("ping:", err)
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}
