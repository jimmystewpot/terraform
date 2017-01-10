package brocadevadc

// Globals JSON encoded composition.
// /api/tm/3.9/config/active/global_settings
type Globals struct {
	Properties *Properties `json:"properties"`
}

type Properties struct {
	GlobalBasic                 *GlobalBasic                 `json:"basic,omitempty"`
	GlobalAdmin                 *GlobalAdmin                 `json:"admin,omitempty"`
	GlobalAppliance             *GlobalAppliance             `json:"appliance,omitempty"`
	GlobalAptimizer             *GlobalAptimizer             `json:"aptimizer,omitempty"`
	GlobalAuditlog              *GlobalAuditlog              `json:"auditlog,omitempty"`
	GlobalAutoScaler            *GlobalAutoScaler            `json:"autoscaler,omitempty"`
	GlobalBandwidth             *GlobalBandwidth             `json:"bandwidth,omitempty"`
	GlobalBgp                   *GlobalBgp                   `json:"bgp,omitempty"`
	GlobalClusterComms          *GlobalClusterComms          `json:"cluster_comms,omitempty"`
	GlobalConnection            *GlobalConnection            `json:"connection,omitempty"`
	GlobalDataPlaneAcceleration *GlobalDataPlaneAcceleration `json:"data_plane_acceleration,omitempty"`
	GlobalDns                   *GlobalDns                   `json:"dns,omitempty"`
	GlobalDnsAutoscale          *GlobalDnsAutoscale          `json:"dns_autoscale,omitempty"`
	GlobalEc2                   *GlobalEc2                   `json:"ec2,omitempty"`
	GlobalEventing              *GlobalEventing              `json:"eventing,omitempty"`
	GlobalFaultTolerance        *GlobalFaultTolerance        `json:"fault_tolerance,omitempty"`
	GlobalFips                  *GlobalFips                  `json:"fips,omitempty"`
	GlobalFtp                   *GlobalFtp                   `json:"ftp,omitempty"`
	GlobalGlb                   *GlobalGlb                   `json:"glb,omitempty"`
	GlobalHistoricalActivity    *GlobalHistoricalActivity    `json:"historical_activity,omitempty"`
	GlobalHTTP                  *GlobalHTTP                  `json:"http,omitempty"`
	GlobalIP                    *GlobalIP                    `json:"ip,omitempty"`
	GlobalJava                  *GlobalJava                  `json:"java,omitempty"`
	GlobalKerberos              *GlobalKerberos              `json:"kerberos,omitempty"`
	GlobalLog                   *GlobalLog                   `json:"log,omitempty"`
	GlobalOspfv2                *GlobalOspfv2                `json:"ospfv2,omitempty"`
	GlobalPeriodicLog           *GlobalPeriodicLog           `json:"periodic_log,omitempty"`
	GlobalProtection            *GlobalProtection            `json:"protection,omitempty"`
	GlobalRecentConnections     *GlobalRecentConnections     `json:"recent_connections,omitempty"`
	GlobalRemoteLicensing       *GlobalRemoteLicensing       `json:"remote_licensing,omitempty"`
	GlobalRestAPI               *GlobalRestAPI               `json:"rest_api,omitempty"`
	GlobalSecurity              *GlobalSecurity              `json:"security,omitempty"`
	GlobalSession               *GlobalSession               `json:"session,omitempty"`
	GlobalSnmp                  *GlobalSnmp                  `json:"snmp,omitempty"`
	GlobalSoap                  *GlobalSoap                  `json:"soap,omitempty"`
	GlobalSourceNat             *GlobalSourceNat             `json:"source_nat,omitempty"`
	GlobalSsl                   *GlobalSsl                   `json:"ssl,omitempty"`
	GlobalSslHardware           *GlobalSslHardware           `json:"ssl_hardware,omitempty"`
	GlobalTrafficscript         *GlobalTrafficscript         `json:"trafficscript,omitempty"`
	GlobalWebCache              *GlobalWebCache              `json:"web_cache,omitempty"`
}

// global basic matches the API uri and JSON but it is actually global_system
// to match the user interface that users are familiar with.
type GlobalBasic struct {
	AcceptingDelay              int      `json:"accepting_delay,omitempty"`
	AfmEnabled                  bool     `json:"afm_enabled,omitempty"`
	ChildControlCommandTimeout  int      `json:"child_control_command_timeout,omitempty"`
	ChildControlKillTimeout     int      `json:"child_control_kill_timeout,omitempty"`
	ChunkSize                   int      `json:"chunk_size,omitempty"`
	ClientFirstOpt              bool     `json:"client_first_opt,omitempty"`
	ClusterIdentifier           string   `json:"cluster_identifier,omitempty"`
	CpuStarvationCheckInterval  int      `json:"cpu_starvation_check_interval,omitempty"`
	CpuStarvationCheckTolerance int      `json:"cpu_starvation_check_tolerance,omitempty"`
	DataPlaneAccelerationCores  string   `json:"data_plane_acceleration_cores,omitempty"`
	DataPlaneAccelerationMode   bool     `json:"data_plane_acceleration_mode,omitempty"`
	Http2noCipherBlacklistCheck bool     `json:"http2_no_cipher_blacklist_check,omitempty"`
	LicenseServers              []string `json:"license_servers,omitempty"`
	MaxFds                      int      `json:"max_fds,omitempty"`
	MonitorMemorySize           int      `json:"monitor_memory_size,omitempty"`
	RateClassLimit              int      `json:"rate_class_limit,omitempty"`
	SharedPoolSize              string   `json:"shared_pool_size,omitempty"`
	SlmClassLimit               int      `json:"slm_class_limit,omitempty"`
	SoRbuffSize                 int      `json:"so_rbuff_size,omitempty"`
	SoWbuffSize                 int      `json:"so_wbuff_size,omitempty"`
	SocketOptimizations         string   `json:"socket_optimizations,omitempty"`
	StorageShared               bool     `json:"storage_shared,omitempty"`
	TipClassLimit               int      `json:"tip_class_limit,omitempty"`
}

type GlobalAdmin struct {
	HonorFallbackScsv            bool     `json:"honor_fallback_scsv,omitempty"`
	Ssl3AllowRehandshake         string   `json:"ssl3_allow_rehandshake,omitempty"`
	Ssl3Ciphers                  string   `json:"ssl3_ciphers,omitempty"`
	Ssl3DiffieHellmanKeyLength   string   `json:"ssl3_diffie_hellman_key_length,omitempty"`
	Ssl3MinRehandshakeInterval   int      `json:"ssl3_min_rehandshake_interval,omitempty"`
	SslEllipticCurves            []string `json:"ssl_elliptic_curves,omitempty"`
	SslInsertExtraFragment       bool     `json:"ssl_insert_extra_fragment,omitempty"`
	SslMaxHandshakeMessageSize   int      `json:"ssl_max_handshake_message_size,omitempty"`
	SslPreventTimingSideChannels bool     `json:"ssl_prevent_timing_side_channels,omitempty"`
	SslSignatureAlgorithms       string   `json:"ssl_signature_algorithms,omitempty"`
	SupportSsl2                  bool     `json:"support_ssl2,omitempty"`
	SupportSsl3                  bool     `json:"support_ssl3,omitempty"`
	SupportTls1                  bool     `json:"support_tls1,omitempty"`
	SupportTls11                 bool     `json:"support_tls11,omitempty"`
	SupportTls12                 bool     `json:"support_tls12,omitempty"`
}

type GlobalAppliance struct {
	BootloaderPassword       string `json:"bootloader_password,omitempty"`
	ManageNcipher            bool   `json:"manage_ncipher,ommitempty"`
	NethsmEsn                string `json:"nethsm_esn,omitempty"`
	NethsmHash               string `json:"nethsm_hash,omitempty"`
	NethsmIp                 string `json:"nethsm_ip,omitempty"`
	NethsmNcipherRfs         string `json:"nethsm_ncipher_rfs,omitempty"`
	ReturnPathRoutingEnabled bool   `json:"return_path_routing_enabled,omitempty"`
}

type GlobalAptimizer struct {
	MaxDependentFetchSize        string `json:"max_dependent_fetch_size,omitempty"`
	MaxOriginalContentBufferSize string `json:"max_original_content_buffer_size,omitempty"`
	WatchdogInterval             int    `json:"watchdog_interval,omitempty"`
	WatchdogLimit                int    `json:"watchdog_limit,omitempty"`
}

type GlobalAuditlog struct {
	ViaEventd bool `json:"via_eventd,omitempty"`
	ViaSyslog bool `json:"via_syslog,omitempty"`
}

type GlobalBandwidth struct {
}

type GlobalAutoScaler struct {
	Verbose bool `json:"verbose,omitempty"`
}

type GlobalBgp struct {
	AsNumber int  `json:"as_number,omitempty"`
	Enabled  bool `json:"enabled,omitempty"`
}

type GlobalClusterComms struct {
	AllowUpdateDefault bool     `json:"allow_update_default,omitempty"`
	AllowedUpdateHosts []string `json:"allowed_update_hosts,omitempty"`
	StateSyncInterval  int      `json:"state_sync_interval,omitempty"`
	StateSyncTimeout   int      `json:"state_sync_timeout,omitempty"`
}

type GlobalConnection struct {
	IdleConnectionsMax int  `json:"idle_connections_max,omitempty"`
	IdleTimeout        int  `json:"idle_timeout,omitempty"`
	ListenQueueSize    int  `json:"listen_queue_size,omitempty"`
	MaxAccepting       int  `json:"max_accepting,omitempty"`
	MultipleAccept     bool `json:"multiple_accept,omitempty"`
}

type GlobalDataPlaneAcceleration struct {
	TcpDelayAck int `json:"tcp_delay_ack,omitempty"`
	TcpWinScale int `json:"tcp_win_scale,omitempty"`
}

type GlobalDns struct {
	MaxTTL         int `json:"max_ttl,omitempty"`
	MinTTL         int `json:"min_ttl,omitempty"`
	NegativeExpiry int `json:"negative_expiry,omitempty"`
	Size           int `json:"size,omitempty"`
	Timeout        int `json:"timeout,omitempty"`
}

type GlobalDnsAutoscale struct {
}

type GlobalEc2 struct {
	AccessKeyID           string `json:"access_key_id,omitempty"`
	AwstoolTimeout        int    `json:"awstool_timeout,omitempty"`
	SecretAccessKey       string `json:"secret_access_key,omitempty"`
	VerifyQueryServerCert bool   `json:"verify_query_server_cert,omitempty"`
}

type GlobalEventing struct {
	MailInterval int `json:"mail_interval,omitempty"`
	MaxAttempts  int `json:"max_attempts,omitempty"`
}

type GlobalFaultTolerance struct {
	ArpCount            int      `json:"arp_count,omitempty"`
	AutoFailback        bool     `json:"auto_failback,omitempty"`
	ChildTimeout        int      `json:"child_timeout,omitempty"`
	FrontendCheckIps    []string `json:"frontend_check_ips,omitempty"`
	HeartbeatMethod     string   `json:"heartbeat_method,omitempty"`
	IgmpInterval        int      `json:"igmp_interval,omitempty"`
	L4AccelChildTimeout int      `json:"l4accel_child_timeout,omitempty"`
	L4AccelSyncPort     int      `json:"l4accel_sync_port,omitempty"`
	MonitorInterval     int      `json:"monitor_interval,omitempty"`
	MonitorTimeout      int      `json:"monitor_timeout,omitempty"`
	MulticastAddress    string   `json:"multicast_address,omitempty"`
	UnicastPort         int      `json:"unicast_port,omitempty"`
	UseBindIP           bool     `json:"use_bind_ip,omitempty"`
	Verbose             bool     `json:"verbose,omitempty"`
}

type GlobalFips struct {
	Enabled bool `json:"enabled,omitempty"`
}

type GlobalFtp struct {
	DataBindLow bool `json:"data_bind_low,omitempty"`
}

type GlobalGlb struct {
	Verbose bool `json:"verbose,omitempty"`
}

type GlobalHistoricalActivity struct {
	KeepDays int `json:"keep_days,omitempty"`
}

type GlobalHTTP struct {
}

type GlobalIP struct {
	ApplianceReturnpath []interface{} `json:"appliance_returnpath,omitempty"`
}

type GlobalJava struct {
	Classpath      string `json:"classpath,omitempty"`
	Command        string `json:"command,omitempty"`
	Enabled        bool   `json:"enabled,omitempty"`
	Lib            string `json:"lib,omitempty"`
	MaxConnections int    `json:"max_connections,omitempty"`
	SessionAge     int    `json:"session_age,omitempty"`
}

type GlobalKerberos struct {
	Verbose bool `json:"verbose,omitempty"`
}

type GlobalLog struct {
	ErrorLevel string `json:"error_level,omitempty"`
	FlushTime  int    `json:"flush_time,omitempty"`
	LogFile    string `json:"log_file,omitempty"`
	Rate       int    `json:"rate,omitempty"`
	Reopen     int    `json:"reopen,omitempty"`
	Time       int    `json:"time,omitempty"`
}

type GlobalOspfv2 struct {
	Area                        string `json:"area,omitempty"`
	AreaType                    string `json:"area_type,omitempty"`
	AuthenticationKeyIDA        int    `json:"authentication_key_id_a,omitempty"`
	AuthenticationKeyIDB        int    `json:"authentication_key_id_b,omitempty"`
	AuthenticationSharedSecretA string `json:"authentication_shared_secret_a,omitempty"`
	AuthenticationSharedSecretB string `json:"authentication_shared_secret_b,omitempty"`
	Enabled                     bool   `json:"enabled,omitempty"`
	HelloInterval               int    `json:"hello_interval,omitempty"`
	RouterDeadInterval          int    `json:"router_dead_interval,omitempty"`
}

type GlobalPeriodicLog struct {
}

type GlobalProtection struct {
	ConncountSize string `json:"conncount_size,omitempty"`
}

type GlobalRecentConnections struct {
	MaxPerProcess int `json:"max_per_process,omitempty"`
	RetainTime    int `json:"retain_time,omitempty"`
	SnapshotSize  int `json:"snapshot_size,omitempty"`
}

type GlobalRemoteLicensing struct {
	RegistrationServer string `json:"registration_server,omitempty"`
	ServerCertificate  string `json:"server_certificate,omitempty"`
}

type GlobalRestAPI struct {
	AuthTimeout         int  `json:"auth_timeout,omitempty"`
	Enabled             bool `json:"enabled,omitempty"`
	HTTPMaxHeaderLength int  `json:"http_max_header_length,omitempty"`
	ReplicateAbsolute   int  `json:"replicate_absolute,omitempty"`
	ReplicateLull       int  `json:"replicate_lull,omitempty"`
	ReplicateTimeout    int  `json:"replicate_timeout,omitempty"`
}

type GlobalSecurity struct {
	LoginBanner                   string `json:"login_banner,omitempty"`
	LoginBannerAccept             bool   `json:"login_banner_accept,omitempty"`
	LoginDelay                    int    `json:"login_delay,omitempty"`
	MaxLoginAttempts              int    `json:"max_login_attempts,omitempty"`
	MaxLoginExternal              bool   `json:"max_login_external,omitempty"`
	MaxLoginSuspensionTime        int    `json:"max_login_suspension_time,omitempty"`
	PasswordAllowConsecutiveChars bool   `json:"password_allow_consecutive_chars,omitempty"`
	PasswordChangesPerDay         int    `json:"password_changes_per_day,omitempty"`
	PasswordMinAlphaChars         int    `json:"password_min_alpha_chars,omitempty"`
	PasswordMinLength             int    `json:"password_min_length,omitempty"`
	PasswordMinNumericChars       int    `json:"password_min_numeric_chars,omitempty"`
	PasswordMinSpecialChars       int    `json:"password_min_special_chars,omitempty"`
	PasswordMinUppercaseChars     int    `json:"password_min_uppercase_chars,omitempty"`
	PasswordReuseAfter            int    `json:"password_reuse_after,omitempty"`
	PostLoginBanner               string `json:"post_login_banner,omitempty"`
	TrackUnknownUsers             bool   `json:"track_unknown_users,omitempty"`
	UIPageBanner                  string `json:"ui_page_banner,omitempty"`
}

type GlobalSession struct {
	AspCacheSize       int `json:"asp_cache_size,omitempty"`
	IPCacheSize        int `json:"ip_cache_size,omitempty"`
	J2EeCacheSize      int `json:"j2ee_cache_size,omitempty"`
	SslCacheSize       int `json:"ssl_cache_size,omitempty"`
	UniversalCacheSize int `json:"universal_cache_size,omitempty"`
}

type GlobalSnmp struct {
	UserCounters int `json:"user_counters,omitempty"`
}

type GlobalSoap struct {
	IdleMinutes int `json:"idle_minutes,omitempty"`
}

type GlobalSourceNat struct {
	IPLimit              int `json:"ip_limit,omitempty"`
	IPLocalPortRangeHigh int `json:"ip_local_port_range_high,omitempty"`
	SharedPoolSize       int `json:"shared_pool_size,omitempty"`
}

type GlobalSsl struct {
	CacheExpiry                        int           `json:"cache_expiry,omitempty"`
	CachePerVirtualserver              bool          `json:"cache_per_virtualserver,omitempty"`
	CacheSize                          int           `json:"cache_size,omitempty"`
	CrlMemSize                         string        `json:"crl_mem_size,omitempty"`
	EllipticCurves                     []interface{} `json:"elliptic_curves,omitempty"`
	HonorFallbackScsv                  bool          `json:"honor_fallback_scsv,omitempty"`
	InsertExtraFragment                bool          `json:"insert_extra_fragment,omitempty"`
	MaxHandshakeMessageSize            int           `json:"max_handshake_message_size,omitempty"`
	OcspCacheSize                      int           `json:"ocsp_cache_size,omitempty"`
	OcspStaplingDefaultRefreshInterval int           `json:"ocsp_stapling_default_refresh_interval,omitempty"`
	OcspStaplingMaximumRefreshInterval int           `json:"ocsp_stapling_maximum_refresh_interval,omitempty"`
	OcspStaplingMemSize                string        `json:"ocsp_stapling_mem_size,omitempty"`
	OcspStaplingTimeTolerance          int           `json:"ocsp_stapling_time_tolerance,omitempty"`
	OcspStaplingVerifyResponse         bool          `json:"ocsp_stapling_verify_response,omitempty"`
	PreventTimingSideChannels          bool          `json:"prevent_timing_side_channels,omitempty"`
	SignatureAlgorithms                string        `json:"signature_algorithms,omitempty"`
	Ssl3AllowRehandshake               string        `json:"ssl3_allow_rehandshake,omitempty"`
	Ssl3Ciphers                        string        `json:"ssl3_ciphers,omitempty"`
	Ssl3DiffieHellmanKeyLength         string        `json:"ssl3_diffie_hellman_key_length,omitempty"`
	Ssl3MinRehandshakeInterval         int           `json:"ssl3_min_rehandshake_interval,omitempty"`
	SupportSsl2                        bool          `json:"support_ssl2,omitempty"`
	SupportSsl3                        bool          `json:"support_ssl3,omitempty"`
	SupportTLS1                        bool          `json:"support_tls1,omitempty"`
	SupportTLS11                       bool          `json:"support_tls1_1,omitempty"`
	SupportTLS12                       bool          `json:"support_tls1_2,omitempty"`
}

type GlobalSslHardware struct {
	Accel                  bool   `json:"accel,omitempty"`
	AzureClientID          string `json:"azure_client_id,omitempty"`
	AzureClientSecret      string `json:"azure_client_secret,omitempty"`
	AzureVaultURL          string `json:"azure_vault_url,omitempty"`
	AzureVerifyRestAPICert bool   `json:"azure_verify_rest_api_cert,omitempty"`
	DriverPkcs11Debug      bool   `json:"driver_pkcs11_debug,omitempty"`
	DriverPkcs11Lib        string `json:"driver_pkcs11_lib,omitempty"`
	DriverPkcs11SlotDesc   string `json:"driver_pkcs11_slot_desc,omitempty"`
	DriverPkcs11SlotType   string `json:"driver_pkcs11_slot_type,omitempty"`
	DriverPkcs11UserPin    string `json:"driver_pkcs11_user_pin,omitempty"`
	FailureCount           int    `json:"failure_count,omitempty"`
	Library                string `json:"library,omitempty"`
}

type GlobalTrafficscript struct {
	ArrayElements            int    `json:"array_elements,omitempty"`
	DataLocalSize            string `json:"data_local_size,omitempty"`
	DataSize                 string `json:"data_size,omitempty"`
	ExecutionTimeWarning     int    `json:"execution_time_warning,omitempty"`
	MaxInstr                 int    `json:"max_instr,omitempty"`
	MemoryWarning            int    `json:"memory_warning,omitempty"`
	RegexCacheSize           int    `json:"regex_cache_size,omitempty"`
	RegexMatchLimit          int    `json:"regex_match_limit,omitempty"`
	RegexMatchWarnPercentage int    `json:"regex_match_warn_percentage,omitempty"`
	VariablePoolUse          bool   `json:"variable_pool_use,omitempty"`
}

type GlobalWebCache struct {
	AvgPathLength  int    `json:"avg_path_length,omitempty"`
	Disk           bool   `json:"disk,omitempty"`
	DiskDir        string `json:"disk_dir,omitempty"`
	MaxFileNum     int    `json:"max_file_num,omitempty"`
	MaxFileSize    string `json:"max_file_size,omitempty"`
	MaxPathLength  int    `json:"max_path_length,omitempty"`
	NormalizeQuery bool   `json:"normalize_query,omitempty"`
	Size           string `json:"size,omitempty"`
	Verbose        bool   `json:"verbose,omitempty"`
}
