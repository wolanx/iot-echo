package web

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func DefaultWeb() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8883", nil)
	log.Fatal(err)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World! 1"))
	fmt.Println("ping", time.Now())
}
