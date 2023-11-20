/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/orange-cloudavenue/cloudavenue-sdk-go"
	clientcloudavenue "github.com/orange-cloudavenue/cloudavenue-sdk-go/pkg/clients/cloudavenue"
	"github.com/spf13/cobra"
)

var (
	c *cloudavenue.Client
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cloudavenue-cli",
	Short: "cloudavenue-cli is the Command Line Interface for CloudAvenue Platform",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	var err error
	c, err = cloudavenue.New(cloudavenue.ClientOpts{
		CloudAvenue: clientcloudavenue.Opts{
			Debug:    false,
			Username: "gaetan.ars",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Args = cobra.MinimumNArgs(1)
}
