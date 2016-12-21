package brocadevadc

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("VADC_URL", nil),
			},
			"username": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("VADC_USERNAME", nil),
			},
			"password": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("VADC_PASSWORD", nil),
			},
			"verify_ssl": {
				Required:    true,
				Type:        schema.TypeBool,
				DefaultFunc: schema.EnvDefaultFunc("VADC_VERIFY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"brocadevadc_global":         resourceGlobal(),
			"brocadevadc_licenses":       resourceLicenses(),
			"brocadevadc_ssl":            resourceSSL(),
			"brocadevadc_virtual_server": resourceVirtualServer(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := ClientConfig{
		URL:       d.Get("url").(string),
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
		VerifySSL: d.Get("ssl_verify").(bool),
	}
	return config, nil
}
