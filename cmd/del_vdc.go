package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// delVDCCmd delete a vdc resource(s)
var delVDCCmd = &cobra.Command{
	Use:               argVDC,
	Example:           "del vdc <name> [<name>] [<name>] ...",
	Short:             "Delete a vdc",
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

		for _, arg := range args {
			s.Stop()
			fmt.Println("delete vdc resource " + arg)
			s.Restart()
			vdc, err := c.V1.VDC().GetVDC(arg)
			if err != nil {
				return fmt.Errorf("Error from vdc: %w", err)
			}
			job, err := vdc.Delete()
			if err != nil {
				return fmt.Errorf("Unable to delete vdc: %w", err)
			}
			err = job.Wait(3, 300)
			if err != nil {
				return fmt.Errorf("Error during vdc Deletion !!: %w", err)
			}
			s.FinalMSG = "vdc resource deleted " + arg + " successfully !!\n"
			s.Stop()
		}
		return nil
	},
}
