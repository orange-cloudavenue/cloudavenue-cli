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

// getBMSCmd return a list of your BMS (bucket) resource(s)
var getBMSCmd = &cobra.Command{
	Use:               argBMS,
	Short:             "A brief list of your BMS resources",
	Long:              "A complete list information of your BMS resources in your CloudAvenue account." + description,
	Example:           "get BMS",
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

		listBMS, err := c.V1.BMS.List()
		if err != nil {
			return fmt.Errorf("CloudAvenue Error from BMS List: %w", err)
		}
		if cmd.Flag(flagName) != nil && cmd.Flag(flagName).Value.String() != "" {
			// Get the specific bms
			for _, b := range *listBMS {
				bmsDetail := b.GetBMS()
				for _, bms := range bmsDetail {
					if bms.Hostname == cmd.Flag(flagName).Value.String() {
						// Create a listBMS with this bms found
						x := v1.BMS{
							BMSNetworks: []v1.BMSNetwork{},
							BMSDetails:  []v1.BMSDetail{bms},
						}
						listBMS = &[]v1.BMS{x}
					}
				}
			}
		}

		// Print the result
		s.Stop()
		flag := cmd.Flag(flagOutput).Value.String()
		w := print.New()
		// var format output.Formatter
		switch flag {
		case flagOutputValueWide:
			w.SetHeader("name", "owner", "creation date", "owner id")
			for _, b := range *listBMS {
				bmsDetail := b.GetBMS()
				for _, bms := range bmsDetail {
					w.AddFields(bms.Hostname, bms.BMSType, bms.OS, bms.BiosConfiguration)
				}
			}
		case flagOutputValueJSON, flagOutputValueYAML:
			x, err := output.New(stringToTypeFormat(flag), listBMS)
			if err != nil {
				return fmt.Errorf("impossible to format output: %w", err)
			}
			x.ToOutput()
		case "":
			w.SetHeader("name", "owner")
			for _, b := range *listBMS {
				bmsDetail := b.GetBMS()
				for _, bms := range bmsDetail {
					w.AddFields(bms.Hostname, bms.OS)
				}
			}
		default:
			return fmt.Errorf("output format %v: %w", flag, errorscustom.ErrNotValidOutput)
		}
		w.PrintTable()
		return nil
	},
}
