package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/ueckoken/backend/json2grpc/pkg/proxy"
	atspb "github.com/ueckoken/backend/json2grpc/spec"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Config struct {
	InternalEndpoint string `required:"true"`
	ListenAddr       string `default:":8080"`
}

func main() {
	var conf Config
	envconfig.MustProcess("", &conf)
	baseCtx := context.Background()
	certpool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalln("systemcertpool failed", err)
	}
	config := &tls.Config{
		RootCAs:    certpool,
		ServerName: conf.InternalEndpoint,
		MinVersion: tls.VersionTLS13,
	}
	cred := credentials.NewTLS(config)

	conn, err := grpc.Dial(
		conf.InternalEndpoint,
		grpc.WithTransportCredentials(cred),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Println("Connection failed in grpc conn initializing", err)
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
		if err := srv.ListenAndServe(); err != nil {
			log.Print(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	shutdownCtx, cancel := context.WithTimeout(baseCtx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Print(err)
	}
}
