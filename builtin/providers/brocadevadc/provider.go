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
				Sensitive:   true,
			},
			"ssl_ignore_verify": {
				Optional:    true,
				Type:        schema.TypeBool,
				DefaultFunc: schema.EnvDefaultFunc("VADC_VERIFY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"brocadevadc_global_settings": resourceGlobalSettings(),
			"brocadevadc_licenses":        resourceLicenses(),
			"brocadevadc_ssl":             resourceSSL(),
			"brocadevadc_virtual_server":  resourceVirtualServer(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &ClientConfig{
		URL:       d.Get("url").(string),
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
		SslVerify: d.Get("ssl_ignore_verify").(bool),
	}
	return config, nil
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"url": "The URL used to connect to the Brocade vADC administration interface\n" +
			"Example: https://127.0.0.1:9070",

		"username": "The Username used to authenticate to the API as.\n" +
			"Example: admin",

		"password": "The corresponding users password.\n" +
			"Example: ?",

		"ssl_ignore_verify": "Verify the SSL certificate as being authentic.\n" +
			"Example: false",
	}
}
