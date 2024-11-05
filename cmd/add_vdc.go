package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/orange-cloudavenue/cloudavenue-sdk-go/v1/infrapi"
	"github.com/spf13/cobra"
)

// addVDCCmd add a vdc resource(s)
var addVDCCmd = &cobra.Command{
	Use:               argVDC,
	Short:             "Add an vdc (virtual data center) to CloudAvenue.",
	Example:           "add vdc --name <vdc name>",
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
		vdcName, err := cmd.Flags().GetString(flagName)
		if err != nil {
			return fmt.Errorf("unable to retrieve flag %v: %w", flagName, err)
		}

		// Add the vdc
		s.Stop()
		fmt.Println("add vdc resource (with basic value)")
		fmt.Println("vdc name: " + vdcName)
		s.Restart()

		if _, err = c.V1.VDC().New(context.Background(), &infrapi.CAVVirtualDataCenter{
			VDC: infrapi.CAVVirtualDataCenterVDC{
				Name:                vdcName,
				ServiceClass:        "STD",
				BillingModel:        "PAYG",
				CPUAllocated:        22000,
				VCPUInMhz:           2200,
				Description:         "vdc add by cloudavenue-cli",
				MemoryAllocated:     30,
				DisponibilityClass:  "ONE-ROOM",
				StorageBillingModel: "PAYG",
				StorageProfiles: []infrapi.StorageProfile{
					{
						Class:   "gold",
						Limit:   500,
						Default: true,
					},
				},
			},
		}); err != nil {
			return fmt.Errorf("unable to add vdc %v: %w", vdcName, err)
		}
		s.FinalMSG = "vdc resource added successfully !!\n"
		s.Stop()
		return nil
	},
}
