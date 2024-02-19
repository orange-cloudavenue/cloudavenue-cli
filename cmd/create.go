package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	cmdCreate       = "create"
	cmdCreateAlias1 = "new"
	cmdCreateAlias2 = "add"
)

var (
	exampleCreate1 = `
	#List all T0
	cav create vdc --name myvdc`
)

// createCmd create a CAV resource
var createCmd = &cobra.Command{
	Use:               cmdCreate,
	Aliases:           []string{cmdCreateAlias1, cmdCreateAlias2},
	Example:           exampleCreate1,
	Short:             "Create resource to CloudAvenue.",
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(createCmd)

	// ? Create command
	createCmd.Args = cobra.NoArgs
	// ? Create for public ip
	createCmd.AddCommand(createPublicIPCmd)
	// ? Create for s3
	createCmd.AddCommand(createS3Cmd)
	// ? Create for edgegateway
	createCmd.AddCommand(createEdgeGatewayCmd)
	// ? Create for vdc
	createCmd.AddCommand(createVDCCmd)

	// ? Options for edgegateway
	createEdgeGatewayCmd.Flags().String(argT0, "", "t0 name")
	createEdgeGatewayCmd.Flags().String(argVDC, "", "vdc name")
	if err := createEdgeGatewayCmd.MarkFlagRequired(argVDC); err != nil {
		fmt.Println("Error from Flag VDC, is require.", err)
		return
	}

	// ? Options for publicip
	createPublicIPCmd.Flags().String(flagName, "", "edgegateway name")
	if err := createPublicIPCmd.MarkFlagRequired(flagName); err != nil {
		fmt.Println("Error from Flag name, is require.", err)
		return
	}

	// ? Options for s3
	createS3Cmd.Flags().String(flagName, "", "s3 bucket name")
	if err := createS3Cmd.MarkFlagRequired(flagName); err != nil {
		fmt.Println("Error from Flag name, is require.", err)
		return
	}

	// ? Options for vdc
	createVDCCmd.Flags().String(flagName, "", "vdc name")
	if err := createVDCCmd.MarkFlagRequired(flagName); err != nil {
		fmt.Println("Error from Flag Name, is require.", err)
		return
	}
}
