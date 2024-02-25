terraform {
  required_providers {
    gopilot = {
      source = "klaxt.com/edu/gopilot"
    }
  }
}

provider "gopilot" {
  host     = "http://localhost:8000"
}

data "gopilot_devices" "example" {}

output "example_coffees" {
  value = data.gopilot_devices.example
}