/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/orange-cloudavenue/cloudavenue-sdk-go"
	clientcloudavenue "github.com/orange-cloudavenue/cloudavenue-sdk-go/pkg/clients/cloudavenue"
	"github.com/spf13/cobra"
)

var (
	c       *cloudavenue.Client
	version = "v0.0.4"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "cloudavenue-cli",
	Short:   "cloudavenue-cli is the Command Line Interface for CloudAvenue Platform",
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {

	// Set client CloudAvenue
	var err error
	c, err = cloudavenue.New(cloudavenue.ClientOpts{
		CloudAvenue: &clientcloudavenue.Opts{},
	})
	if err != nil {
		fmt.Println("Error in CloudAvenue parameter", err)
		return err
	}

	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error in Command", err)
		return err
	}
	return nil
}

// NewRootCmd returns the root command
func NewRootCmd() *cobra.Command {
	// rootCmd.Args = cobra.MinimumNArgs(1)
	// viper.AutomaticEnv() // read in environment variables that match
	return rootCmd
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("=== Command %s took %s second(s) ===", name, elapsed)
}
