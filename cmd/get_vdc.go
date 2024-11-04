package cmd

import (
	"fmt"
	"time"

	"github.com/orange-cloudavenue/common-go/print"
	"github.com/spf13/cobra"
	"github.com/vmware/go-vcloud-director/v2/types/v56"

	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/errorscustom"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output"
)

// getVDCCmd return a list of your vdc resource(s)
var getVDCCmd = &cobra.Command{
	Use:               argVDC,
	Short:             "A brief list of your vdc resources",
	Long:              "A complete list information of your s3 resources in your CloudAvenue account." + description,
	Example:           "get vdc",
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

		// Get the list of vdc or a specific vdc
		var vdcs []*types.QueryResultOrgVdcRecordType
		var vdc *types.QueryResultOrgVdcRecordType
		if cmd.Flag(flagName) != nil && cmd.Flag(flagName).Value.String() != "" {
			// Get the specific vdc
			vdc, err = c.V1.Querier().Get().VDC(cmd.Flag(flagName).Value.String())
			if err != nil || vdc == nil {
				return fmt.Errorf("CloudAvenue Error from VDC Get (VDC exist ?): %w", err)
			}
			// Create a list of one vdc
			vdcs = append(vdcs, vdc)
		} else {
			// Get the list of vdc
			vdcs, err = c.V1.Querier().List().VDC()
			if err != nil {
				return fmt.Errorf("CloudAvenue Error from VDC List: %w", err)
			}
		}

		// Print the result
		s.Stop()
		flag := cmd.Flag(flagOutput).Value.String()
		w := print.New()
		switch flag {
		case flagOutputValueWide:
			w.SetHeader("name", "status", "cpu used (mhz)", "memory used (mb)", "storage used (mb)", "number of vm(s)", "number of vapp(s)")
			for _, v := range vdcs {
				w.AddFields(v.Name, v.Status, *v.CpuUsedMhz, *v.MemoryUsedMB, *v.StorageUsedMB, *v.NumberOfVMs, *v.NumberOfVApps)
			}
		case flagOutputValueJSON, flagOutputValueYAML:
			x, err := output.New(stringToTypeFormat(flag), vdcs)
			if err != nil {
				return fmt.Errorf("impossible to format output: %w", err)
			}
			x.ToOutput()
		case "":
			w.SetHeader("name", "status")
			for _, v := range vdcs {
				w.AddFields(v.Name, v.Status)
			}
		default:
			return fmt.Errorf("output format %v: %w", flag, errorscustom.ErrNotValidOutput)
		}
		w.PrintTable()
		return nil
	},
}
