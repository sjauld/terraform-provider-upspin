package upspin

import (
	"github.com/hashicorp/terraform/helper/schema"

	"upspin.io/rpc"
	"upspin.io/upspin"
)

func dataSourceUpspinUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUpspinRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceUpspinRead(d *schema.ResourceData, meta interface{}) error {
	conn := rpc.PublicUserKeyService(meta.(upspin.Config))
	username := d.Get("username").(string)

	pubkey, err := conn(upspin.UserName(username))
	if err != nil {
		return err
	}

	// Set the attributes
	d.SetId(username)
	d.Set("username", username)
	d.Set("public_key", pubkey)

	return nil
}
