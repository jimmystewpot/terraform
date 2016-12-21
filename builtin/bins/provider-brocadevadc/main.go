package main

import (
	"github.com/hashicorp/terraform/builtin/providers/brocadevadc"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: brocadevadc.Provider,
	})
}
