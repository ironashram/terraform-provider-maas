package maas

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const TestACCDataMaasMachine_basic = `
resource "maas_machine" "virsh_vm1" {
	power_type = "virsh"
	hostname = "test-virsh-vm1"
	power_parameters = {
	  power_address = "qemu+ssh://ubuntu@10.113.1.26/system"
	  power_id = "test-vm1"
	}
	pxe_mac_address = "52:54:00:89:f5:3e"
  }
data "maas_machine" "virsh_vm1" {
	hostname = "test-virsh-vm1"
}
  `

const TestACCDataMaasMachine_pxe_mac = `
resource "maas_machine" "virsh_vm1" {
	power_type = "virsh"
	hostname = "test-virsh-vm1"
	power_parameters = {
	  power_address = "qemu+ssh://ubuntu@10.113.1.26/system"
	  power_id = "test-vm1"
	}
	pxe_mac_address = "52:54:00:89:f5:3e"
  }
data "maas_machine" "virsh_vm1" {
	pxe_mac_address = "52:54:00:89:f5:3e"
}
  `

func TestACCDataMaasMachine(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestACCDataMaasMachine_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.maas_machine.virsh_vm1", "pxe_mac_address", "52:54:00:89:f5:3e"),
				),
			},
			{
				Config: TestACCDataMaasMachine_pxe_mac,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.maas_machine.virsh_vm1", "hostname", "test-virsh-vm1"),
				),
			},
		},
	})
}
