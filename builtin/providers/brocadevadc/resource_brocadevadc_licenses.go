package brocadevadc

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLicenses() *schema.Resource {
	return &schema.Resource{
		Create: resourceLicensesCreate,
		Read:   resourceLicensesRead,
		Update: resourceLicensesUpdate,
		Delete: resourceLicensesDelete,

		Schema: map[string]*schema.Schema{
			"accepting_delay": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  50,
			},
			"accepting_delays": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  50,
			},
		},
	}
}

func resourceLicensesCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceLicensesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceLicensesUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceLicensesDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
