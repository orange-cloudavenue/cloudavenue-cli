package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// addPublicIPCmd add a public ip resource(s)
var addPublicIPCmd = &cobra.Command{
	Use:               argPublicIP,
	Aliases:           []string{argPublicIPAlias1},
	Short:             "Add a public ip",
	Example:           "add publicip --name <EdgeGateway>",
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

		// Get the name from the command line
		gwName, err := cmd.Flags().GetString(flagName)
		if err != nil {
			return fmt.Errorf("unable to retrieve flag %v: %w", flagName, err)
		}

		// Add a public ip
		s.Stop()
		fmt.Println("add public ip resource")
		fmt.Println("for EdgeGateway name: " + gwName)
		s.Restart()

		job, err := c.V1.PublicIP.New(gwName)
		if err != nil {
			return fmt.Errorf("unable to add job: %w", err)
		}
		err = job.Wait(5, 300)
		if err != nil {
			return fmt.Errorf("error during public ip creation for edgegateway %v: %w", gwName, err)
		}
		s.FinalMSG = "public ip resource added successfully !!"
		s.Stop()
		return nil
	},
}
