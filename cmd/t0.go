/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	jsontmpl "github.com/orange-cloudavenue/cloudavenue-cli/pkg/templates/json"
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

}

// listCmd represents the list command
var t0ListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief list of your t0 resources",
	Run: func(cmd *cobra.Command, args []string) {

		// Get the list of t0
		t0s, err := c.V1.T0.GetT0s()
		if err != nil {
			log.Default().Println("Error from t0 List", err)
			return
		}

		// Struct to print a basic view
		type basict0 = struct {
			T0Vrf          string `json:"t0_vrf"`
			T0Provider     string `json:"t0_provider"`
			T0ClassService string `json:"t0_class_service"`
		}
		basict0s := []*basict0{}

		// Set the struct
		for _, t0 := range *t0s {
			x := &basict0{
				T0Vrf:          t0.Tier0Vrf,
				T0Provider:     t0.Tier0Provider,
				T0ClassService: t0.Tier0ClassService,
			}
			basict0s = append(basict0s, x)
		}

		// Print the result
		jsontmpl.Format(jsontmpl.JsonTemplate{
			Fields: []string{"t0_vrf", "t0_provider", "t0_class_service"},
			Data:   basict0s,
		})
	},
}
