package cmd

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output/model"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go"
	"github.com/spf13/cobra"
)

const (
	// ? Flags for arguments
	flagOutputValueJSON = "json"
	flagOutputValueYAML = "yaml"
	flagOutputValueWide = "wide"
	flagOutput          = "output"
	flagTime            = "time"
	flagName            = "name"
	flagIPAdress        = "ip"
	flagVDC             = "vdc"
	// ? Arguments for commands
	argVDC               = "vdc"
	argPublicIP          = "publicip"
	argPublicIPAlias1    = "pip"
	argS3                = "s3"
	argS3Alias           = "bucket"
	argEdgeGateway       = "edgegateway"
	argEdgeGatewayAlias1 = "gw"
	argEdgeGatewayAlias2 = "egw"
	argT0                = "t0"
	// Path of cav configuration file
	configPath     = "/.cav"
	fileConfig     = "config.yaml"
	fileConfigPath = configPath + "/" + fileConfig
)

var (
	c *cloudavenue.Client
	// Version of the CLI
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

// timeTrack function to print time elapsed for command
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("\n=== Command %s took %s second(s) ===\n", name, elapsed)
}
