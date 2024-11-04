package cmd

import (
	"github.com/spf13/cobra"
)

const (
	cmdDelete = "del"
)

var (
	exampleDel1 = `
	#Delete a Public IP
	cav del ip 192.168.0.2`
	exampleDel2 = `
	#Delete several vdc named xxxx and yyyy
	cav del vdc xxxx yyyy`
	exampleDel3 = `
	#Delete a edgegateway named zzzz
	cav del egw zzzz`
)

// delCmd delete a CAV resource
var delCmd = &cobra.Command{
	Use:               cmdDelete,
	Example:           exampleDel1 + "\n" + exampleDel2 + "\n" + exampleDel3,
	Short:             "Delete resource(s) from CloudAvenue.",
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(delCmd)

	// ? Delete command
	delCmd.Args = cobra.NoArgs

	// ? Delete for public ip
	delCmd.AddCommand(delPublicIPCmd)
	delPublicIPCmd.Args = cobra.MinimumNArgs(1)

	// ? Delete for s3
	delCmd.AddCommand(delS3Cmd)
	delS3Cmd.Args = cobra.MinimumNArgs(1)

	// ? Delete for edgegateway
	delCmd.AddCommand(delEdgeGatewayCmd)
	delEdgeGatewayCmd.Args = cobra.MinimumNArgs(1)

	// ? Delete for vdc
	delCmd.AddCommand(delVDCCmd)
	delVDCCmd.Args = cobra.MinimumNArgs(1)
}
