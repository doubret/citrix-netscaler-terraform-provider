# netscaler_csaction

Terraform resource name : ```netscaler_csaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|name|No|No|string|
|targetlbvserver|No|No|[lbvserver.name](/doc/resources/lbvserver.md)|
|targetvserver|No|No|string|
|targetvserverexpr|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/content-switching/csaction/csaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_csaction" "<resource_name>" {

    comment = "abc"
    name = "abc"
    targetlbvserver = "${netscaler_lbvserver.<resource_name>.name}"
    targetvserver = "abc"
    targetvserverexpr = "abc"
}
```

