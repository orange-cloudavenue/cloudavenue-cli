/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"time"

	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/print"
	"github.com/spf13/cobra"
)

// t0Cmd represents the t0 command
var t0Cmd = &cobra.Command{
	Use:     "t0",
	Example: "t0 list",
	Short:   "Option to list your t0 (provider gateway) on CloudAvenue.",
}

func init() {
	rootCmd.AddCommand(t0Cmd)

	// ? List command
	t0Cmd.Args = cobra.NoArgs
	t0Cmd.AddCommand(t0ListCmd)
	t0Cmd.PersistentFlags().StringP("output", "o", "", "Print all resources informations")

}

// listCmd represents the list command
var t0ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "get"},
	Short:   "A brief list of your t0 resources",
	Long:    "A complete list information of your T0 resources in your CloudAvenue account.",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the list of t0
		t0s, err := c.V1.T0.GetT0s()
		if err != nil {
			log.Default().Println("Error from t0 List", err)
			return
		}

		// Print the result
		flag := cmd.Flag("output").Value
		w := print.New()
		switch flag.String() {
		case "wide":
			w.SetHeader("name", "t0 provider name", "t0 class service", "services", "class service")
			for _, t0 := range *t0s {
				w.AddFields(t0.Tier0Vrf, t0.Tier0Provider, t0.Tier0ClassService, t0.Services, t0.ClassService)
			}
		default:
			w.SetHeader("name", "t0 provider name")
			for _, t0 := range *t0s {
				w.AddFields(t0.Tier0Vrf, t0.Tier0Provider)
			}
		}
		w.PrintTable()
	},
}
