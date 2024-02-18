package main

import (
	"fmt"
	"os"

	"github.com/orange-cloudavenue/cloudavenue-cli/cmd"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/customErrors"
)

const (
	errorMessage = "** Error **: "
)

//go:generate go run tools/tools.go

func main() {
	if err := cmd.Execute(); err != nil {
		switch {
		case customErrors.IsNoHomeDirectory(err):
			fmt.Println(errorMessage+"Check your system to set and access write to your home directory.", err)
		case customErrors.IsConfigFile(err):
			fmt.Println(errorMessage+"Please check your configuration file.", err)
		case customErrors.IsClient(err):
			fmt.Println(errorMessage+"Unable to initialize client", err)
			fmt.Println("Please check your configuration (https://orange-cloudavenue.github.io/cloudavenue-cli/getting-started/configuration/).")
		case customErrors.IsNotValidOutput(err):
			fmt.Println(errorMessage+"Please read help to check format output is possible.", err)
		default:
			fmt.Println(errorMessage + err.Error())
		}
		os.Exit(1)
	}
	os.Exit(0)
}
