/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/print"
	"github.com/spf13/cobra"
)

const CmdPublicIP = "publicip"

// publicipCmd represents the vdc command
var publicipCmd = &cobra.Command{
	Use:     CmdPublicIP,
	Example: CmdPublicIP + " <list | create | delete>",
	Short:   "Option to manage your public ip on CloudAvenue.",
}

func init() {
	rootCmd.AddCommand(publicipCmd)

	// ? List command
	publicipCmd.Args = cobra.NoArgs
	publicipCmd.AddCommand(publicipListCmd)
	publicipCmd.PersistentFlags().StringP("output", "o", "", "Print all resources informations")

	// ? Delete command
	publicipCmd.AddCommand(publicipDelCmd)
	publicipDelCmd.Args = cobra.MinimumNArgs(1)

	// ? Create command
	publicipCmd.AddCommand(publicipCreateCmd)
	publicipCreateCmd.PersistentFlags().String("name", "", "edge gateway name")
	if err := publicipCreateCmd.MarkPersistentFlagRequired("name"); err != nil {
		fmt.Println("Error from Flag name, is require.", err)
		return
	}
}

// listCmd represents the list command
var publicipListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "get"},
	Short:   "A brief list of your publicip resources",
	Long:    "A complete list information of your Public IP resources in your CloudAvenue account.",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the list of vdc
		ips, err := c.V1.PublicIP.GetIPs()
		if err != nil {
			fmt.Println("Error from IP List", err)
			return
		}

		// Print the result
		flag := cmd.Flag("output").Value
		w := print.New()
		switch flag.String() {
		case "wide":
			w.SetHeader("public ip", "edge gateway name", "ip natted")
			for _, i := range ips.NetworkConfig {
				w.AddFields(i.UplinkIP, i.EdgeGatewayName, i.TranslatedIP)
			}
		default:
			w.SetHeader("public ip", "edge gateway name")
			for _, i := range ips.NetworkConfig {
				w.AddFields(i.UplinkIP, i.EdgeGatewayName)
			}
		}
		w.PrintTable()
	},
}

// deleteCmd represents the delete command
var publicipDelCmd = &cobra.Command{
	Use:     "delete",
	Example: "publicip delete <ip> [<ip>] [<ip>] ...",
	Short:   "Delete publicip resource(s)",

	Run: func(cmd *cobra.Command, args []string) {
		defer timeTrack(time.Now(), cmd.CommandPath())

		for _, arg := range args {
			fmt.Println("delete publicip resource " + arg)
			ip, err := c.V1.PublicIP.GetIP(arg)
			if err != nil {
				fmt.Println("Error from ip: ", err)
				return
			}
			job, err := ip.Delete()
			if err != nil {
				fmt.Println("Unable to delete ip: ", err)
				return
			}
			err = job.Wait(15, 300)
			if err != nil {
				fmt.Println("Error during ip Deletion !!", err)
				return
			}
			fmt.Println("ip resource deleted " + arg + " successfully !!")
			fmt.Println("\nip resource list after deletion:")
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
		defer timeTrack(time.Now(), cmd.CommandPath())

		// Get the name from the command line
		gwName, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Malformed argument EdgeGateway Name ", err)
			return
		}

		// Create a public ip
		fmt.Println("create public ip resource")
		fmt.Println("for EdgeGateway name: " + gwName)

		job, err := c.V1.PublicIP.New(gwName)
		if err != nil {
			fmt.Println("Unable to create public ip", err)
			return
		}
		err = job.Wait(5, 300)
		if err != nil {
			fmt.Println("Error during public ip creation !!", err)
			return
		}
		fmt.Println("public ip resource created successfully !")
		fmt.Println("\npublic ip resource list after creation:")
		publicipListCmd.Run(cmd, []string{})

	},
}
