package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	atspb "github.com/ueckoken/backend/json2grpc/spec"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Send struct {
	Sensor int32 `validate:"required,min=1,max=52"`
}

func main() {
	var client atspb.AtsClient

	//grpcサーバーとのコネクションの確立
	address := "localhost:8888" //grpcサーバーのアドレス
	conn, err := grpc.Dial(
		address,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed")
		return
	}
	defer conn.Close()
	client = atspb.NewAtsClient(conn)

	http.HandleFunc("/sensor", func(w http.ResponseWriter, r *http.Request) {
		var mutex = &sync.RWMutex{}
		var sensor Send

		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodPost:
			var temp Send
			if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}
			validate := validator.New()
			if err := validate.Struct(temp); err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusBadRequest)
				return
			}
			mutex.Lock()
			sensor = temp
			fmt.Printf("%+v\n", sensor)

			req := &atspb.SendStatusRequest{
				Sensor: atspb.SendStatusRequest_SensorName(sensor.Sensor),
			}

			res, err := client.SendStatus(context.Background(), req)
			if err != nil {
				fmt.Println(err)
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			} else {
				fmt.Println(res.String())
			}
			mutex.Unlock()

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status": "ok"}`))

		default:
			http.Error(w, `{"status":"permits only POST"}`, http.StatusMethodNotAllowed)
		}
	})

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
