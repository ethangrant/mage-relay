package handlers

import(
	"fmt"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dashboard display here.")
}