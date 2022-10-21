package internal

import (
	"net/http"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"ueckoken/plarail2022-speed/pkg/storeSpeed"
	"ueckoken/plarail2022-speed/pkg/train2IP"
	pb "ueckoken/plarail2022-speed/spec"

	"github.com/gorilla/websocket"
)

type speedStruct struct {
	TrainName string `json:"train_name"`
	Speed     int    `json:"speed"`
}

func (m ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	ctx, cancel := context.WithCancel(context.Background())
	n := make(chan storeSpeed.TrainConf, 64)
	unregister := make(chan struct{})
	notifier := Client{ClientNotifier{n, unregister}}
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
	go handleClientCommand(ctx, c, &m)
	for notification := range notifier.notifier.Notifier {
		sendData := speedStruct{TrainName: notification.GetTrain().Name.String(), Speed: int(notification.GetSpeed())}
		err := c.WriteJSON(sendData)
		if err != nil {
			notifier.notifier.Unregister <- struct{}{}
			cancel()
			break
		}
	}
}

func handleClientCommand(ctx context.Context, c *websocket.Conn, m *ClientHandler) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			r, err := unpackClientSendData(c, m.TrainName2Id)
			if err != nil {
				log.Println(err)
				return
			}
			m.ClientCommand <- r
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

func unpackClientSendData(c *websocket.Conn, nameDir train2IP.Name2Id) (storeSpeed.TrainConf, error) {
	_, msg, err := c.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("websocket read failed: %e", err)
	}
	var ud speedStruct
	err = json.Unmarshal(msg, &ud)
	if err != nil {
		return nil, fmt.Errorf("bad json format: %e", err)
	}
	speed := storeSpeed.NewTrainConf(
		storeSpeed.NewTrain(pb.SendSpeed_Train(pb.SendSpeed_Train_value[ud.TrainName]),
			nameDir.SearchIP(ud.TrainName)),
	)
	if err := speed.SetSpeed(int32(ud.Speed)); err != nil {
		return nil, err
	}
	return speed, err
}
