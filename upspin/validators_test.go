package upspin

import (
	"testing"
)

func TestValidateEndpoint(t *testing.T) {
	validEndpoints := []string{
		"key.upspin.io",
		"key.upspin.io:443",
		"custom.endpoint",
		"custom.endpoint:8080",
		"custom.endpoint:0", // not sure if port zero should be made invalid?
		"remote,key.upspin.io",
		"remote,key.upspin.io:443",
		"remote,custom.endpoint",
		"remote,custom.endpoint:8080",
		"inprocess,",
		"inprocess,uri.ignored",
		"unassigned,",
		"unassigned,uri.ignored",
	}
	for _, v := range validEndpoints {
		_, errors := validateEndpoint(v, "endpoint")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid endpoint: %q", v, errors)
		}
	}

	invalidEndpoints := []string{
		"wrongtransport,",
		"wrongtransport,custom.endpoint",
		"negativeport:-1",
		"bigport:65536",
		"stringport:eleven",
		"bad[uri]",
	}
	for _, v := range invalidEndpoints {
		_, errors := validateEndpoint(v, "endpoint")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid endpoint: %q", v, errors)
		}
	}

	return
}

func TestValidatePacking(t *testing.T) {
	validPackings := []string{
		"ee",
		"eeintegrity",
		"plain",
	}
	for _, v := range validPackings {
		_, errors := validatePacking(v, "packing")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid packing: %q", v, errors)
		}
	}

	invalidPackings := []string{
		"bob",
		"FAT32",
		"\\d\\d\\\\d\\\\//\\/",
	}
	for _, v := range invalidPackings {
		_, errors := validatePacking(v, "packing")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid packing: %q", v, errors)
		}
	}
}

func TestValidateTransport(t *testing.T) {
	validTransports := []string{
		"inprocess",
		"remote",
		"unassigned",
	}
	for _, v := range validTransports {
		_, errors := validateTransport(v, "transport")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid transport name: %q", v, errors)
		}
	}

	invalidTransports := []string{
		"notinprocess",
		"rebote",
		"b[if]tek",
		"key.upspin.io",
	}
	for _, v := range invalidTransports {
		_, errors := validateTransport(v, "transport")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid transport name: %q", v, errors)
		}
	}
}
