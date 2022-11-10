package proxy

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	atspb "github.com/ueckoken/backend/json2grpc/spec"
)

type Handler struct {
	grpcClient atspb.AtsClient
}

type StatusResponse struct {
	Status        string `json:"status"`
	InternalError error  `json:"internalerror,omitempty"`
}

type Send struct {
	Sensor int32 `validate:"required,min=1,max=52"`
}

func NewHandler(client atspb.AtsClient) *Handler {
	return &Handler{grpcClient: client}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		var sensor Send
		if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
			if err := json.NewEncoder(w).Encode(StatusResponse{Status: "err", InternalError: err}); err != nil {
				log.Println(err)
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		validate := validator.New()
		if err := validate.Struct(sensor); err != nil {
			if err := json.NewEncoder(w).Encode(StatusResponse{Status: "err", InternalError: err}); err != nil {
				log.Println(err)
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("%+v\n", sensor)
		req := &atspb.SendStatusRequest{
			Sensor: atspb.SendStatusRequest_SensorName(sensor.Sensor),
		}
		res, err := h.grpcClient.SendStatus(r.Context(), req)
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(StatusResponse{Status: res.GetResponse().String()}); err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, `{"status":"permits only POST"}`, http.StatusMethodNotAllowed)
	}
}
