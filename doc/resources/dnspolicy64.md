# netscaler_dnspolicy64

Terraform resource name : ```netscaler_dnspolicy64```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[dnsaction64.actionname](/doc/resources/dnsaction64.md)|
|name|No|No|string|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/domain-name-service/dnspolicy64/dnspolicy64/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_dnspolicy64" "<resource_name>" {

    action = "${netscaler_dnsaction64.<resource_name>.actionname}"
    name = "abc"
    rule = "abc"
}
```

