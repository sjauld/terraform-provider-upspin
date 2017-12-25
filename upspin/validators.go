package upspin

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var validTransports = map[string]bool{
	"inprocess":  true,
	"remote":     true,
	"unassigned": true,
}

func validateEndpoint(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	if !regexp.MustCompile(`^[a-z]*,?[a-z.]*:?\d*$`).MatchString(value) {
		errors = append(errors, fmt.Errorf("%q does not contain a valid endpoint", k))
	}

	ep := strings.Split(value, ",")
	var transport, uri string
	switch len(ep) {
	case 2:
		transport, uri = ep[0], ep[1]
	case 1:
		uri = ep[0]
	}

	if transport != "" {
		_, terrors := validateTransport(transport, k)
		errors = append(errors, terrors...)
	}

	if transport == "remote" || transport == "" {
		_, uerrors := validateURI(uri, k)
		errors = append(errors, uerrors...)
	}

	return
}

func validatePacking(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(`^(ee|eeintegrity|plain)$`).MatchString(value) {
		errors = append(errors, fmt.Errorf("%q does not contain a valid packing", k))
	}

	return
}

func validateTransport(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !validTransports[value] {
		errors = append(errors, fmt.Errorf("%v does not contain a supported transport", k))
	}
	return
}

func validateURI(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(`^[a-z.]+:?\d*?$`).MatchString(value) {
		errors = append(errors, fmt.Errorf("%q does not contain a valid uri", k))
	}

	uri := strings.Split(value, ":")

	if len(uri) == 2 {
		// check that the port provided is a 16 bit unsigned integer
		_, err := strconv.ParseUint(uri[1], 10, 16)
		if err != nil {
			errors = append(errors, fmt.Errorf("%q contains an invalid port", k))
		}
	}

	return
}
