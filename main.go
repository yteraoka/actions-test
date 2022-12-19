package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	response_code := http.StatusOK
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(response_code)
	fmt.Fprintf(w, "OK\n")
}

func main() {
	http.HandleFunc("/", handler)
	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = "8080"
	}
	listenAddr := os.Getenv("LISTEN_ADDR")
	log.Printf("Listening %s:%s", listenAddr, listenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", listenAddr, listenPort), nil))
}
