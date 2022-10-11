package sendSpeed

import (
	"bytes"
	"encoding/json"
	"net/http"
	"ueckoken/plarail2021-soft-speed/pkg/storeSpeed"
)

type SendSpeed struct {
	Train  storeSpeed.TrainConf
	client *http.Client
}

type sendElements struct {
	Speed int32 `json:"speed"`
}

func NewSendSpeed(client *http.Client) *SendSpeed {
	return &SendSpeed{
		client: client,
	}
}
func (s *SendSpeed) Send() error {
	return s.sendJson()
}
func (s *SendSpeed) sendJson() error {
	b, err := s.getJson()
	if err != nil {
		return err
	}
	_, err = s.client.Post(s.Train.GetTrain().Addr, "application/json; charset=utf-8", bytes.NewReader(b))
	return err
}
func (s *SendSpeed) getJson() ([]byte, error) {
	payload := sendElements{Speed: s.Train.GetSpeed()}
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return b, nil
}
