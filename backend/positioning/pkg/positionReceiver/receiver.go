package positionReceiver

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"ueckoken/plarail2022-positioning/pkg/addressChecker"
	"ueckoken/plarail2022-positioning/pkg/hallsensor"
	"ueckoken/plarail2022-positioning/pkg/trainState"
)

type PositionReceiveHandler struct {
	registerReceivedPosition chan trainState.State
	checker                  *addressChecker.AddressChecker
}

func NewPositionReceiverHandler(registerReceivedPosition chan trainState.State) PositionReceiveHandler {
	checker, err := addressChecker.NewAddressChecker()
	if err != nil {
		log.Fatalln(err)
	}
	return PositionReceiveHandler{
		registerReceivedPosition: registerReceivedPosition,
		checker:                  checker,
	}
}

type ReceivedPosition struct {
	MacAddress string `json:"mac_address"`
	Pin        int    `json:"pin"`
	TrainID    int    `json:"train_id"`
}

func (p PositionReceiveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip := getClientIP(r)
	if !p.checker.CheckIfOk(ip) {
		w.WriteHeader(http.StatusForbidden)
	}
	if r.Method != "POST" {
		if _, err := w.Write([]byte("Use POST")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var buf []byte
	if _, err := r.Body.Read(buf); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var receivedPosition ReceivedPosition
	if err := json.Unmarshal(buf, &receivedPosition); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h, err := hallsensor.NewEsp32PinSetting()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	name, err := h.Search(receivedPosition.MacAddress, receivedPosition.Pin)
	if err != nil {
		if _, err := w.Write([]byte("not found such address or pin")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//TODO: validate trainId
	dat := trainState.State{
		TrainID:          receivedPosition.TrainID,
		HallSensorName:   name,
		FetchedTimeStump: time.Now(),
	}
	p.registerReceivedPosition <- dat
	if _, err := w.Write([]byte("accepted")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getClientIP(r *http.Request) string {
	f := r.Header.Get("X-FORWARDED-FOR")
	if f != "" {
		return f
	}
	return r.RemoteAddr
}
