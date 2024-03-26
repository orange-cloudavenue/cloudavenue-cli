package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	cmdAdd = "add"
)

var (
	exampleAdd1 = `
	#Add a VDC (Virtual Data Center) to CloudAvenue
	cav add vdc --name myvdc`
)

// addCmd add a CAV resource
var addCmd = &cobra.Command{
	Use:               cmdAdd,
	Example:           exampleAdd1,
	Short:             "Add resource to CloudAvenue.",
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// ? Add command
	addCmd.Args = cobra.NoArgs
	// ? Add for public ip
	addCmd.AddCommand(addPublicIPCmd)
	// ? Add for s3
	addCmd.AddCommand(addS3Cmd)
	// ? add for edgegateway
	addCmd.AddCommand(addEdgeGatewayCmd)
	// ? Add for vdc
	addCmd.AddCommand(addVDCCmd)

	// ? Options for edgegateway
	addEdgeGatewayCmd.Flags().String(argT0, "", "t0 name")
	addEdgeGatewayCmd.Flags().String(argVDC, "", "vdc name")
	if err := addEdgeGatewayCmd.MarkFlagRequired(argVDC); err != nil {
		fmt.Println("Error from Flag VDC, is require.", err)
		return
	}

	// ? Options for publicip
	addPublicIPCmd.Flags().String(flagName, "", "edgegateway name")
	if err := addPublicIPCmd.MarkFlagRequired(flagName); err != nil {
		fmt.Println("Error from Flag name, is require.", err)
		return
	}

	// ? Options for s3
	addS3Cmd.Flags().String(flagName, "", "s3 bucket name")
	if err := addS3Cmd.MarkFlagRequired(flagName); err != nil {
		fmt.Println("Error from Flag name, is require.", err)
		return
	}

	// ? Options for vdc
	addVDCCmd.Flags().String(flagName, "", "vdc name")
	if err := addVDCCmd.MarkFlagRequired(flagName); err != nil {
		fmt.Println("Error from Flag Name, is require.", err)
		return
	}
}
