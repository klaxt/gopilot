package gopilot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetCoffees - Returns list of coffees (no auth required)
func (c *Client) GetDevices() ([]Device, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/devices", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	devices := []Device{}
	err = json.Unmarshal(body, &devices)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

// GetCoffee - Returns specific coffee (no auth required)
func (c *Client) GetDevice(deviceId int64) (*Device, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/devices/%d", c.HostURL, deviceId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	device := Device{}
	err = json.Unmarshal(body, &device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

// CreateCoffee - Create new coffee
func (c *Client) CreateDevice(device Device) (*Device, error) {
	rb, err := json.Marshal(device)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/devices", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newDevice := Device{}
	err = json.Unmarshal(body, &newDevice)
	if err != nil {
		return nil, err
	}

	return &newDevice, nil
}

// UpdateOrder - Updates an order
func (c *Client) UpdateDevice(orderID int64, device Device) (*Device, error) {
	rb, err := json.Marshal(device)
	if err != nil {
		return nil, err
	}

	println(fmt.Sprintf("%s/api/devices%d", c.HostURL, orderID))
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/devices/%d", c.HostURL, orderID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	_, errr := c.doRequest(req)
	if errr != nil {
		return nil, errr
	}

	newDevice := Device{}
	// err = json.Unmarshal(body, &newDevice)
	// if err != nil {
	// 	return nil, err
	// }

	return &newDevice, nil
}

// DeleteDevice deletes a device with the specified order ID.
// It sends a DELETE request to the API endpoint and returns an error if any.
func (c *Client) DeleteDevice(orderID int64) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/devices/%d", c.HostURL, orderID), nil)
	if err != nil {
		return err
	}

	_, errr := c.doRequest(req)
	if err != nil {
		return errr
	}

	return nil
}
