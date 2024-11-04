package cmd

import (
	"fmt"
	"time"

	v1 "github.com/orange-cloudavenue/cloudavenue-sdk-go/v1"
	"github.com/orange-cloudavenue/common-go/print"
	"github.com/spf13/cobra"

	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/customerrors"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output"
)

// getPublicIPCmd return a list of your publicip
var getPublicIPCmd = &cobra.Command{
	Use:               argPublicIP,
	Aliases:           []string{argPublicIPAlias1},
	Short:             "A brief list of your public ip resources",
	Long:              "A complete list information of your Public IP resources in your CloudAvenue account." + description,
	Example:           "get publicip",
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

		// Get the list of publicip or a specific publicip
		var ips *v1.IPs
		var ip *v1.IP
		if cmd.Flag(flagIPAdress) != nil && cmd.Flag(flagIPAdress).Value.String() != "" {
			// Get the specific publicip
			ip, err = c.V1.PublicIP.GetIP(cmd.Flag(flagIPAdress).Value.String())
			if err != nil {
				return fmt.Errorf("CloudAvenue Error from IP Get: %w", err)
			}
			// Create a list of one publicip
			ips = &v1.IPs{
				NetworkConfig: []v1.IP{*ip},
			}
		} else {
			// Get the list of publicip
			ips, err = c.V1.PublicIP.GetIPs()
			if err != nil {
				return fmt.Errorf("CloudAvenue Error from IP List: %w", err)
			}
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
				return fmt.Errorf("impossible to format output: %w", err)
			}
			x.ToOutput()
		case "":
			w.SetHeader("public ip", "edge gateway name")
			for _, i := range ips.NetworkConfig {
				w.AddFields(i.UplinkIP, i.EdgeGatewayName)
			}
		default:
			return fmt.Errorf("output format %v: %w", flag, customerrors.ErrNotValidOutput)
		}
		w.PrintTable()
		return nil
	},
}
