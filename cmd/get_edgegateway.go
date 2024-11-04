package cmd

import (
	"fmt"
	"time"

	v1 "github.com/orange-cloudavenue/cloudavenue-sdk-go/v1"
	"github.com/orange-cloudavenue/common-go/print"
	"github.com/spf13/cobra"

	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/errorscustom"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output"
)

// getEdgeGatewayCmd return a list of your edgegateway resource(s)
var getEdgeGatewayCmd = &cobra.Command{
	Use:               argEdgeGateway,
	Aliases:           []string{argEdgeGatewayAlias1, argEdgeGatewayAlias2},
	Short:             "A brief list of your edgegateway resources",
	Long:              "A complete list information of your EdgeGateway resources in your CloudAvenue account." + description,
	Example:           "get edgegateway",
	DisableAutoGenTag: true,
	SilenceErrors:     true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		// init Config File & Client
		if err := initConfig(); err != nil {
			return fmt.Errorf("unable to initialize: %w", err)
		}

		// Check if time flag is set and print time elapsed
		if cmd.Flag(flagTime).Value.String() == trueValue {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the list of edgegateway or a specific edgegateway
		var edgeGateways *v1.EdgeGateways
		var edgeGw *v1.EdgeGw
		if cmd.Flag(flagName) != nil && cmd.Flag(flagName).Value.String() != "" {
			// Get the specific edgegateway
			edgeGw, err = c.V1.EdgeGateway.GetByName(cmd.Flag(flagName).Value.String())
			if err != nil {
				return fmt.Errorf("CloudAvenue Error from EdgeGateway Get: %w", err)
			}
			// Create a list of one edgegateway
			edgeGateways = &v1.EdgeGateways{*edgeGw}
		} else {
			// Get the list of edgegateway
			edgeGateways, err = c.V1.EdgeGateway.List()
			if err != nil {
				return fmt.Errorf("CloudAvenue Error from EdgeGateway List: %w", err)
			}
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
				return fmt.Errorf("impossible to format output: %w", err)
			}
			x.ToOutput()
		case "":
			w.SetHeader("name", "owner")
			for _, e := range *edgeGateways {
				w.AddFields(e.EdgeName, e.OwnerName)
			}
		default:
			return fmt.Errorf("output format %v: %w", flag, errorscustom.ErrNotValidOutput)
		}
		w.PrintTable()
		return nil
	},
}
