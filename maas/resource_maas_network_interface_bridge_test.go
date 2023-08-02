package maas

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const TestACCResourceMaasNetworkInterfaceBridgeConfig_basic = `
resource "maas_network_interface_bridge" "test" {
	machine = "mq4s3r"
	name = "cloud-brmgmt"
	parent = "bond0.3342"
  }
  `

const TestACCResourceMaasNetworkInterfaceBridgeConfig_update = `
resource "maas_network_interface_bridge" "test" {
	machine = "mq4s3r"
	name = "cloud-brmgmt"
	parent = "bond0.3342"
	bridge_stp = true
  }
  `

func TestACCResourceMaasNetworkInterfaceBridge(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestACCResourceMaasNetworkInterfaceBridgeConfig_basic,
			},
			{
				Config: TestACCResourceMaasNetworkInterfaceBridgeConfig_update,
				// ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("maas_network_interface_bridge.test", "bridge_stp", "true"),
				),
			},
		},
	})
}
