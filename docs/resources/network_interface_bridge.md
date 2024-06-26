---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "maas_network_interface_bridge Resource - terraform-provider-maas"
subcategory: ""
description: |-
  Provides a resource to manage MAAS network Bridges.
---

# maas_network_interface_bridge (Resource)

Provides a resource to manage MAAS network Bridges.

## Example Usage

```terraform
resource "maas_network_interface_bridge" "example" {
  machine = maas_machine.example.id
  name    = "example"
  parent  = maas_network_interface_vlan.example.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `machine` (String) The identifier (system ID, hostname, or FQDN) of the machine with the bridge interface.

### Optional

- `accept_ra` (Boolean) Accept router advertisements. (IPv6 only).
- `bridge_fd` (Number) Set bridge forward delay to time seconds. (Default: 15).
- `bridge_stp` (Boolean) Turn spanning tree protocol on or off. (Default: False).
- `bridge_type` (String) The type of bridge to create. Possible values are: ``standard``, ``ovs``.
- `mac_address` (String) The bridge interface MAC address.
- `mtu` (Number) The MTU of the bridge interface.
- `name` (String) The bridge interface name.
- `parent` (String) Parent interface name for this bridge interface.
- `tags` (Set of String) A set of tag names to be assigned to the bridge interface.
- `vlan` (Number) Database ID of the VLAN the bridge interface is connected to.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# A bridge network interface can be imported using the machine identifier (system ID, hostname, or FQDN) and ID. e.g.
$ terraform import maas_network_interface_bridge.example vm1:id
```
