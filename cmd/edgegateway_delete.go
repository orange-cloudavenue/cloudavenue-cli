/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	jsontmpl "github.com/orange-cloudavenue/cloudavenue-cli/pkg/templates/json"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var deleteCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		edgeGateways, err := c.V1.EdgeGateway.List()
		if err != nil {
			panic(err)
		}

		jsontmpl.Format(jsontmpl.JsonTemplate{
			Fields: []string{"edgeName", "edgeID", "ownerType", "ownerName", "bandwidth", "description"},
			Data:   edgeGateways,
		})

	},
}

func init() {
	edgegatewayCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
