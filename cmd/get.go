/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	jsontmpl "github.com/orange-cloudavenue/cloudavenue-cli/pkg/templates/json"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		edgeGateway, err := c.V1.EdgeGateway.GetByName("tn01e02ocb0006205spt101")
		if err != nil {
			panic(err)
		}

		jsontmpl.Format2(jsontmpl.JsonTemplate{
			Fields: []string{"edgeName", "edgeId", "ownerType", "ownerName", "rateLimit", "description"},
			Data:   edgeGateway,
		})
	},
}

func init() {
	edgegatewayCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
