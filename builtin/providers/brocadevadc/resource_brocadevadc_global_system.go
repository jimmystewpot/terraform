package brocadevadc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceGlobalSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlobalSystemCreate,
		Read:   resourceGlobalSystemRead,
		Update: resourceGlobalSystemUpdate,
		Delete: resourceGlobalSystemDelete,

		Schema: map[string]*schema.Schema{
			"accepting_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  50,
			},
			"afm_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
			"child_control_command_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"child_control_kill_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"chunk_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  16385,
			},
			"client_first_opt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cluster_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
			"cpu_starvation_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"cpu_starvation_check_tolerance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"data_plane_acceleration_cores": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "one",
				ValidateFunc: validateDataPlaneCores,
			},
			"data_plane_acceleration_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"http2_no_cipher_blacklist_check": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"internal_config_logging": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"license_servers": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"max_fds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1048576,
			},
			"monitor_memory_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
			"rate_class_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25000,
			},
			"shared_pool_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10MB",
			},
			"slm_class_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"so_rbuff_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"so_wbuff_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"socket_optimizations": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "auto",
			},
			"storage_shared": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tip_class_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
		},
	}
}

func mapGlobalSystemType(d *schema.ResourceData) *Globals {
	var licenses []string

	//if _, ok := d.GetOk("license_servers"); ok {
	//	for _, l := range d.Get("license_servers").(*schema.Set) {
	//		licenses = append(licenses, l.(string))
	//	}
	//}
	return &Globals{
		Properties: &Properties{
			GlobalBasic: &GlobalBasic{
				AcceptingDelay:              d.Get("accepting_delay").(int),
				AfmEnabled:                  d.Get("afm_enabled").(bool),
				ChildControlCommandTimeout:  d.Get("child_control_command_timeout").(int),
				ChildControlKillTimeout:     d.Get("child_control_kill_timeout").(int),
				ChunkSize:                   d.Get("chunk_size").(int),
				ClientFirstOpt:              d.Get("client_first_opt").(bool),
				ClusterIdentifier:           d.Get("cluster_identifier").(string),
				CpuStarvationCheckInterval:  d.Get("cpu_starvation_check_interval").(int),
				CpuStarvationCheckTolerance: d.Get("cpu_starvation_check_tolerance").(int),
				DataPlaneAccelerationMode:   d.Get("data_plane_acceleration_mode").(bool),
				DataPlaneAccelerationCores:  d.Get("data_plane_acceleration_cores").(string),
				Http2noCipherBlacklistCheck: d.Get("http2_no_cipher_blacklist_check").(bool),
				LicenseServers:              licenses,
				MaxFds:                      d.Get("max_fds").(int),
				MonitorMemorySize:           d.Get("monitor_memory_size").(int),
				RateClassLimit:              d.Get("rate_class_limit").(int),
				SharedPoolSize:              d.Get("shared_pool_size").(string),
				SlmClassLimit:               d.Get("slm_class_limit").(int),
				SoRbuffSize:                 d.Get("so_rbuff_size").(int),
				SoWbuffSize:                 d.Get("so_wbuff_size").(int),
				SocketOptimizations:         d.Get("socket_optimizations").(string),
				StorageShared:               d.Get("storage_shared").(bool),
				TipClassLimit:               d.Get("tip_class_limit").(int),
			},
		},
	}
}

func resourceGlobalSystemCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ClientConfig)

	// Create a lock, there should be only one module running concurrently.
	client.Mutex.Lock()
	defer client.Mutex.Unlock()

	global_system := mapGlobalSystemType(d)

	log.Printf("GlobalSystemCreate mapGlobalSystemType: %+v \n", *global_system)

	jsonpayload := jsonEncoder(global_system)

	system_req, err := client.Put(fmt.Sprintf("%s/global_settings", endpoint), jsonpayload)

	if !handleHttpCodes(system_req) {
		log.Printf("GlobalSystemCreate system_req status code: %+v\n", system_req.StatusCode)
		return err
	}

	// Read the returned JSON into a buffer
	RequestBuffer := new(bytes.Buffer)
	RequestBuffer.ReadFrom(system_req.Body)

	log.Printf("GlobalSystemCreate system_req body: %+v \n", RequestBuffer)

	if err != nil {
		return err
	}

	var vAdcReturnedData Globals
	err = json.NewDecoder(RequestBuffer).Decode(&vAdcReturnedData)
	if jsonDecodeError(err) {
		return err
	}

	d.SetId("global_settings/system")

	return resourceGlobalSystemRead(d, meta)
}

func resourceGlobalSystemRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ClientConfig)
	global_system_req, err := client.Get(fmt.Sprintf("%s/global_settings", endpoint))

	if err != nil {
		return err
	}

	var global Globals

	decoder := json.NewDecoder(global_system_req.Body)
	err = decoder.Decode(&global)

	if jsonDecodeError(err) {
		return err
	}

	d.Set("accepting_delay", global.Properties.GlobalBasic.AcceptingDelay)
	d.Set("afm_enabled", global.Properties.GlobalBasic.AfmEnabled)
	d.Set("child_control_command_timeout", global.Properties.GlobalBasic.ChildControlCommandTimeout)
	d.Set("child_control_kill_timeout", global.Properties.GlobalBasic.ChildControlKillTimeout)
	d.Set("chunk_size", global.Properties.GlobalBasic.ChunkSize)
	d.Set("client_first_opt", global.Properties.GlobalBasic.ClientFirstOpt)
	d.Set("cluster_identifier", global.Properties.GlobalBasic.ClusterIdentifier)
	d.Set("cpu_starvation_check_interval", global.Properties.GlobalBasic.CpuStarvationCheckInterval)
	d.Set("cpu_starvation_check_tolerance", global.Properties.GlobalBasic.CpuStarvationCheckTolerance)
	d.Set("data_plane_acceleration_mode", global.Properties.GlobalBasic.DataPlaneAccelerationMode)
	d.Set("data_plane_acceleration_cores", global.Properties.GlobalBasic.DataPlaneAccelerationCores)
	d.Set("http2_no_cipher_blacklist_check", global.Properties.GlobalBasic.Http2noCipherBlacklistCheck)
	d.Set("license_servers", global.Properties.GlobalBasic.LicenseServers)
	d.Set("max_fds", global.Properties.GlobalBasic.MaxFds)
	d.Set("monitor_memory_size", global.Properties.GlobalBasic.MonitorMemorySize)
	d.Set("rate_class_limit", global.Properties.GlobalBasic.RateClassLimit)
	d.Set("shared_pool_size", global.Properties.GlobalBasic.SharedPoolSize)
	d.Set("slm_class_limit", global.Properties.GlobalBasic.SlmClassLimit)
	d.Set("so_rbuff_size", global.Properties.GlobalBasic.SoRbuffSize)
	d.Set("so_wbuff_size", global.Properties.GlobalBasic.SoWbuffSize)
	d.Set("socket_optimizations", global.Properties.GlobalBasic.SocketOptimizations)
	d.Set("storage_shared", global.Properties.GlobalBasic.StorageShared)
	d.Set("tip_class_limit", global.Properties.GlobalBasic.TipClassLimit)

	return nil
}

func resourceGlobalSystemUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceGlobalSystemDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
