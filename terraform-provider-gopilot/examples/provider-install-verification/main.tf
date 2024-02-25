terraform {
  required_providers {
    gopilot = {
      source = "klaxt.com/edu/gopilot"
    }
  }
}

provider "gopilot" {}

data "gopilot_devices" "example" {}
