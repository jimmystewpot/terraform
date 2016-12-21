package brocadevadc

import ()

// Config is the configuration structure used to instantiate the brocade virtual application delivery controller
type ClientConfig struct {
	URL       string
	Username  string
	Password  string
	VerifySSL bool
}
