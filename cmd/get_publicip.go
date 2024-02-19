package cmd

import (
	"fmt"
	"time"

	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/customErrors"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output"
	"github.com/orange-cloudavenue/common-go/print"
	"github.com/spf13/cobra"
)

// getPublicIPCmd return a list of your publicip
var getPublicIPCmd = &cobra.Command{
	Use:               argPublicIP,
	Aliases:           []string{argPublicIPAlias1},
	Short:             "A brief list of your public ip resources",
	Long:              "A complete list information of your Public IP resources in your CloudAvenue account." + description,
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

		// Get the list of publicip
		ips, err := c.V1.PublicIP.GetIPs()
		if err != nil {
			return fmt.Errorf("CloudAvenue Error from IP List: %w", err)
		}

		// Print the result
		s.Stop()
		flag := cmd.Flag(flagOutput).Value.String()
		w := print.New()
		switch flag {
		case flagOutputValueWide:
			w.SetHeader("public ip", "edge gateway name", "ip natted")
			for _, i := range ips.NetworkConfig {
				w.AddFields(i.UplinkIP, i.EdgeGatewayName, i.TranslatedIP)
			}
		case flagOutputValueJSON, flagOutputValueYAML:
			x, err := output.New(stringToTypeFormat(flag), ips)
			if err != nil {
				return fmt.Errorf("Impossible to format output: %w", err)
			}
			x.ToOutput()
		case "":
			w.SetHeader("public ip", "edge gateway name")
			for _, i := range ips.NetworkConfig {
				w.AddFields(i.UplinkIP, i.EdgeGatewayName)
			}
		default:
			return fmt.Errorf("Output format %v: %w", flag, customErrors.ErrNotValidOutput)
		}
		w.PrintTable()
		return nil
	},
}
