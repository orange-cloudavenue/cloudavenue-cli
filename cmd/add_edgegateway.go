package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// addEdgeGatewayCmd add a edgegateway resource(s)
var addEdgeGatewayCmd = &cobra.Command{
	Use:               argEdgeGateway,
	Aliases:           []string{argEdgeGatewayAlias1, argEdgeGatewayAlias2},
	Short:             "Add an edgeGateway",
	Long:              "Add an edgeGateway in a VDC. If the T0 is not specified, the first one will be used. No need to specify a name, the edgeGateway name is auto-generated.",
	Example:           "add edgegateway --vdc <vdc name> [--t0 <t0 name>]",
	DisableAutoGenTag: true,
	SilenceErrors:     true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// init Config File & Client
		if err := initConfig(); err != nil {
			return fmt.Errorf("unable to initialize: %w", err)
		}

		// Check if time flag is set
		if cmd.Flag(flagTime).Value.String() == trueValue {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the vdc name from the command line
		vdc, err := cmd.Flags().GetString(flagVDC)
		if err != nil {
			return fmt.Errorf("unable to retrieve flag %v: %w", flagVDC, err)
		}

		// Get the t0 name
		// if flag is not precise, get the first one
		var t0 string
		if cmd.Flag(argT0).Value.String() == "" {
			t0s, err := c.V1.T0.GetT0s()
			if err != nil || (len(*t0s) > 1 || len(*t0s) == 0) {
				return fmt.Errorf("unable to retrieve your first T0: %w", err)
			}
			t0 = (*t0s)[0].Tier0Vrf
		} else {
			t0, err = cmd.Flags().GetString("t0")
			if err != nil {
				return fmt.Errorf("unable to retrieve T0 with VDC name %v: %w", flagVDC, err)
			}
		}
		// Add the edgeGateway
		s.Stop()
		fmt.Println("Creating EdgeGateway resource")
		fmt.Println("vdc name: " + vdc)
		fmt.Println("t0 name: " + t0)
		s.Restart()
		job, err := c.V1.EdgeGateway.New(vdc, t0)
		if err != nil {
			return fmt.Errorf("unable to add job: %w", err)
		}
		err = job.Wait(3, 300)
		if err != nil {
			return fmt.Errorf("error during EdgeGateway creation in VDC %v: %w", vdc, err)
		}
		s.FinalMSG = "EdgeGateway resource added successfully !!"
		s.Stop()
		return nil
	},
}
