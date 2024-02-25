// cmd/main.go
package main

import (
	"log"
	"net/http"

	"gopilot/api"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	api.RegisterRoutes(r)

	log.Fatal(http.ListenAndServe(":8000", r))
}
