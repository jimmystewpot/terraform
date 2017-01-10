package brocadevadc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceGlobalSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlobalSettingsCreate,
		Read:   resourceGlobalSettingsRead,
		Update: resourceGlobalSettingsUpdate,
		Delete: resourceGlobalSettingsDelete,

		Schema: map[string]*schema.Schema{
			// Start global_settings/basic
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
			// Start global_settings/admin {}
			"honor_fallback_scsv": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssl3_allow_rehandshake": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "rfc5746",
				ValidateFunc: validateSsl3Handshake,
			},
			"ssl3_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_RSA_WITH_AES_128_GCM_SHA256,SSL_RSA_WITH_AES_128_CBC_SHA256,SSL_RSA_WITH_AES_128_CBC_SHA,SSL_RSA_WITH_AES_256_GCM_SHA384,SSL_RSA_WITH_AES_256_CBC_SHA256,SSL_RSA_WITH_AES_256_CBC_SHA,SSL_RSA_WITH_3DES_EDE_CBC_SHA,SSL_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,SSL_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,SSL_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,SSL_DHE_DSS_WITH_AES_128_CBC_SHA,SSL_DHE_DSS_WITH_AES_256_CBC_SHA,SSL_DHE_DSS_WITH_3DES_EDE_CBC_SHA",
			},
			"ssl3_diffie_hellman_key_length": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "dh_2048",
				ValidateFunc: validateSS3diffieHellmanKl,
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
				Default:  true,
			},
			// start global_settings/appliance
			"bootloader_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"manage_ncipher": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"nethsm_esn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nethsm_hash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nethsm_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nethsm_ncipher_rfs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"return_path_routing_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// Start global_settings/aptimizer
			"max_dependent_fetch_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "2MB",
			},
			"max_original_content_buffer_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "2MB",
			},
			"watchdog_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"watchdog_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
			// Start global_settings/auditlog
			"via_eventd": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"via_syslog": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// Start global_settings/autoscaler
			"autoscaler_verbose": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// Start global_settings/bgp
			"bgp_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"as_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  65534,
			},
			// Start global_settings/cluster_comms
			"allow_update_default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"allowed_update_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state_sync_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
			"state_sync_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
			// Start global_settings/connection
			"idle_connections_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"listen_queue_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_accepting": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"multiple_accept": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// Start global_settings/data_pane_acceleration
			"tcp_delay_ack": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_win_scale": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Start global_settings/dns
			"max_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"min_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"negative_expiry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10867,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  12,
			},
			// Start global_settings/ec2
			"access_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"awstool_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"secret_access_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"verify_query_server_cert": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// Start global_settings/eventing
			"mail_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"max_attempts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			// Start global_settings/fault_tolerance
			"arp_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"auto_failback": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"child_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"frontend_check_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 1,
				Elem: schema.Schema{
					Type: schema.TypeString},
			},
			"heartbeat_method": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "unicast",
				ValidateFunc: validateHeartBeatMethod,
			},
			"igmp_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"l4accel_child_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"l4accel_sync_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10240,
			},
			"monitor_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  500,
			},
			"monitor_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"multicast_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "239.100.1.1:9090",
			},
			"unicast_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9090,
			},
			"use_bind_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fault_tolerance_verbose": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// Start global_settings/fips
			"fips_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// Start global_settings/ftp
			"data_bind_low": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// Start global_settings/glb
			"glb_verbose": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// Start global_settings/historical_activity
			"keep_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  90,
			},
		},
	}
}

func mapGlobalSettingsType(d *schema.ResourceData) *Globals {
	var licenses []string
	var elliptic []string
	var allowed_update_hosts []string
	var frontend_check_ips []string

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
			GlobalAdmin: &GlobalAdmin{
				HonorFallbackScsv:            d.Get("honor_fallback_scsv").(bool),
				Ssl3AllowRehandshake:         d.Get("ssl3_allow_rehandshake").(string),
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
				SupportTls11:                 d.Get("support_tls11").(bool),
				SupportTls12:                 d.Get("support_tls12").(bool),
			},
			GlobalAppliance: &GlobalAppliance{
				BootloaderPassword:       d.Get("bootloader_password").(string),
				ManageNcipher:            d.Get("manage_ncipher").(bool),
				NethsmEsn:                d.Get("nethsm_esn").(string),
				NethsmHash:               d.Get("nethsm_hash").(string),
				NethsmIp:                 d.Get("nethsm_ip").(string),
				NethsmNcipherRfs:         d.Get("nethsm_ncipher_rfs").(string),
				ReturnPathRoutingEnabled: d.Get("return_path_routing_enabled").(bool),
			},
			GlobalAptimizer: &GlobalAptimizer{
				MaxDependentFetchSize:        d.Get("max_dependent_fetch_size").(string),
				MaxOriginalContentBufferSize: d.Get("max_original_content_buffer_size").(string),
				WatchdogInterval:             d.Get("watchdog_interval").(int),
				WatchdogLimit:                d.Get("watchdog_limit").(int),
			},
			GlobalAuditlog: &GlobalAuditlog{
				ViaEventd: d.Get("via_eventd").(bool),
				ViaSyslog: d.Get("via_syslog").(bool),
			},
			GlobalBandwidth: &GlobalBandwidth{},
			GlobalAutoScaler: &GlobalAutoScaler{
				Verbose: d.Get("autoscaler_verbose").(bool),
			},
			GlobalBgp: &GlobalBgp{
				AsNumber: d.Get("as_number").(int),
				Enabled:  d.Get("bgp_enabled").(bool),
			},
			GlobalClusterComms: &GlobalClusterComms{
				AllowUpdateDefault: d.Get("allow_update_default").(bool),
				AllowedUpdateHosts: allowed_update_hosts,
				StateSyncInterval:  d.Get("state_sync_interval").(int),
				StateSyncTimeout:   d.Get("state_sync_timeout").(int),
			},
			GlobalConnection: &GlobalConnection{
				IdleConnectionsMax: d.Get("idle_connections_max").(int),
				IdleTimeout:        d.Get("idle_timeout").(int),
				ListenQueueSize:    d.Get("listen_queue_size").(int),
				MaxAccepting:       d.Get("max_accepting").(int),
				MultipleAccept:     d.Get("multiple_accept").(bool),
			},
			GlobalDataPlaneAcceleration: &GlobalDataPlaneAcceleration{
				TcpDelayAck: d.Get("tcp_delay_ack").(int),
				TcpWinScale: d.Get("tcp_win_scale").(int),
			},
			GlobalDns: &GlobalDns{
				MaxTTL:         d.Get("max_ttl").(int),
				MinTTL:         d.Get("min_ttl").(int),
				NegativeExpiry: d.Get("negative_expiry").(int),
				Size:           d.Get("size").(int),
				Timeout:        d.Get("timeout").(int),
			},
			GlobalDnsAutoscale: &GlobalDnsAutoscale{},
			GlobalEc2: &GlobalEc2{
				AccessKeyID:           d.Get("access_key_id").(string),
				AwstoolTimeout:        d.Get("awstool_timeout").(int),
				SecretAccessKey:       d.Get("secret_access_key").(string),
				VerifyQueryServerCert: d.Get("verify_query_server_cert").(bool),
			},
			GlobalEventing: &GlobalEventing{
				MailInterval: d.Get("mail_interval").(int),
				MaxAttempts:  d.Get("max_attempts").(int),
			},
			GlobalFaultTolerance: &GlobalFaultTolerance{
				ArpCount:            d.Get("arp_count").(int),
				AutoFailback:        d.Get("auto_failback").(bool),
				ChildTimeout:        d.Get("child_timeout").(int),
				FrontendCheckIps:    frontend_check_ips,
				HeartbeatMethod:     d.Get("heartbeat_method").(string),
				IgmpInterval:        d.Get("igmp_interval").(int),
				L4AccelChildTimeout: d.Get("l4accel_child_timeout").(int),
				L4AccelSyncPort:     d.Get("l4accel_sync_port").(int),
				MonitorInterval:     d.Get("monitor_interval").(int),
				MonitorTimeout:      d.Get("monitor_timeout").(int),
				MulticastAddress:    d.Get("multicast_address").(string),
				UnicastPort:         d.Get("unicast_port").(int),
				UseBindIP:           d.Get("use_bind_ip").(bool),
				Verbose:             d.Get("fault_tolerance_verbose").(bool),
			},
			GlobalFips: &GlobalFips{
				Enabled: d.Get("fips_enabled").(bool),
			},
			GlobalFtp: &GlobalFtp{
				DataBindLow: d.Get("data_bind_low").(bool),
			},
			GlobalGlb: &GlobalGlb{
				Verbose: d.Get("glb_verbose").(bool),
			},
			GlobalHistoricalActivity: &GlobalHistoricalActivity{
				KeepDays: d.Get("keep_days").(int),
			},
		},
	}
}

func resourceGlobalSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ClientConfig)

	// Create a lock, there should be only one module running concurrently.
	client.Lock()
	defer client.Unlock()

	globalSettings := mapGlobalSettingsType(d)

	log.Printf("GlobalSettingsCreate mapGlobalSettingsType: %+v \n", *globalSettings)

	jsonpayload := jsonEncoder(globalSettings)

	system_req, err := client.Put(fmt.Sprintf("%s/global_settings", apipath), jsonpayload)

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

	return resourceGlobalSettingsRead(d, meta)
}

func resourceGlobalSettingsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ClientConfig)
	global_system_req, err := client.Get(fmt.Sprintf("%s/global_settings", apipath))

	if err != nil {
		return err
	}

	var global Globals

	decoder := json.NewDecoder(global_system_req.Body)
	err = decoder.Decode(&global)

	if jsonDecodeError(err) {
		return err
	}

	// global_settings/basic
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

	// global_settings/admin
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

	// global_settings/appliance
	d.Set("bootloader_password", global.Properties.GlobalAppliance.BootloaderPassword)
	d.Set("manage_ncipher", global.Properties.GlobalAppliance.ManageNcipher)
	d.Set("nethsm_esn", global.Properties.GlobalAppliance.NethsmEsn)
	d.Set("nethsm_hash", global.Properties.GlobalAppliance.NethsmHash)
	d.Set("nethsm_ip", global.Properties.GlobalAppliance.NethsmIp)
	d.Set("nethsm_ncipher_rfs", global.Properties.GlobalAppliance.NethsmNcipherRfs)
	d.Set("return_path_routing_enabled", global.Properties.GlobalAppliance.ReturnPathRoutingEnabled)

	// global_settings/aptimizer
	d.Set("max_dependent_fetch_size", global.Properties.GlobalAptimizer.MaxDependentFetchSize)
	d.Set("max_original_content_buffer_size", global.Properties.GlobalAptimizer.MaxOriginalContentBufferSize)
	d.Set("watchdog_interval", global.Properties.GlobalAptimizer.WatchdogInterval)
	d.Set("watchdog_limit", global.Properties.GlobalAptimizer.WatchdogLimit)

	// global_settings/audit_log
	d.Set("via_eventd", global.Properties.GlobalAuditlog.ViaEventd)
	d.Set("via_syslogd", global.Properties.GlobalAuditlog.ViaSyslog)

	// global_settings/autoscaler
	d.Set("autoscaler_verbose", global.Properties.GlobalAutoScaler.Verbose)

	// global_settings/bgp
	d.Set("as_number", global.Properties.GlobalBgp.AsNumber)
	d.Set("bgp_enabled", global.Properties.GlobalBgp.Enabled)

	// global_settings/cluster_comms
	d.Set("allow_update_default", global.Properties.GlobalClusterComms.AllowUpdateDefault)
	d.Set("allowed_update_hosts", global.Properties.GlobalClusterComms.AllowedUpdateHosts)
	d.Set("state_sync_interval", global.Properties.GlobalClusterComms.StateSyncInterval)
	d.Set("state_sync_timeout", global.Properties.GlobalClusterComms.StateSyncTimeout)

	// global_settings/connection
	d.Set("idle_connections_max", global.Properties.GlobalConnection.IdleConnectionsMax)
	d.Set("idle_timeout", global.Properties.GlobalConnection.IdleTimeout)
	d.Set("max_accepting", global.Properties.GlobalConnection.MaxAccepting)
	d.Set("multiple_accept", global.Properties.GlobalConnection.MultipleAccept)

	// global_settings/data_plane_acceleration
	d.Set("max_ttl", global.Properties.GlobalDns.MaxTTL)
	d.Set("min_ttl", global.Properties.GlobalDns.MinTTL)
	d.Set("negative_expiry", global.Properties.GlobalDns.NegativeExpiry)
	d.Set("size", global.Properties.GlobalDns.Size)
	d.Set("timeout", global.Properties.GlobalDns.Timeout)

	// global_settings/ec2
	d.Set("access_key_id", global.Properties.GlobalEc2.AccessKeyID)
	d.Set("awstool_timeout", global.Properties.GlobalEc2.AwstoolTimeout)
	d.Set("secret_access_key", global.Properties.GlobalEc2.SecretAccessKey)
	d.Set("verify_query_server_cert", global.Properties.GlobalEc2.VerifyQueryServerCert)

	// global_settings/eventing
	d.Set("mail_interval", global.Properties.GlobalEventing.MailInterval)
	d.Set("max_attempts", global.Properties.GlobalEventing.MaxAttempts)

	// global_settings/fault_tolerance
	d.Set("arp_count", global.Properties.GlobalFaultTolerance.ArpCount)
	d.Set("auto_failback", global.Properties.GlobalFaultTolerance.AutoFailback)
	d.Set("child_timeout", global.Properties.GlobalFaultTolerance.ChildTimeout)
	d.Set("frontend_check_ips", global.Properties.GlobalFaultTolerance.FrontendCheckIps)
	d.Set("heartbeat_method", global.Properties.GlobalFaultTolerance.HeartbeatMethod)
	d.Set("igmp_interval", global.Properties.GlobalFaultTolerance.IgmpInterval)
	d.Set("l4accel_child_timeout", global.Properties.GlobalFaultTolerance.L4AccelChildTimeout)
	d.Set("l4accel_sync_port", global.Properties.GlobalFaultTolerance.L4AccelSyncPort)
	d.Set("monitor_interval", global.Properties.GlobalFaultTolerance.MonitorInterval)
	d.Set("monitor_timeout", global.Properties.GlobalFaultTolerance.MonitorTimeout)
	d.Set("multicast_address", global.Properties.GlobalFaultTolerance.MulticastAddress)
	d.Set("unicast_port", global.Properties.GlobalFaultTolerance.UnicastPort)
	d.Set("use_bind_ip", global.Properties.GlobalFaultTolerance.UseBindIP)
	d.Set("fault_tolerance_verbose", global.Properties.GlobalFaultTolerance.Verbose)

	// global_settings/fips
	d.Set("fips_enabled", global.Properties.GlobalFips.Enabled)

	// global_settings/ftp
	d.Set("data_bind_low", global.Properties.GlobalFtp.DataBindLow)

	// Start global_settings/glb
	d.Set("glb_verbose", global.Properties.GlobalGlb.Verbose)

	// Start global_settings/historical_activity
	d.Set("keep_days", global.Properties.GlobalHistoricalActivity.KeepDays)

	return nil
}

func resourceGlobalSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceGlobalSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
