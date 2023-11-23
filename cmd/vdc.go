/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	jsontmpl "github.com/orange-cloudavenue/cloudavenue-cli/pkg/templates/json"
	v1 "github.com/orange-cloudavenue/cloudavenue-sdk-go/v1"
	"github.com/spf13/cobra"
)

const VDC = "vdc"

// vdcCmd represents the vdc command
var vdcCmd = &cobra.Command{
	Use:     VDC,
	Example: VDC + " <list | create | delete>",
	Short:   "Option to manage your vdc (Virtual Data Center) on CloudAvenue.",
}

func init() {
	rootCmd.AddCommand(vdcCmd)

	// ? List command
	vdcCmd.Args = cobra.NoArgs
	vdcCmd.AddCommand(vdcListCmd)

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
	Use:   "list",
	Short: "A brief list of your vdc resources",
	Run: func(cmd *cobra.Command, args []string) {

		// Get the list of vdc
		vdcs, err := c.V1.VDC.List()
		if err != nil {
			fmt.Println("Error from VDC List", err)
			return
		}

		// Struct to print a basic view
		type basicVdc = struct {
			VdcGroup        string `json:"vdc_group"`
			VdcName         string `json:"vdc_name"`
			VcpuInMhz2      int    `json:"vcpu_in_mhz2"`
			MemoryAllocated int    `json:"memory_allocated"`
			CPUAllocated    int    `json:"cpu_allocated"`
			Description     string `json:"description"`
		}
		basicVdcs := []*basicVdc{}

		// Set the struct
		for _, dc := range *vdcs {
			x := &basicVdc{
				VdcGroup:        dc.VdcGroup,
				VdcName:         dc.Vdc.Name,
				VcpuInMhz2:      dc.Vdc.VcpuInMhz2,
				MemoryAllocated: dc.Vdc.MemoryAllocated,
				CPUAllocated:    dc.Vdc.CPUAllocated,
				Description:     dc.Vdc.Description,
			}
			basicVdcs = append(basicVdcs, x)
		}

		// Print the result
		jsontmpl.Format(jsontmpl.JsonTemplate{
			Fields: []string{"vdc_name", "vdc_group", "vcpu_in_mhz2", "memory_allocated", "cpu_allocated", "description"},
			Data:   basicVdcs,
		})
	},
}

// deleteCmd represents the delete command
var vdcDelCmd = &cobra.Command{
	Use:     "delete",
	Example: "vdc delete <name> [<name>] [<name>] ...",
	Short:   "Delete a vdc",

	Run: func(cmd *cobra.Command, args []string) {

		for _, arg := range args {
			fmt.Println("delete vdc resource " + arg)
			vdc, err := c.V1.VDC.Get(arg)
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
	Short:   "Ceate an vdc",
	Example: "vdc create --name <vdc name>",

	Run: func(cmd *cobra.Command, args []string) {

		// Get the vdc name from the command line
		vdcName, err := cmd.Flags().GetString("vdc")
		if err != nil {
			fmt.Println("Malformed VDC name", err)
			return
		}

		// Create the vdc
		fmt.Println("create vdc resource (with basic value)")
		fmt.Println("vdc name: " + vdcName)

		_, err = c.V1.VDC.New(&v1.CAVVirtualDataCenter{Vdc: v1.CAVVirtualDataCenterVDC{
			Name:                vdcName,
			ServiceClass:        "STD",
			BillingModel:        "PAYG",
			CPUAllocated:        22000,
			VcpuInMhz2:          2200,
			Description:         "vdc created by cloudavenue-cli",
			MemoryAllocated:     30,
			DisponibilityClass:  "ONE-ROOM",
			StorageBillingModel: "PAYG",
			StorageProfiles: []v1.VDCStrorageProfile{
				v1.VDCStrorageProfile{ //nolint
					Class:   "gold",
					Limit:   500,
					Default: true,
				},
			},
		}})

		if err != nil {
			fmt.Println("Error from vdc", err)
			return
		}

		fmt.Println("vdc resource created successfully !")
		fmt.Println("\nvdc resource list after creation:")
		vdcListCmd.Run(cmd, []string{})

	},
}
