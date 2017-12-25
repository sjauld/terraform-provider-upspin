---
layout: "upspin"
page_title: "Provider: Upspin"
sidebar_current: "docs-upspin-index"
description: |-
  The Upspin provider is used to interact with the resources supported by Upspin.
---

# Upspin Provider

The Upspin provider is used to interact with the resources supported by Upspin.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Upspin Provider
provider "upspin" {
  version = "~> 0.0.1"
}

# Get a user
data "upspin_user" "rob" {
  username = "r@golang.org"
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `key_server` - (Optional) The key server you'd like to use (key.upspin.io is the default)
* `key_server_port` - (Optional) A custom port for your key server (default 443)
* `transport` - (Optional) Supports inprocess, remote (default) or unassigned
