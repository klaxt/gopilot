package gopilot

type Device struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Model  string `json:"model"`
	Status string `json:"status"`
	Color  string `json:"color"`
}
