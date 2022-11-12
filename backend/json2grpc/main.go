package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/ueckoken/backend/json2grpc/pkg/proxy"
	atspb "github.com/ueckoken/backend/json2grpc/spec"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	GRPCEndpoint string `required:"true"`
	ListenAddr       string `required:"true"`
}

func main() {
	var conf Config
	envconfig.MustProcess("", &conf)
	ctx := context.Background()
	ctxGrpcConn, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(
		ctxGrpcConn,
		conf.GRPCEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithBlock(), // 接続が確立するまでブロッキングする
	)
	if err != nil {
		log.Println("Connection failed")
		return
	}
	defer conn.Close()
	client := atspb.NewAtsClient(conn)
	mux := http.NewServeMux()
	mux.Handle("/sensor", proxy.NewHandler(client))
	srv := &http.Server{
		Addr:              conf.ListenAddr,
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	go func() {
		log.Printf("listen in %s\n", conf.ListenAddr)
		if err := srv.ListenAndServe(); err != nil {
			log.Print(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctxShutDown, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctxShutDown); err != nil {
		log.Print(err)
	}
}
