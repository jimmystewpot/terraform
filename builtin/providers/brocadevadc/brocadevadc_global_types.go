package brocadevadc

// Globals JSON encoded composition.
// /api/tm/3.9/config/active/global_settings
type Globals struct {
	Properties *Properties `json:"properties"`
}

type Properties struct {
	GlobalBasic     *GlobalBasic     `json:"basic,omitempty"`
	GlobalAdmin     *GlobalAdmin     `json:"admin,omitempty"`
	GlobalAppliance *GlobalAppliance `json:"appliance,omitempty"`
	//Globalaptimizer `json:"aptimizer,omitempty"`
	//Globalauditlog  `json:"auditlog,omitempty"`
	//Globalbandwidth `json:"bandwidth,omitempty"`
}

// global basic matches the API uri and JSON but it is actually global_system
// to match the user interface that users are familiar with.
type GlobalBasic struct {
	Uuid                        string   `json:"uuid,omitempty"`
	AcceptingDelay              uint     `json:"accepting_delay,omitempty"`
	AfmEnabled                  bool     `json:"afm_enabled,omitempty"`
	ChildControlCommandTimeout  uint     `json:"child_control_command_timeout,omitempty"`
	ChildControlKillTimeout     uint     `json:"child_control_kill_timeout,omitempty"`
	ChunkSize                   uint     `json:"chunk_size,omitempty"`
	ClientFirstOpt              bool     `json:"client_first_opt,omitempty"`
	ClusterIdentifier           string   `json:"cluster_identifier,omitempty"`
	CpuStarvationCheckInterval  uint     `json:"cpu_starvation_check_interval,omitempty"`
	CpuStarvationCheckTolerance uint     `json:"cpu_starvation_check_tolerance,omitempty"`
	DataPlaneAccelerationCores  string   `json:"data_plane_acceleration_cores,omitempty"`
	DataPlaneAccelerationMode   bool     `json:"data_plane_acceleration_mode,omitempty"`
	Http2noCipherBlacklistCheck bool     `json:"http2_no_cipher_blacklist_check,omitempty"`
	LicenseServers              []string `json:"license_servers,omitempty"`
	MaxFds                      uint     `json:"max_fds,omitempty"`
	MonitorMemorySize           uint     `json:"monitor_memory_size,omitempty"`
	RateClassLimit              uint     `json:"rate_class_limit,omitempty"`
	SharedPoolSize              string   `json:"shared_pool_size,omitempty"`
	SlmClassLimit               uint     `json:"slm_class_limit,omitempty"`
	SoRbuffSize                 uint     `json:"so_rbuff_size,omitempty"`
	SoWbuffSize                 uint     `json:"so_wbuff_size,omitempty"`
	SocketOptimizations         string   `json:"socket_optimizations,omitempty"`
	StorageShared               bool     `json:"storage_shared,omitempty"`
	TipClassLimit               uint     `json:"tip_class_limit,omitempty"`
}

type GlobalAdmin struct {
	Uuid                         string   `json:"uuid,omitempty"`
	HonorFallbackScsv            bool     `json:"honor_fallback_scsv,omitempty"`
	Ssl3AllowRehandshake         string   `json:"ssl3_allow_rehandshake,omitempty"`
	Ssl3Ciphers                  string   `json:"ssl3_ciphers,omitempty"`
	Ssl3DiffieHellmanKeyLength   uint     `json:"ssl3_diffie_hellman_key_length,omitempty"`
	Ssl3MinRehandshakeInterval   uint     `json:"ssl3_min_rehandshake_interval,omitempty"`
	SslEllipticCurves            []string `json:"ssl_elliptic_curves,omitempty"`
	SslInsertExtraFragment       bool     `json:"ssl_insert_extra_fragment,omitempty"`
	SslMaxHandshakeMessageSize   uint     `json:"ssl_max_handshake_message_size,omitempty"`
	SslPreventTimingSideChannels bool     `json:"ssl_prevent_timing_side_channels,omitempty"`
	SslSignatureAlgorithms       string   `json:"ssl_signature_algorithms,omitempty"`
	SupportSsl2                  bool     `json:"support_ssl2,omitempty"`
	SupportSsl3                  bool     `json:"support_ssl3,omitempty"`
	SupportTls1                  bool     `json:"support_tls1,omitempty"`
	SupportTls11                 bool     `json:"support_tls11,omitempty"`
	SupportTls12                 bool     `json:"support_tls12,omitempty"`
}

type GlobalAppliance struct {
	Bootloader_password         string `json:"bootloader_password,omitempty"`
	Manage_ncipher              bool   `json:"manage_ncipher,ommitempty"`
	Nethsm_esn                  string `json:"nethsm_esn,omitempty"`
	Nethsm_hash                 string `json:"nethsm_hash,omitempty"`
	Nethsm_ip                   string `json:"nethsm_ip,omitempty"`
	Nethsm_ncipher_rfs          string `json:"nethsm_ncipher_rfs,omitempty"`
	Return_path_routing_enabled bool   `json:"return_path_routing_enabled,omitempty"`
}

type Globalaptimizer struct {
}

type Globalauditlog struct {
}

type Globalbandwidth struct {
}
