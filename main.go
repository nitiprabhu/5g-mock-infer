package main

import (
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	r   = mux.NewRouter()
	adr = os.Getenv("ADDRESS")
)

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	b := RandBool()
	if b {
		responseWriter.Write([]byte("true"))
	} else {
		responseWriter.Write([]byte("false"))

	}
}

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func main() {
	r.HandleFunc("/predict", handler).Methods("GET")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
