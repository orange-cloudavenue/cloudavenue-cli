package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// createEdgeGatewayCmd create a edgegateway resource(s)
var createEdgeGatewayCmd = &cobra.Command{
	Use:               argEdgeGateway,
	Short:             "Create an edgeGateway",
	Aliases:           []string{"gw", "egw"},
	Example:           "edgegateway create --vdc <vdc name> [--t0 <t0 name>]",
	DisableAutoGenTag: true,
	SilenceErrors:     true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// init Config File & Client
		if err := initConfig(); err != nil {
			return fmt.Errorf("Unable to initialize: %w", err)
		}

		// Check if time flag is set
		if cmd.Flag(flagTime).Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the vdc name from the command line
		vdc, err := cmd.Flags().GetString(flagVDC)
		if err != nil {
			return fmt.Errorf("Unable to retrieve flag %v: %w", flagVDC, err)
		}

		// Get the t0 name
		// if flag is not precise, get the first one
		var t0 string
		if cmd.Flag(argT0).Value.String() == "" {
			t0s, err := c.V1.T0.GetT0s()
			if err != nil || (len(*t0s) > 1 || len(*t0s) == 0) {
				return fmt.Errorf("Unable to retrieve your first T0: %w", err)
			}
			t0 = (*t0s)[0].Tier0Vrf
		} else {
			t0, err = cmd.Flags().GetString("t0")
			if err != nil {
				return fmt.Errorf("Unable to retrieve T0 with VDC name %v: %w", flagVDC, err)
			}
		}
		// Create the edgeGateway
		s.Stop()
		fmt.Println("Creating EdgeGateway resource")
		fmt.Println("vdc name: " + vdc)
		fmt.Println("t0 name: " + t0)
		s.Restart()
		job, err := c.V1.EdgeGateway.New(vdc, t0)
		if err != nil {
			return fmt.Errorf("Unable to create job: %w", err)
		}
		err = job.Wait(3, 300)
		if err != nil {
			return fmt.Errorf("Error during EdgeGateway creation in VDC %v: %w", vdc, err)
		}
		s.FinalMSG = "EdgeGateway resource created successfully !!"
		s.Stop()
		return nil
	},
}
