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


resource "gopilot_device" "iPad" {
  name = "edu"
  model = "iPad"
  status = "active"
  color = "red"
}

output "iPad" {
  value = gopilot_device.iPad
}
