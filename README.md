# GoPilot

Experimenting with Copilot while learning Go and Terraform providers

## API

API built entirely with Copilot

## Provider

Copilot not doing well at setting up a provider, took a few attempts to get to use the Terraform Plugin Framework instead of SKD. After manually creating a go project

Don't know enough here to understand what is wrong so is a combination of asking to generate code and pull from https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider

- Setup
    - `I want to add a terraform provider for the gopilot device api using the terraform plugin framework`
- Use Plugin Framework
    - `Use the Terraform Plugin Framework instead of the SDK`
    - This generated code imports and references that were wrong

## TODO

What is correct setup of `go.work`?