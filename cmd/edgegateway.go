/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	jsontmpl "github.com/orange-cloudavenue/cloudavenue-cli/pkg/templates/json"
	"github.com/spf13/cobra"
)

const CmdEdgeGateway = "edgegateway"

// edgegatewayCmd represents the edgegateway command
var gwCmd = &cobra.Command{
	Use:     CmdEdgeGateway,
	Example: CmdEdgeGateway + " <list | create | delete>",
	Short:   "Option to manage your edgeGateway NSX on CloudAvenue.",
}

func init() {
	rootCmd.AddCommand(gwCmd)

	// ? List command
	gwCmd.Args = cobra.NoArgs
	gwCmd.AddCommand(gwListCmd)

	// ? Delete command
	gwCmd.AddCommand(gwDelCmd)
	gwDelCmd.Args = cobra.MinimumNArgs(1)

	// ? Create command
	gwCmd.AddCommand(gwCreateCmd)
	gwCreateCmd.PersistentFlags().String("vdc", "", "vdc name")
	gwCreateCmd.PersistentFlags().String("t0", "", "t0 name")
	if err := gwCreateCmd.MarkPersistentFlagRequired("vdc"); err != nil {
		fmt.Println("Error from Flag VDC, is require.", err)
		return
	}
	if err := gwCreateCmd.MarkPersistentFlagRequired("t0"); err != nil {
		fmt.Println("Error from Flag T0, is require.", err)
		return
	}
}

// listCmd represents the list command
var gwListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief list of your edgegateway resources",
	Run: func(cmd *cobra.Command, args []string) {
		defer timeTrack(time.Now(), cmd.CommandPath())

		edgeGateways, err := c.V1.EdgeGateway.List()
		if err != nil {
			fmt.Println("Error from EdgeGateway", err)
		}

		jsontmpl.Format(jsontmpl.JsonTemplate{
			Fields: []string{"edgeName", "edgeId", "ownerType", "ownerName", "rateLimit", "description"},
			Data:   edgeGateways,
		})

	},
}

// deleteCmd represents the delete command
var gwDelCmd = &cobra.Command{
	Use:     "delete",
	Example: "edgegateway delete <id or name> [<id or name>] [<id or name>] ...",
	Short:   "Delete an edgeGateway (name or id)",

	Run: func(cmd *cobra.Command, args []string) {
		defer timeTrack(time.Now(), cmd.CommandPath())

		for _, arg := range args {
			fmt.Println("delete EdgeGateway resource " + arg)
			gw, err := c.V1.EdgeGateway.GetByName(arg)
			if err != nil {
				gw, err = c.V1.EdgeGateway.GetByID(arg)
				if err != nil {
					fmt.Println("Unable to find EdgeGateway ID or Name", err)
					return
				}
			}
			job, err := gw.Delete()
			if err != nil {
				fmt.Println("Unable to delete EdgeGateway", err)
				return
			}
			err = job.Wait(3, 300)
			if err != nil {
				fmt.Println("Error during EdgeGateway Deletion !!", err)
				return
			}
			fmt.Println("EdgeGateway resource deleted " + arg + " successfully !!")
			fmt.Println("\nEdgeGateway resource list after deletion:")
			gwListCmd.Run(cmd, []string{})
		}

	},
}

// createCmd represents the create command
var gwCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an edgeGateway",
	Example: "edgegateway create --vdc <vdc name> --t0 <t0 name>",

	Run: func(cmd *cobra.Command, args []string) {
		defer timeTrack(time.Now(), cmd.CommandPath())

		// Get the vdc name from the command line
		vdc, err := cmd.Flags().GetString("vdc")
		if err != nil {
			fmt.Println("Error from VDC", err)
			return
		}

		// Get the t0 name from the command line
		t0, err := cmd.Flags().GetString("t0")
		if err != nil {
			fmt.Println("Error from T0", err)
			return
		}

		// Create the edgeGateway
		fmt.Println("create EdgeGateway resource")
		fmt.Println("vdc name: " + vdc)
		fmt.Println("t0 name: " + t0)
		job, err := c.V1.EdgeGateway.New(vdc, t0)
		if err != nil {
			fmt.Println("Error from EdgeGateway", err)
			return
		}
		err = job.Wait(3, 300)
		if err != nil {
			fmt.Println("Error during EdgeGateway Creation !!", err)
			return
		}
		fmt.Println("EdgeGateway resource created successfully !")
		fmt.Println("\nEdgeGateway resource list after creation:")
		gwListCmd.Run(cmd, []string{})

	},
}
