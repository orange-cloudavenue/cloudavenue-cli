package cmd

import (
	"fmt"
	"time"

	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/customErrors"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output"
	"github.com/orange-cloudavenue/common-go/print"
	"github.com/spf13/cobra"
)

// getEdgeGatewayCmd return a list of your edgegateway resource(s)
var getEdgeGatewayCmd = &cobra.Command{
	Use:               argEdgeGateway,
	Aliases:           []string{argEdgeGatewayAlias1, argEdgeGatewayAlias2},
	Short:             "A brief list of your edgegateway resources",
	Long:              "A complete list information of your EdgeGateway resources in your CloudAvenue account." + description,
	DisableAutoGenTag: true,
	SilenceErrors:     true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// init Config File & Client
		if err := initConfig(); err != nil {
			return fmt.Errorf("Unable to initialize: %w", err)
		}

		// Check if time flag is set and print time elapsed
		if cmd.Flag(flagTime).Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		edgeGateways, err := c.V1.EdgeGateway.List()
		if err != nil {
			return fmt.Errorf("CloudAvenue Error from EdgeGateway List: %w", err)
		}

		// Print the result
		s.Stop()
		flag := cmd.Flag(flagOutput).Value.String()
		w := print.New()
		switch flag {
		case flagOutputValueWide:
			w.SetHeader("name", "id", "owner name", "owner type", "ratelimit (mb/s)", "description", "tier0 vrf name")
			for _, e := range *edgeGateways {
				w.AddFields(e.EdgeName, e.EdgeID, e.OwnerName, e.OwnerType, e.Bandwidth, e.Description, e.Tier0VrfName)
			}
		case flagOutputValueJSON, flagOutputValueYAML:
			x, err := output.New(stringToTypeFormat(flag), edgeGateways)
			if err != nil {
				return fmt.Errorf("Impossible to format output: %w", err)
			}
			x.ToOutput()
		case "":
			w.SetHeader("name", "owner")
			for _, e := range *edgeGateways {
				w.AddFields(e.EdgeName, e.OwnerName)
			}
		default:
			return fmt.Errorf("Output format %v: %w", flag, customErrors.ErrNotValidOutput)
		}
		w.PrintTable()
		return nil
	},
}
