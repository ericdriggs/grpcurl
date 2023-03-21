package main

import (
	"fmt"
	"testing"
)

func TestGrpcurl(t *testing.T) {

	osArgs := []string{
		"-d",
		`{"configName": "assertion", "dimensionQueries":[{"key":"countryOrRegion", "sVal":"BR"}]}`,
		"--insecure",
		"config-service-qa.us-east-1.dpegrid.net:443",
		"ConfigServer.ConfigServer/GetConfig",
	}
	setOsArgs(osArgs)

	ecm := doMain()
	print(ecm.toJSON())
	fmt.Printf("ecm.exitCode: %d ", ecm.ExitCode)
	fmt.Printf("ecm.message: %s ", ecm.Message)

}
