# netscaler_lbvserver_tmtrafficpolicy_binding

Terraform resource name : ```netscaler_lbvserver_tmtrafficpolicy_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|gotopriorityexpression|No|No|string|
|invoke|No|No|bool|
|labelname|No|No|string|
|labeltype|No|No|reqvserver, resvserver, policylabel|
|name|No|No|[lbvserver.name](/doc/resources/lbvserver.md)|
|policyname|No|No|[tmtrafficpolicy.name](/doc/resources/tmtrafficpolicy.md)|
|priority|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbvserver_tmtrafficpolicy_binding/lbvserver_tmtrafficpolicy_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_lbvserver_tmtrafficpolicy_binding" "<resource_name>" {

    gotopriorityexpression = "abc"
    invoke = true
    labelname = "abc"
    labeltype = "reqvserver"
    name = "${netscaler_lbvserver.<resource_name>.name}"
    policyname = "${netscaler_tmtrafficpolicy.<resource_name>.name}"
    priority = 42
}
```

