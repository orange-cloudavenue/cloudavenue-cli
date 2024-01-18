/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/adampresley/sigint"
	"github.com/briandowns/spinner"
	cloudavenue "github.com/orange-cloudavenue/cloudavenue-sdk-go"
	clientcloudavenue "github.com/orange-cloudavenue/cloudavenue-sdk-go/pkg/clients/cloudavenue"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	c                   *cloudavenue.Client
	version             = "dev"
	commit              = "none"
	date                = "unknown"
	builtBy             = "unknown"
	s                   = spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	cloudavenueOrg      string
	cloudavenueUsername string
	cloudavenuePassword string
	cloudavenueDebug    bool
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

	// Set default file configuration and create it if not exist
	home := os.Getenv("HOME")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.cav")
	if home == "" {
		return fmt.Errorf("Error in Get HOME Directory")
	}
	if _, err := os.Stat(home + "/.cav/config.yaml"); os.IsNotExist(err) {
		if err := os.MkdirAll(home+"/.cav", 0755); err != nil {
			return err
		}
		viper.SetDefault("cloudavenue_username", "")
		viper.SetDefault("cloudavenue_password", "")
		viper.SetDefault("cloudavenue_org", "")
		viper.AutomaticEnv()
		viper.SetDefault("cloudavenue_debug", false)

		if err := viper.SafeWriteConfig(); err != nil {
			return err
		}
	}

	// check if variable is set if not, use configuration file
	if os.Getenv("CLOUDAVENUE_USERNAME") == "" || os.Getenv("CLOUDAVENUE_PASSWORD") == "" || os.Getenv("CLOUDAVENUE_ORG") == "" {
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		cloudavenueUsername = viper.GetString("cloudavenue_username")
		cloudavenuePassword = viper.GetString("cloudavenue_password")
		cloudavenueOrg = viper.GetString("cloudavenue_org")
		cloudavenueDebug = viper.GetBool("cloudavenue_debug")
	} else {
		cloudavenueUsername = os.Getenv("CLOUDAVENUE_USERNAME")
		cloudavenuePassword = os.Getenv("CLOUDAVENUE_PASSWORD")
		cloudavenueOrg = os.Getenv("CLOUDAVENUE_ORG")
		x, err := strconv.ParseBool(os.Getenv("CLOUDAVENUE_DEBUG"))
		if err != nil {
			return err
		}
		cloudavenueDebug = x
	}

	// Set client CloudAvenue
	var err error
	c, err = cloudavenue.New(cloudavenue.ClientOpts{
		CloudAvenue: &clientcloudavenue.Opts{
			Username: cloudavenueUsername,
			Password: cloudavenuePassword,
			Org:      cloudavenueOrg,
			Debug:    cloudavenueDebug,
		},
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
