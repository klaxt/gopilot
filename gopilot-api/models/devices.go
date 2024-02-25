// models/device.go
package models

type Device struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Model  string `json:"model"`
	Color  string `json:"color"`
}
