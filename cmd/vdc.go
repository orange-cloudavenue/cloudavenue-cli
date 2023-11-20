/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	jsontmpl "github.com/orange-cloudavenue/cloudavenue-cli/pkg/templates/json"
	"github.com/spf13/cobra"
)

// vdcCmd represents the vdc command
var vdcCmd = &cobra.Command{
	Use:     "vdc",
	Example: "vdc <list | create | delete>",
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

	// TODO : Create command
	// ? Create command
	vdcCmd.AddCommand(vdcCreateCmd)
	vdcCreateCmd.PersistentFlags().String("vdc", "", "vdc name")
	vdcCreateCmd.PersistentFlags().String("t0", "", "t0 name")
	if err := vdcCreateCmd.MarkPersistentFlagRequired("vdc"); err != nil {
		log.Default().Println("Error from Flag VDC, is require.", err)
		return
	}
	if err := vdcCreateCmd.MarkPersistentFlagRequired("t0"); err != nil {
		log.Default().Println("Error from Flag T0, is require.", err)
		return
	}
}

// listCmd represents the list command
var vdcListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief list of your vdc resources",
	Run: func(cmd *cobra.Command, args []string) {

		vdcs, err := c.V1.VDC.List()
		if err != nil {
			log.Default().Println("Error from VDC List", err)
			return
		}

		type printVdc = struct {
			VdcGroup        string `json:"vdc_group"`
			VdcName         string `json:"vdc_name"`
			VcpuInMhz2      int    `json:"vcpu_in_mhz2"`
			MemoryAllocated int    `json:"memory_allocated"`
			CPUAllocated    int    `json:"cpu_allocated"`
			Description     string `json:"description"`
		}
		dimensionVdc := []*printVdc{}

		for _, dc := range *vdcs {

			x := &printVdc{
				VdcGroup:        dc.VdcGroup,
				VdcName:         dc.Vdc.Name,
				VcpuInMhz2:      dc.Vdc.VcpuInMhz2,
				MemoryAllocated: dc.Vdc.MemoryAllocated,
				CPUAllocated:    dc.Vdc.CPUAllocated,
				Description:     dc.Vdc.Description,
			}
			dimensionVdc = append(dimensionVdc, x)
		}

		jsontmpl.Format(jsontmpl.JsonTemplate{
			Fields: []string{"vdc_name", "vdc_group", "vcpu_in_mhz2", "memory_allocated", "cpu_allocated", "description"},
			Data:   dimensionVdc,
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
				log.Default().Println("Error from vdc", err)
				return
			}
			job, err := vdc.Delete()
			if err != nil {
				log.Default().Println("Unable to delete vdc", err)
			}
			err = job.Wait(15, 300)
			if err != nil {
				log.Default().Println("Error during vdc Deletion !!", err)
			}
			fmt.Println("vdc resource deleted " + arg + " successfully !!\n")
			fmt.Println("vdc resource list after deletion:")
			vdcListCmd.Run(cmd, []string{})
		}

	},
}

// TODO : Create command
// createCmd represents the create command
var vdcCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "Ceate an vdc",
	Example: "vdc create --vdc <vdc name> --t0 <t0 name>",

	Run: func(cmd *cobra.Command, args []string) {

		// Get the vdc name from the command line
		vdc, err := cmd.Flags().GetString("vdc")
		if err != nil {
			log.Default().Println("Error from VDC", err)
			return
		}

		// Get the t0 name from the command line
		t0, err := cmd.Flags().GetString("t0")
		if err != nil {
			log.Default().Println("Error from T0", err)
			return
		}

		// Create the vdc
		fmt.Println("create vdc resource")
		fmt.Println("vdc name: " + vdc)
		fmt.Println("t0 name: " + t0)
		// job, err := c.V1.VDC.New(&v1.CAVVirtualDataCenter{Vdc: v1.CAVVirtualDataCenterVDC{
		// 	Name:                vdc,
		// 	ServiceClass:        "STD",
		// 	BillingModel:        "PAYG",
		// 	CPUAllocated:        22000,
		// 	VcpuInMhz2:          2200,
		// 	Description:         "vdc created by cloudavenue-cli",
		// 	MemoryAllocated:     30,
		// 	DisponibilityClass:  "ONE-ROOM",
		// 	StorageBillingModel: "PAYG",
		// 	StorageProfiles: []v1.VDCStrorageProfile{
		// 		Class:   ,
		// 		Limit:   500,
		// 		Default: true},
		// }})
		if err != nil {
			log.Default().Println("Error from vdc", err)
			return
		}
		// err = job.Wait(15, 300)
		// if err != nil {
		// 	log.Default().Println("Error during vdc Creation !!", err)
		// 	return
		// }
		fmt.Println("vdc resource created successfully !")
		fmt.Println("\nvdc resource list after creation:")
		vdcListCmd.Run(cmd, []string{})

	},
}
