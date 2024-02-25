# GoPilot API

Using Copilot chat to create project. It is building a project off it's internal chat state only? Not taking into account changes I make (e.g. file names)

- Setting up a new project.
    - `Start a new go project to build an API service`
- Adding our functionality
    - `Add CRUD API abilities to manage devices`
- Fix project name in generated code, was not using what was in `go.mod`
    - `I have renamed my project to gopilot`
- Adding device info
    - `Add fields to device for model and color`
    - `Model should be one of iPad, iPhone, or Mac`
        - Missed setting initial devices model


`go run ./cmd/main.go`

`curl -X POST -H "Content-Type: application/json" -d '{"id":"3", "name":"Device 3", "status":"active", "model":"iPad", "color":"Green"}' http://localhost:8000/api/devices`