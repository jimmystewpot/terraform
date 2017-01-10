package brocadevadc

import (
	"sync"
)

// Config is the configuration structure used to instantiate the brocade virtual application delivery controller
type ClientConfig struct {
	sync.Mutex
	URL       string
	Username  string
	Password  string
	SslVerify bool
}
