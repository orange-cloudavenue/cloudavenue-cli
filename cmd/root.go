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
	"github.com/mitchellh/go-homedir"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output/model"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go"
	clientcloudavenue "github.com/orange-cloudavenue/cloudavenue-sdk-go/pkg/clients/cloudavenue"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagOutputValueJSON  = "json"
	flagOutputValueYAML  = "yaml"
	flagOutputValueWide  = "wide"
	flagOutput           = "output"
	flagTime             = "time"
	flagName             = "name"
	argVDC               = "vdc"
	argPublicIP          = "publicip"
	argPublicIPAlias1    = "ip"
	argS3                = "s3"
	argS3Alias           = "bucket"
	argEdgeGateway       = "edgegateway"
	argEdgeGatewayAlias1 = "gw"
	argEdgeGatewayAlias2 = "egw"
	argT0                = "t0"
)

var (
	c       *cloudavenue.Client
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
	s       = spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	RootCmd = rootCmd

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:               "cav",
		Short:             "cav is the Command Line Interface for CloudAvenue Platform",
		DisableAutoGenTag: true,
	}
)

// Use for YAML configuration file
type cloudavenueConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Org      string `yaml:"org"`
	URL      string `yaml:"url"`
	Debug    bool   `yaml:"debug"`
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() (err error) {
	// ctrl+c handler
	sigint.ListenForSIGINT(func() {
		fmt.Println("SIGINT received. Exiting...")
		os.Exit(0)
	})

	// Execute root command
	if err = rootCmd.Execute(); err != nil {
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
	rootCmd.InitDefaultCompletionCmd()
}

// function to transform String to output.TypeFormat
func stringToTypeFormat(s string) model.TypeFormat {
	switch s {
	case flagOutputValueJSON:
		return model.TypeJSON
	case flagOutputValueYAML:
		return model.TypeYAML
	default:
		return model.TypeFormat("")
	}
}

// function to initialize the configuration file
func initConfig() (*viper.Viper, error) {
	// Set default file configuration and create it if not exist
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	// Set default file configuration
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(home + "/.cav")
	if home == "" {
		return nil, fmt.Errorf("Error in Get HOME Directory")
	}
	// Create configuration file if not exist
	if _, err = os.Stat(home + "/.cav/config.yaml"); os.IsNotExist(err) {
		if err = os.MkdirAll(home+"/.cav", 0755); err != nil {
			return nil, err
		}
		// Set default configuration
		cloudavenueConfig := cloudavenueConfig{}
		// set struct to viper
		v.Set("cloudavenue", cloudavenueConfig)

		// Write configuration file
		if err = v.SafeWriteConfig(); err != nil {
			return nil, err
		}
		s.FinalMSG = `
					***
					Configuration file is created in " + home + "/.cav/config.yaml
					Please fill it with your credentials and re-run the command.
					***`
		s.Stop()
		os.Exit(0)
	}
	return v, nil
}

func initClient(v *viper.Viper) (err error) {
	// Read configuration file
	err = v.ReadInConfig()
	if err != nil {
		fmt.Println("Unable to read config:", err)
	}
	// Set client CloudAvenue
	c, err = cloudavenue.New(&cloudavenue.ClientOpts{
		CloudAvenue: &clientcloudavenue.Opts{
			Username: v.GetString("cloudavenue.username"),
			Password: v.GetString("cloudavenue.password"),
			Org:      v.GetString("cloudavenue.org"),
			Debug:    v.GetBool("cloudavenue.debug"),
		},
	})
	if err != nil {
		return err
	}
	return nil
}
