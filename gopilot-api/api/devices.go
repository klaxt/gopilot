// api/devices.go
package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"gopilot/models"

	"github.com/gorilla/mux"
)

var devices = []models.Device{
	{ID: 1, Name: "Device 1", Status: "active", Model: "iPad", Color: "white"},
}

func GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetDevicesHandler called")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devices)
}

func GetDeviceHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetDeviceHandler called")
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	log.Printf("GetDeviceHandler %d", i)
	for _, item := range devices {
		if item.ID == i {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Device{})
}

func CreateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateDeviceHandler called")
	var device models.Device
	_ = json.NewDecoder(r.Body).Decode(&device)
	if device.Model != "iPad" && device.Model != "iPhone" && device.Model != "Mac" {
		http.Error(w, "Invalid model", http.StatusBadRequest)
		return
	}
	device.ID = len(devices) + 1
	devices = append(devices, device)
	json.NewEncoder(w).Encode(device)
}

func DeleteDeviceHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteDeviceHandler called")
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	log.Printf("DeleteDeviceHandler %d", i)
	for index, item := range devices {
		if item.ID == i {
			devices = append(devices[:index], devices[index+1:]...)
			break
		}
	}
	// json.NewEncoder(w).Encode(devices)
}

func UpdateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateDeviceHandler called")
	var newDevice models.Device
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	log.Printf("UpdateDeviceHandler %d", i)
	_ = json.NewDecoder(r.Body).Decode(&newDevice)
	if newDevice.Model != "iPad" && newDevice.Model != "iPhone" && newDevice.Model != "Mac" {
		http.Error(w, "Invalid model", http.StatusBadRequest)
		return
	}
	for index, item := range devices {
		if item.ID == i {
			devices[index] = newDevice
			break
		}
	}
	json.NewEncoder(w).Encode(devices)
}
