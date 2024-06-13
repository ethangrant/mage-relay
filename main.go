package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/ethangrant/mage-relay/handlers"
	"github.com/ethangrant/mage-relay/socket"
)

var AppRelay = socket.NewRelay()

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "mage-relay")
}

func main() {
	relay := socket.NewRelay()

	http.HandleFunc("/", handlers.Dashboard)
	http.HandleFunc("/dashboard", handlers.Dashboard)
	http.HandleFunc("/relay", func(w http.ResponseWriter, r *http.Request) {
		handlers.Relay(relay, w, r)
	})
	log.Fatal(http.ListenAndServe(":8083", nil))
}