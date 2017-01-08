package brocadevadc

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"log"
)

func resourceGlobalAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlobalAdminCreate,
		Read:   resourceGlobalAdminRead,
		Update: resourceGlobalAdminUpdate,
		Delete: resourceGlobalAdminDelete,

		Schema: map[string]*schema.Schema{
			"honor_fallback_scsv": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssl3_allow_rehandshake": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "rfc5746",
			},
			"ssl3_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_RSA_WITH_AES_128_GCM_SHA256,SSL_RSA_WITH_AES_128_CBC_SHA256,SSL_RSA_WITH_AES_128_CBC_SHA,SSL_RSA_WITH_AES_256_GCM_SHA384,SSL_RSA_WITH_AES_256_CBC_SHA256,SSL_RSA_WITH_AES_256_CBC_SHA,SSL_RSA_WITH_3DES_EDE_CBC_SHA,SSL_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,SSL_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,SSL_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,SSL_DHE_DSS_WITH_AES_128_CBC_SHA,SSL_DHE_DSS_WITH_AES_256_CBC_SHA,SSL_DHE_DSS_WITH_3DES_EDE_CBC_SHA",
			},
			"ssl3_diffie_hellman_key_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "dh_2048",
			},
			"ssl3_min_rehandshake_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"ssl_elliptic_curves": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ssl_insert_extra_fragment": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssl_max_handshake_message_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10240,
			},
			"ssl_prevent_timing_side_channels": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssl_signature_algorithms": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"support_ssl2": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"support_ssl3": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"support_tls1": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"support_tls11": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"support_tls12": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func mapGlobalAdminType(d *schema.ResourceData) *Globals {
	var elliptic []string
	//if d.Get("ssl_elliptic_curves") != nil {
	//	for _, e := range d.Get("ssl_elliptic_curves").([]interface{}) {
	//		elliptic = append(elliptic, e.(string))
	//	}
	//}

	return &Globals{
		Properties: &Properties{
			GlobalAdmin: &GlobalAdmin{
				HonorFallbackScsv:            d.Get("honor_fallback_scsv").(bool),
				Ssl3AllowRehandshake:         d.Get("ssl_allow_rehandshake").(string),
				Ssl3Ciphers:                  d.Get("ssl3_ciphers").(string),
				Ssl3DiffieHellmanKeyLength:   d.Get("ssl3_diffie_hellman_key_length").(string),
				Ssl3MinRehandshakeInterval:   d.Get("ssl3_min_rehandshake_interval").(int),
				SslEllipticCurves:            elliptic,
				SslInsertExtraFragment:       d.Get("ssl_insert_extra_fragment").(bool),
				SslMaxHandshakeMessageSize:   d.Get("ssl_max_handshake_message_size").(int),
				SslPreventTimingSideChannels: d.Get("ssl_prevent_timing_side_channels").(bool),
				SslSignatureAlgorithms:       d.Get("ssl_signature_algorithms").(string),
				SupportSsl2:                  d.Get("support_ssl2").(bool),
				SupportSsl3:                  d.Get("support_ssl3").(bool),
				SupportTls1:                  d.Get("support_tls1").(bool),
				SupportTls11:                 d.Get("Support_tls11").(bool),
				SupportTls12:                 d.Get("support_tls12").(bool),
			},
		},
	}
}

func resourceGlobalAdminCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*ClientConfig)
	global_admin := mapGlobalAdminType(d)

	log.Printf("[DEBUG] mapGlobalAdminType: %+v \n", global_admin)

	jsonpayload := jsonEncoder(global_admin)

	system_req, err := client.Put(fmt.Sprintf("%s/global_settings", endpoint), jsonpayload)

	log.Printf("[DEBUG] system_req status code: %+v\n", system_req.StatusCode)
	io, _ := ioutil.ReadAll(system_req.Body)
	log.Printf("[DEBUG] systaem_req body: %+v \n", string(io))

	if err != nil {
		return err
	}

	var globals Globals
	decoder := json.NewDecoder(system_req.Body)
	err = decoder.Decode(&globals)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return resourceGlobalAdminRead(d, m)
}

func resourceGlobalAdminRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*ClientConfig)
	global_system_req, err := client.Get(fmt.Sprintf("%s/global_settings", endpoint))

	if err != nil {
		return err
	}

	var global Globals

	decoder := json.NewDecoder(global_system_req.Body)
	err = decoder.Decode(&global)
	if err != nil {
		return err
	}

	d.Set("honor_fallback_scsv", global.Properties.GlobalAdmin.HonorFallbackScsv)
	d.Set("ssl_allow_rehandshake", global.Properties.GlobalAdmin.Ssl3AllowRehandshake)
	d.Set("ssl3_ciphers", global.Properties.GlobalAdmin.Ssl3Ciphers)
	d.Set("ssl3_diffie_hellman_key_length", global.Properties.GlobalAdmin.Ssl3DiffieHellmanKeyLength)
	d.Set("ssl3_min_rehandshake_interval", global.Properties.GlobalAdmin.Ssl3MinRehandshakeInterval)
	d.Set("ssl_elliptic_curves", global.Properties.GlobalAdmin.SslEllipticCurves)
	d.Set("ssl_insert_extra_fragment", global.Properties.GlobalAdmin.SslInsertExtraFragment)
	d.Set("ssl_max_handshake_message_size", global.Properties.GlobalAdmin.SslMaxHandshakeMessageSize)
	d.Set("ssl_prevent_timing_side_channels", global.Properties.GlobalAdmin.SslPreventTimingSideChannels)
	d.Set("ssl_signature_algorithms", global.Properties.GlobalAdmin.SslSignatureAlgorithms)
	d.Set("support_ssl2", global.Properties.GlobalAdmin.SupportSsl2)
	d.Set("support_ssl3", global.Properties.GlobalAdmin.SupportSsl3)
	d.Set("support_tls1", global.Properties.GlobalAdmin.SupportTls1)
	d.Set("Support_tls11", global.Properties.GlobalAdmin.SupportTls11)
	d.Set("support_tls12", global.Properties.GlobalAdmin.SupportTls12)
	return nil
}

func resourceGlobalAdminUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceGlobalAdminDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
