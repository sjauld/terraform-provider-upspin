package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sjauld/terraform-provider-upspin/upspin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: upspin.Provider})
}
