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

// getT0Cmd return a list of your t0 resource(s)
var getT0Cmd = &cobra.Command{
	Use:               argT0,
	Short:             "A brief list of your t0 resources",
	Long:              "A complete list information of your T0 resources in your CloudAvenue account." + description,
	Example:           "get t0",
	DisableAutoGenTag: true,
	SilenceErrors:     true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		// init Config File & Client
		if err = initConfig(); err != nil {
			return fmt.Errorf("unable to initialize: %w", err)
		}

		// Check if time flag is set and print time elapsed
		if cmd.Flag(flagTime).Value.String() == trueValue {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the list of t0 or a specific t0
		var t0s *v1.T0s
		var t0 *v1.T0
		if cmd.Flag(flagName) != nil && cmd.Flag(flagName).Value.String() != "" {
			// Get the specific t0
			t0, err = c.V1.T0.GetT0(cmd.Flag(flagName).Value.String())
			if err != nil {
				return fmt.Errorf("CloudAvenue Error from T0 Get: %w", err)
			}
			// Create a list of one t0
			t0s = &v1.T0s{*t0}
		} else {
			// Get the list of t0
			t0s, err = c.V1.T0.GetT0s()
			if err != nil {
				return fmt.Errorf("CloudAvenue Error from T0 List: %w", err)
			}
		}

		// Print the result
		s.Stop()
		flag := cmd.Flag(flagOutput).Value.String()
		w := print.New()
		switch flag {
		case flagOutputValueWide:
			w.SetHeader("name", "t0 provider name", "t0 class service", "services", "class service")
			for _, t0 := range *t0s {
				w.AddFields(t0.Tier0Vrf, t0.Tier0Provider, t0.Tier0ClassService, t0.Services, t0.ClassService)
			}
		case flagOutputValueJSON, flagOutputValueYAML:
			x, err := output.New(stringToTypeFormat(flag), t0s)
			if err != nil {
				return fmt.Errorf("impossible to format output: %w", err)
			}
			x.ToOutput()
		case "":
			w.SetHeader("name", "t0 provider name")
			for _, t0 := range *t0s {
				w.AddFields(t0.Tier0Vrf, t0.Tier0Provider)
			}
		default:
			return fmt.Errorf("output format %v: %w", flag, errorscustom.ErrNotValidOutput)
		}
		w.PrintTable()
		return nil
	},
}
