// api/routes.go
package api

import (
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
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
