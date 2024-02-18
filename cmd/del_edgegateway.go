package cmd

import (
	"fmt"
	"time"

	"github.com/orange-cloudavenue/cloudavenue-sdk-go/pkg/uuid"
	v1 "github.com/orange-cloudavenue/cloudavenue-sdk-go/v1"
	"github.com/spf13/cobra"
)

// deleteCmd delete a edgeGateway resource(s)
var delEdgeGatewayCmd = &cobra.Command{
	Use:               argEdgeGateway,
	Aliases:           []string{argEdgeGatewayAlias1, argEdgeGatewayAlias2},
	Example:           "delete edgegateway <id or name> [<id or name>] [<id or name>] ...",
	Short:             "Delete an edgeGateway (name or id)",
	DisableAutoGenTag: true,
	SilenceErrors:     true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			gw  *v1.EdgeGw
			err error
		)

		// init Config File & Client
		if err = initConfig(); err != nil {
			return fmt.Errorf("Unable to initialize: %w", err)
		}

		// Check if time flag is set
		if cmd.Flag(flagTime).Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		for _, arg := range args {
			s.Stop()
			fmt.Println("Delete EdgeGateway resource " + arg)
			s.Restart()
			if uuid.IsUUIDV4(arg) {
				gw, err = c.V1.EdgeGateway.GetByID(arg)
			} else {
				gw, err = c.V1.EdgeGateway.GetByName(arg)
			}
			if err != nil {
				return fmt.Errorf("Unable to retrieve EdgeGateway: %w.", err)
			}

			job, err := gw.Delete()
			if err != nil {
				return fmt.Errorf("Unable to delete EdgeGateway: %w", err)
			}
			err = job.Wait(3, 300)
			if err != nil {
				return fmt.Errorf("Error during jobs edgeGateway deletion : %w", err)
			}
			s.FinalMSG = "EdgeGateway resource deleted " + arg + " successfully !!\n"
			s.Stop()
		}
		return nil
	},
}
