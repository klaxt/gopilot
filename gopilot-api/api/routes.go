// api/routes.go
package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/hello", helloHandler).Methods("GET")
	r.HandleFunc("/api/devices", GetDevicesHandler).Methods("GET")
	r.HandleFunc("/api/devices/{id}", GetDeviceHandler).Methods("GET")
	r.HandleFunc("/api/devices", CreateDeviceHandler).Methods("POST")
	r.HandleFunc("/api/devices/{id}", DeleteDeviceHandler).Methods("DELETE")
	r.HandleFunc("/api/devices/{id}", UpdateDeviceHandler).Methods("PUT")
	// hack way to get into error - how to actually wildcard?
	r.HandleFunc("/api/{error}", errorHandler).Methods("GET", "POST", "PUT", "DELETE")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	// add log message to the console
	log.Printf("errorHandler called for %s", r.URL.Path)
	// return 404 not found error
	http.Error(w, "404 Not Found", http.StatusNotFound)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
