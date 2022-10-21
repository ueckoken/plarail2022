package Send2Json

import (
	"bytes"
	"fmt"
	"net/http"
)

type SendJSON struct {
	client *http.Client
	url    string
	json   []byte
}

func NewSendJSON(client *http.Client, url string, json []byte) *SendJSON {
	return &SendJSON{
		client: client,
		url:    url,
		json:   json,
	}
}
func (sj *SendJSON) Send() error {
	res, err := sj.client.Post(sj.url, "application/json", bytes.NewBuffer(sj.json))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if !(200 <= res.StatusCode && res.StatusCode < 300) {
		return fmt.Errorf("http status is not succeed. status is %s ", res.Status)
	}
	return nil
}
