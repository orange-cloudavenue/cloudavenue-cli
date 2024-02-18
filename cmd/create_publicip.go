package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// createPublicIPCmd create a public ip resource(s)
var createPublicIPCmd = &cobra.Command{
	Use:               argPublicIP,
	Short:             "Create an ip",
	Example:           "ip create --name <EdgeGateway>",
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

		// Get the name from the command line
		gwName, err := cmd.Flags().GetString(flagName)
		if err != nil {
			return fmt.Errorf("Unable to retrieve flag %v: %w", flagName, err)
		}

		// Create a public ip
		s.Stop()
		fmt.Println("create public ip resource")
		fmt.Println("for EdgeGateway name: " + gwName)
		s.Restart()

		job, err := c.V1.PublicIP.New(gwName)
		if err != nil {
			return fmt.Errorf("Unable to create job: %w", err)
		}
		err = job.Wait(5, 300)
		if err != nil {
			return fmt.Errorf("Error during public ip creation for edgegateway %v: %w", gwName, err)
		}
		s.FinalMSG = "public ip resource created successfully !!"
		s.Stop()
		return nil
	},
}
