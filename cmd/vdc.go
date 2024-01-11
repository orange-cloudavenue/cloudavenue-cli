/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/print"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go/v1/infrapi"
	"github.com/spf13/cobra"
)

const CmdVDC = "vdc"

// vdcCmd represents the vdc command
var vdcCmd = &cobra.Command{
	Use:     CmdVDC,
	Example: CmdVDC + " <list | create | delete>",
	Short:   "Option to manage your vdc (Virtual Data Center) on CloudAvenue.",
}

func init() {
	rootCmd.AddCommand(vdcCmd)

	// ? List command
	vdcCmd.Args = cobra.NoArgs
	vdcCmd.AddCommand(vdcListCmd)
	vdcCmd.PersistentFlags().StringP("output", "o", "", "Print all resources informations")

	// ? Delete command
	vdcCmd.AddCommand(vdcDelCmd)
	vdcDelCmd.Args = cobra.MinimumNArgs(1)

	// ? Create command
	vdcCmd.AddCommand(vdcCreateCmd)
	vdcCreateCmd.PersistentFlags().String("name", "", "vdc name")
	if err := vdcCreateCmd.MarkPersistentFlagRequired("name"); err != nil {
		fmt.Println("Error from Flag name, is require.", err)
		return
	}
}

// listCmd represents the list command
var vdcListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "get"},
	Short:   "A brief list of your vdc resources",
	Long:    "A complete list information of your s3 resources in your CloudAvenue account.",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the list of vdc
		vdcs, err := c.V1.Querier().List().VDC()
		if err != nil {
			fmt.Println("Error from VDC List", err)
			return
		}

		// Print the result
		flag := cmd.Flag("output").Value
		w := print.New()
		switch flag.String() {
		case "wide":
			w.SetHeader("name", "status", "cpu used (mhz)", "memory used (mb)", "storage used (mb)", "number of vm(s)", "number of vapp(s)")
			for _, v := range vdcs {
				w.AddFields(v.Name, v.Status, *v.CpuUsedMhz, *v.MemoryUsedMB, *v.StorageUsedMB, *v.NumberOfVMs, *v.NumberOfVApps)
			}
		default:
			w.SetHeader("name", "status")
			for _, v := range vdcs {
				w.AddFields(v.Name, v.Status)
			}
		}
		w.PrintTable()
	},
}

// deleteCmd represents the delete command
var vdcDelCmd = &cobra.Command{
	Use:     "delete",
	Example: "vdc delete <name> [<name>] [<name>] ...",
	Short:   "Delete a vdc",

	Run: func(cmd *cobra.Command, args []string) {
		defer timeTrack(time.Now(), cmd.CommandPath())

		for _, arg := range args {
			fmt.Println("delete vdc resource " + arg)
			// vdc, err := c.V1.VDC.Get(arg)
			vdc, err := c.V1.VDC().GetVDC(arg)
			if err != nil {
				fmt.Println("Error from vdc", err)
				return
			}
			job, err := vdc.Delete()
			if err != nil {
				fmt.Println("Unable to delete vdc", err)
				return
			}
			err = job.Wait(3, 300)
			if err != nil {
				fmt.Println("Error during vdc Deletion !!", err)
				return
			}
			fmt.Println("vdc resource deleted " + arg + " successfully !!")
		}
		fmt.Println("\nvdc resource list after deletion:")
		vdcListCmd.Run(cmd, []string{})

	},
}

// createCmd represents the create command
var vdcCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an vdc",
	Example: "vdc create --name <vdc name>",

	Run: func(cmd *cobra.Command, args []string) {
		defer timeTrack(time.Now(), cmd.CommandPath())

		// Get the vdc name from the command line
		vdcName, err := cmd.Flags().GetString("vdc")
		if err != nil {
			fmt.Println("Malformed VDC name", err)
			return
		}

		// Create the vdc
		fmt.Println("create vdc resource (with basic value)")
		fmt.Println("vdc name: " + vdcName)

		_, err = c.V1.VDC().New(&infrapi.CAVVirtualDataCenter{VDC: infrapi.CAVVirtualDataCenterVDC{
			Name:                vdcName,
			ServiceClass:        "STD",
			BillingModel:        "PAYG",
			CPUAllocated:        22000,
			VCPUInMhz:           2200,
			Description:         "vdc created by cloudavenue-cli",
			MemoryAllocated:     30,
			DisponibilityClass:  "ONE-ROOM",
			StorageBillingModel: "PAYG",
			StorageProfiles: []infrapi.StorageProfile{
				{
					Class:   "gold",
					Limit:   500,
					Default: true,
				},
			},
		},
		})

		if err != nil {
			fmt.Println("Error from vdc", err)
			return
		}

		fmt.Println("vdc resource created successfully !")
		fmt.Println("\nvdc resource list after creation:")
		vdcListCmd.Run(cmd, []string{})

	},
}
