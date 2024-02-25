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

	body, err := c.doRequest(req, nil)
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
func (c *Client) GetDevice(coffeeID string) ([]Device, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/devices/%s", c.HostURL, coffeeID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
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

// CreateCoffee - Create new coffee
func (c *Client) CreateCoffee(coffee Device, authToken *string) (*Device, error) {
	rb, err := json.Marshal(coffee)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/coffees", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newCoffee := Device{}
	err = json.Unmarshal(body, &newCoffee)
	if err != nil {
		return nil, err
	}

	return &newCoffee, nil
}
