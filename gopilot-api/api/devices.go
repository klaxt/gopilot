// api/devices.go
package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gopilot/models"

	"github.com/gorilla/mux"
)

var devices = []models.Device{
	{ID: 1, Name: "Device 1", Status: "active"},
}

func GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(devices)
}

func GetDeviceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range devices {
		i, _ := strconv.Atoi(params["id"])
		if item.ID == i {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Device{})
}

func CreateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	_ = json.NewDecoder(r.Body).Decode(&device)
	if device.Model != "iPad" && device.Model != "iPhone" && device.Model != "Mac" {
		http.Error(w, "Invalid model", http.StatusBadRequest)
		return
	}
	devices = append(devices, device)
	json.NewEncoder(w).Encode(device)
}

func DeleteDeviceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range devices {
		i, _ := strconv.Atoi(params["id"])
		if item.ID == i {
			devices = append(devices[:index], devices[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(devices)
}

func UpdateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var newDevice models.Device
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&newDevice)
	if newDevice.Model != "iPad" && newDevice.Model != "iPhone" && newDevice.Model != "Mac" {
		http.Error(w, "Invalid model", http.StatusBadRequest)
		return
	}
	for index, item := range devices {
		i, _ := strconv.Atoi(params["id"])
		if item.ID == i {
			devices[index] = newDevice
			break
		}
	}
	json.NewEncoder(w).Encode(devices)
}
