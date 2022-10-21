package msg2Esp

import (
	"encoding/json"
	"errors"
	"net/http"
	"ueckoken/plarail2022-internal/pkg/Send2Json"
	"ueckoken/plarail2022-internal/pkg/station2espIp"
)

type sendData struct {
	State string `json:"state"`
	Pin   int    `json:"pin"`
	Angle int    `json:"angle"`
}
type send2nodeExistAngle struct {
	Station  *station2espIp.StationDetail
	sendData *sendData
	client   *http.Client
}

type sendDataNoAngle struct {
	State string `json:"state"`
	Pin   int    `json:"pin"`
}

type send2nodeNoAngle struct {
	sendData *sendDataNoAngle
	Station  *station2espIp.StationDetail
	client   *http.Client
}

type Send2Node interface {
	Send() error
}

func NewSend2Node(c *http.Client, sta *station2espIp.StationDetail, state string, angle int) Send2Node {
	if sta.IsAngleDefined() {
		return &send2nodeExistAngle{
			Station: sta,
			client:  c,
			sendData: &sendData{
				State: "ANGLE",
				Pin:   sta.Pin,
				Angle: angle,
			},
		}
	}
	return &send2nodeNoAngle{
		Station: sta,
		client:  c,
		sendData: &sendDataNoAngle{
			State: state,
			Pin:   sta.Pin,
		},
	}
}

func (s *send2nodeExistAngle) Send() error {
	if s.sendData == nil {
		return errors.New("send data is nil")
	}

	b, err := json.Marshal(s.sendData)
	if err != nil {
		return err
	}
	sender := Send2Json.NewSendJSON(s.client, s.Station.Address, b)
	err = sender.Send()
	if err != nil {
		return err
	}
	return nil
}

func (s *send2nodeNoAngle) Send() error {
	if s.sendData == nil {
		return errors.New("send data is nil")
	}
	b, err := json.Marshal(s.sendData)
	if err != nil {
		return err
	}
	sender := Send2Json.NewSendJSON(s.client, s.Station.Address, b)
	err = sender.Send()
	if err != nil {
		return err
	}
	return nil
}
