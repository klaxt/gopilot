# Setup

## Provider

**It is required to be named terraform-provider-NAME**

https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider

```
go install .
terraform -chdir=examples/provider-install-verification plan 
```

## Configure Provider 

https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider-configure

```
go install .
terraform -chdir=examples/devices plan
```