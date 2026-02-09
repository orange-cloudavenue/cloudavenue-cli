package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// deleteCmd delete a edgeGateway resource(s)
var delEdgeGatewayCmd = &cobra.Command{
	Use:               argEdgeGateway,
	Aliases:           []string{argEdgeGatewayAlias1, argEdgeGatewayAlias2},
	Example:           "del edgegateway <id or name> [<id or name>] [<id or name>] ...",
	Short:             "Delete an edgeGateway (name or id).",
	Long:              "Delete an edgeGateway (name or id), multiple edgeGateway can be deleted at once.",
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

		for _, arg := range args {
			s.Stop()
			fmt.Println("Delete EdgeGateway resource " + arg)
			s.Restart()
			gw, err := c.V1.EdgeGateway.Get(arg)
			if err != nil {
				return fmt.Errorf("unable to retrieve EdgeGateway: %w", err)
			}

			job, err := gw.Delete()
			if err != nil {
				return fmt.Errorf("unable to delete EdgeGateway: %w", err)
			}
			err = job.Wait(3, 300)
			if err != nil {
				return fmt.Errorf("error during jobs edgeGateway deletion : %w", err)
			}
			s.FinalMSG = "EdgeGateway resource deleted " + arg + " successfully !!\n"
			s.Stop()
		}
		return nil
	},
}
