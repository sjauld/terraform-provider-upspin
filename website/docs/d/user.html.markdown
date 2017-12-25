---
layout: "upspin"
page_title: "Upspin: upspin_user"
sidebar_current: "docs-upspin-datasource-user"
description: |-
  Retrieve a user's details from the keyserver
---

# upspin_user

User this data source to retrieve information about a user from the key server

## Example Usage

```hcl
data "upspin_user" "main" {
  username = "r@golang.org"
}
```

## Argument Reference

 * `username` - (Required) The email address of the user, as registered in the key server

## Attributes Reference

 * `username` - The email address of the user, as registered in the key server
 * `public_key` - THe user's public key
