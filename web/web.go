package web

import (
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func DefaultWeb() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8883", nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World! 1"))
	fmt.Println("ping", time.Now())
}
