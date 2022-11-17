package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ueckoken/plarail2022/infra/logviewer/client"
)

type logHandler struct {
	password  string
	logviewer client.LogViewer
}

func isAllowed(appname string) bool {
	return true
}

func (l logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	v := r.URL.Query()
	if v == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "bad request")
		return
	}
	if v.Get("passwd") != l.password {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "password mismatch")
	}
	an := v.Get("app")
	if !isAllowed(an) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("got", an)
		fmt.Fprintln(w, "this appname is forbidden")
		return
	}
	logs, err := l.logviewer.GetLog(ctx, an, "plarail2021")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, logs)
}

func main() {
	passwd := os.Getenv("PASSWD")
	if passwd == "" {
		log.Fatalln("please specify password with PASSWD")
	}
	c, err := client.NewClient()
	if err != nil {
		log.Fatalln("failed to create a client", err)
	}
	l := logHandler{password: passwd, logviewer: c}
	http.Handle("/", l)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("failed to serve", err)
	}
}
