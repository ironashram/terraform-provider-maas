package maas

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/maas/gomaasclient/client"
)

func dataSourceMaasMachine() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMachineRead,

		Schema: map[string]*schema.Schema{
			"power_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"power_parameters": {
				Type:      schema.TypeMap,
				Computed:  true,
				Sensitive: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pxe_mac_address": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"hostname", "pxe_mac_address"},
			},
			"architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_hwe_kernel": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"hostname", "pxe_mac_address"},
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceMachineRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)
	var identifier string

	if hostname, ok := d.Get("hostname").(string); ok {
		identifier = hostname
	}

	if identifier == "" {
		if pxeMacAddress, ok := d.Get("pxe_mac_address").(string); ok {
			identifier = pxeMacAddress
		}
	}

	machine, err := getMachine(client, identifier)
	if err != nil {
		return diag.FromErr(err)
	}
	powerParams, err := client.Machine.GetPowerParameters(machine.SystemID)
	if err != nil {
		return diag.FromErr(err)
	}
	tfState := map[string]interface{}{
		"id":               machine.SystemID,
		"hostname":         machine.Hostname,
		"power_type":       machine.PowerType,
		"power_parameters": powerParams,
		"pxe_mac_address":  machine.BootInterface.MACAddress,
		"architecture":     machine.Architecture,
	}
	if err := setTerraformState(d, tfState); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
