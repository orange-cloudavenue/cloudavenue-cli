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
	Long: `cloudavenue-cli give you a different way to manage your IaaS without terraform command.
	Is able to List, Show, Create and Delete resources.
	For example: 
	- cloudavenue-cli list edgegateway => will list all your VMs.
	- cloudavenue-cli show edgegateway <id> => will show you the VM with the id <id>.
	- cloudavenue-cli create edgegateway => will create a VM.
	- cloudavenue-cli delete edgegateway <id> => will delete the VM with the id <id>.`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cloudavenue-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
