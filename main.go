package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Start orvibo discovery
	go Start()

	// Setup http server
	r := mux.NewRouter()
	r.HandleFunc("/devices", devicesHandler)
	r.HandleFunc("/switch", switchHandler)

	srv := &http.Server{
		Handler: handlers.LoggingHandler(os.Stdout, r),
		Addr:    "0.0.0.0:8000",
	}
	// Start http server
	log.Fatal(srv.ListenAndServe())

}
