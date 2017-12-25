package upspin

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testUserName = "r@golang.org"

func TestAccDataSourceUpspinUser(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceUpspinUserConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceUpspinUserCheck("data.upspin_user.main"),
					resource.TestCheckResourceAttr("data.upspin_user.main", "username", testUserName),
					resource.TestCheckResourceAttrSet("data.upspin_user.main", "public_key"),
				),
			},
		},
	})
}

func testAccDataSourceUpspinUserCheck(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("User ID not set")
		}
		return nil
	}
}

var testAccDataSourceUpspinUserConfig = fmt.Sprintf(`
data upspin_user "main" {
	username = "%s"
}
`, testUserName)
