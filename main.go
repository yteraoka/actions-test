package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	version = "unknown"
	commit = "unknown"
	date = "unknown"
)

func handler(w http.ResponseWriter, r *http.Request) {
	response_code := http.StatusOK
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response_code)
	fmt.Fprintf(w, "{\"status\":\"ok\",\"version\":\"%s\",\"commit\":\"%s\",\"date\":\"%s\"}\n",
			version, commit, date)
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
