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

// publicipCmd represents the vdc command
var publicipCmd = &cobra.Command{
	Use:     "publicip",
	Example: "publicip <list | create | delete>",
	Short:   "Option to manage your public ip on CloudAvenue.",
}

func init() {
	rootCmd.AddCommand(publicipCmd)

	// ? List command
	publicipCmd.Args = cobra.NoArgs
	publicipCmd.AddCommand(publicipListCmd)

	// ? Delete command
	publicipCmd.AddCommand(publicipDelCmd)
	publicipDelCmd.Args = cobra.MinimumNArgs(1)

	// ? Create command
	publicipCmd.AddCommand(publicipCreateCmd)
	publicipCreateCmd.PersistentFlags().String("name", "", "vdc name")
	if err := publicipCreateCmd.MarkPersistentFlagRequired("name"); err != nil {
		log.Default().Println("Error from Flag name, is require.", err)
		return
	}
}

// listCmd represents the list command
var publicipListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief list of your publicip resources",
	Run: func(cmd *cobra.Command, args []string) {

		// Get the list of vdc
		ips, err := c.V1.PublicIP.GetIPs()
		if err != nil {
			log.Default().Println("Error from IP List", err)
			return
		}

		// Struct to print a basic view
		type basicIP = struct {
			IP              string `json:"ip"`
			IPnat           string `json:"ip_nat"`
			EdgeGatewayName string `json:"edge_gateway_name"`
		}
		basicIPs := []*basicIP{}

		// Set the struct
		for _, ip := range ips.NetworkConfig {
			x := &basicIP{
				IP:              ip.UplinkIP,
				IPnat:           ip.TranslatedIP,
				EdgeGatewayName: ip.EdgeGatewayName,
			}
			basicIPs = append(basicIPs, x)
		}

		// Print the result
		jsontmpl.Format(jsontmpl.JsonTemplate{
			Fields: []string{"ip", "ip_nat", "edge_gateway_name"},
			Data:   basicIPs,
		})
	},
}

// deleteCmd represents the delete command
var publicipDelCmd = &cobra.Command{
	Use:     "delete",
	Example: "publicip delete <name> [<name>] [<name>] ...",
	Short:   "Delete publicip resource(s)",

	Run: func(cmd *cobra.Command, args []string) {

		for _, arg := range args {
			fmt.Println("delete publicip resource " + arg)
			ip, err := c.V1.PublicIP.GetIP(arg)
			if err != nil {
				log.Default().Println("Error from ip", err)
				return
			}
			job, err := ip.Delete()
			if err != nil {
				log.Default().Println("Unable to delete ip", err)
				return
			}
			err = job.Wait(15, 300)
			if err != nil {
				log.Default().Println("Error during ip Deletion !!", err)
				return
			}
			fmt.Println("ip resource deleted " + arg + " successfully !!\n")
			fmt.Println("ip resource list after deletion:")
			publicipListCmd.Run(cmd, []string{})
		}

	},
}

// createCmd represents the create command
var publicipCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an ip",
	Example: "ip create --name <edgegateway name>",

	Run: func(cmd *cobra.Command, args []string) {

		// Get the name from the command line
		gwName, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Default().Println("Malformed EdgeGateway Name ", err)
			return
		}

		// // Get EdgeGateway from name
		// gw, err := c.V1.EdgeGateway.GetByName(gwName)
		// if err != nil {
		// 	log.Default().Println("EdgeGateway not found", err)
		// 	return
		// }

		// Create the vdc
		fmt.Println("create public ip resource")
		fmt.Println("for EdgeGateway name: " + gwName)

		// c.V1.EdgeGateway.

		// _, err = c.V1.VDC.New(&v1.CAVVirtualDataCenter{Vdc: v1.CAVVirtualDataCenterVDC{
		// 	Name:                vdcName,
		// 	ServiceClass:        "STD",
		// 	BillingModel:        "PAYG",
		// 	CPUAllocated:        22000,
		// 	VcpuInMhz2:          2200,
		// 	Description:         "vdc created by cloudavenue-cli",
		// 	MemoryAllocated:     30,
		// 	DisponibilityClass:  "ONE-ROOM",
		// 	StorageBillingModel: "PAYG",
		// 	StorageProfiles: []v1.VDCStrorageProfile{
		// 		v1.VDCStrorageProfile{ //nolint
		// 			Class:   "gold",
		// 			Limit:   500,
		// 			Default: true,
		// 		},
		// 	},
		// }})

		if err != nil {
			log.Default().Println("Error from vdc", err)
			return
		}

		fmt.Println("vdc resource created successfully !")
		fmt.Println("\nvdc resource list after creation:")
		publicipListCmd.Run(cmd, []string{})

	},
}
