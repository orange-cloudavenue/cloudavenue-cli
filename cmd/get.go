package cmd

import (
	"github.com/spf13/cobra"
)

const (
	cmdGet       = "get"
	cmdGetAlias1 = "ls"
	cmdGetAlias2 = "list"
	description  = `
					The default output format print the minimal necessary information like name, status or group.
					You can use the -o flag to specify the output format.
					"wide" will print some additional information.
					"json" or "yaml" will print the result in the specified format.`
)

var (
	exampleGet1 = `
	#List all T0
	cav get t0`
	exampleGet2 = `
	#List all T0 in wide format
	cav get t0 -o wide`
	exampleGet3 = `
	#List all Public IP
	cav get publicip`
	exampleGet4 = `
	#List all VDC in yaml format
	cav get vdc -o yaml`
	exampleGet5 = `
	#List all S3 in json format
	cav get s3 -o json`
)

// getCmd list a CAV resource
var getCmd = &cobra.Command{
	Use:               cmdGet,
	Aliases:           []string{cmdGetAlias1, cmdGetAlias2},
	Example:           exampleGet1 + "\n" + exampleGet2 + "\n" + exampleGet3 + "\n" + exampleGet4 + "\n" + exampleGet5,
	Short:             "Get resource to retrieve information from CloudAvenue.",
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(getCmd)

	// ? Get command
	getCmd.Args = cobra.NoArgs
	getCmd.AddCommand(getT0Cmd)
	getCmd.AddCommand(getPublicIPCmd)
	getCmd.AddCommand(getS3Cmd)
	getCmd.AddCommand(getEdgeGatewayCmd)
	getCmd.AddCommand(getVDCCmd)
	getCmd.PersistentFlags().StringP(flagOutput, "o", "", "Output format. One of: (wide, json, yaml)")
}
