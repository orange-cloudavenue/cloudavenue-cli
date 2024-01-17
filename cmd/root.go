/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/adampresley/sigint"
	"github.com/briandowns/spinner"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go"
	clientcloudavenue "github.com/orange-cloudavenue/cloudavenue-sdk-go/pkg/clients/cloudavenue"
	"github.com/spf13/cobra"
)

var (
	c       *cloudavenue.Client
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
	s       = spinner.New(spinner.CharSets[43], 100*time.Millisecond)
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cav",
	Short: "cav is the Command Line Interface for CloudAvenue Platform",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	// ctrl+c handler
	sigint.ListenForSIGINT(func() {
		fmt.Println("SIGINT received. Exiting...")
		os.Exit(0)
	})

	// Set client CloudAvenue
	var err error
	c, err = cloudavenue.New(cloudavenue.ClientOpts{
		CloudAvenue: &clientcloudavenue.Opts{},
	})
	if err != nil {
		s.Stop()
		fmt.Println("Error in CloudAvenue parameter, please check your configuration (https://github.com/orange-cloudavenue/cloudavenue-cli/blob/main/docs/index.md)", err)
		return err
	}

	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		s.Stop()
		fmt.Println("Error in Command", err)
		return err
	}
	return nil
}

// NewRootCmd returns the root command
func NewRootCmd() *cobra.Command {
	return rootCmd
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("\n=== Command %s took %s second(s) ===", name, elapsed)
}

func init() {
	s.Start()
	rootCmd.PersistentFlags().BoolP("time", "t", false, "time elapsed for command")
	rootCmd.AddCommand(versionCmd())
}

// func versionCmd() return the version of the CLI
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of cav",
		Long:  `All software has versions. This is cav's`,
		Run: func(cmd *cobra.Command, args []string) {
			s.Stop()
			fmt.Printf("Version: %s\nCommit: %s\nBuilt at: %s\nBuilt by: %s\n", version, commit, date, builtBy)
		},
	}
}
