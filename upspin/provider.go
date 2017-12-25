package upspin

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"upspin.io/config"

	// This import collects all the transport implementations in the core Upspin
	// libraries and adds them to your program.
	_ "upspin.io/key/transports"
)

// Provider returns a terraform.ResourceProvider
func Provider() terraform.ResourceProvider {
	cfg := config.New()

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"cache": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      cfg.Value("cache"),
				Description:  "Whether to use a local store and directory cache server.",
				ValidateFunc: validateEndpoint,
			},
			"dirserver": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      cfg.DirEndpoint().String(),
				Description:  "Server that holds user's directory tree.",
				ValidateFunc: validateEndpoint,
			},
			"keyserver": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      cfg.KeyEndpoint().String(),
				Description:  "Which key server to use.",
				ValidateFunc: validateEndpoint,
			},
			"packing": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      cfg.Packing().String(),
				Description:  "Algorithm to encrypt or protect data.",
				ValidateFunc: validatePacking,
			},
			"secrets": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     string(cfg.Value("secrets")),
				Description: "Directory holding private keys.",
			},
			"storeserver": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      cfg.StoreEndpoint().String(),
				Description:  "Server to write new storage.",
				ValidateFunc: validateEndpoint,
			},
			"tlscerts": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     cfg.Value("tlscerts"),
				Description: "Directory holding TLS certificates.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultUserName(),
				Description: "E-mail address of user.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"upspin_user": dataSourceUpspinUser(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func defaultUserName() string {
	cfg := config.New()
	return string(cfg.UserName())
}

// returns an upspin.Config that can be used by the resources to make API calls
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var buffer bytes.Buffer
	username := d.Get("username").(string)
	buffer.WriteString(fmt.Sprintf("username: %s\n", username))
	buffer.WriteString(fmt.Sprintf("packing: %s\n", d.Get("packing").(string)))
	buffer.WriteString(fmt.Sprintf("keyserver: %s\n", d.Get("keyserver").(string)))
	buffer.WriteString(fmt.Sprintf("dirserver: %s\n", d.Get("dirserver").(string)))
	buffer.WriteString(fmt.Sprintf("storeserver: %s\n", d.Get("storeserver").(string)))

	cache := d.Get("cache").(string)
	if cache != "" {
		buffer.WriteString(fmt.Sprintf("cache: %s\n", cache))
	}

	// If no username is specified, we're in read only mode and should specify
	// "secrets: none" to avoid attempts to find our local keys
	var secrets string
	if username == defaultUserName() {
		secrets = "none"
	} else {
		secrets = d.Get("secrets").(string)
	}
	if secrets != "" {
		buffer.WriteString(fmt.Sprintf("secrets: %s\n", secrets))
	}

	tlscerts := d.Get("tlscerts").(string)
	if tlscerts != "" {
		buffer.WriteString(fmt.Sprintf("tlscerts: %s\n", tlscerts))
	}

	cfg, err := config.InitConfig(strings.NewReader(buffer.String()))
	// Remove the ErrNoFactotum if it is expected
	if err == config.ErrNoFactotum && secrets == "none" {
		err = nil
	}

	return cfg, err
}
