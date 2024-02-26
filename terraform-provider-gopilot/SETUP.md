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

## Data Source

https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-data-source-read

```
go install .
TF_LOG=INFO terraform -chdir=examples/devices plan
```

## Create and Read

https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-resource-create

```
go install .
terraform -chdir=examples/device apply        
terraform -chdir=examples/device show
```

## Update

https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-resource-update


```
go install .
TF_LOG=INFO terraform -chdir=examples/device apply
```

Still some rough edges if server state changes outside context of terraform



curl -v -X POST -H "Content-Type: application/json" -d '{"name":"Device Test", "status":"active", "model":"iPad", "color":"Green"}' http://localhost:8000/api/devices 
curl -v -X PUT -H "Content-Type: application/json" -d '{"id":3, "name":"Device Test", "status":"active", "model":"iPad", "color":"Blue"}' http://localhost:8000/api/devices/3
curl -v -X DELETE http://localhost:8000/api/devices/2