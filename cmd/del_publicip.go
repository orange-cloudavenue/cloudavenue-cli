package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// deleteCmd delete a public ip resource(s)
var delPublicIPCmd = &cobra.Command{
	Use:               argPublicIP,
	Aliases:           []string{argPublicIPAlias1},
	Example:           "delete publicip <ip> [<ip>] [<ip>] ...",
	Short:             "Delete public ip resource(s)",
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
			fmt.Println("Delete publicip resource " + arg)
			s.Restart()
			ip, err := c.V1.PublicIP.GetIP(arg)
			if err != nil {
				return fmt.Errorf("Unable to retrieve ip: %w", err)
			}
			job, err := ip.Delete()
			if err != nil {
				return fmt.Errorf("Unable to delete ip: %w", err)
			}
			err = job.Wait(15, 300)
			if err != nil {
				return fmt.Errorf("Job errors: %w", err)
			}
			s.FinalMSG = "ip resource deleted " + arg + " successfully !!\n"
			s.Stop()
		}
		return nil
	},
}
