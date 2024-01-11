package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/print"
	"github.com/spf13/cobra"
)

var (
	exget1 = "#List all T0\ncav get t0\n\n"
	exget2 = "#List all T0 in wide format\ncav get t0 -o wide\n\n"
)

// getCmd represents the t0 command
var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"ls", "list"},
	Example: exget1 + exget2,
	Short:   "Get resource to retrieve information from CloudAvenue.",
}

func init() {
	rootCmd.AddCommand(getCmd)

	// ? Get command
	getCmd.Args = cobra.NoArgs
	getCmd.AddCommand(getT0Cmd)
	getCmd.AddCommand(getPublicIPCmd)
	getCmd.AddCommand(getS3Cmd)
	getCmd.AddCommand(getEdgeGatewayCmd)
	getCmd.AddCommand(getVDCCmd)
	getCmd.PersistentFlags().StringP("output", "o", "", "Output format. One of: (wide, json, yaml)")

}

// getT0Cmd represents the Get of T0 command
var getT0Cmd = &cobra.Command{
	Use: "t0",
	// Aliases: []string{"t0s"},
	Short: "A brief list of your t0 resources",
	Long:  "A complete list information of your T0 resources in your CloudAvenue account.",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if time flag is set
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

// listCmd represents the list command
var getPublicIPCmd = &cobra.Command{
	Use:     "publicip",
	Aliases: []string{"ip"},
	Short:   "A brief list of your public ip resources",
	Long:    "A complete list information of your Public IP resources in your CloudAvenue account.",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if time flag is set
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the list of publicip
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

// getS3Cmd represents the list command
var getS3Cmd = &cobra.Command{
	Use:     "s3",
	Aliases: []string{"ls", "get"},
	Short:   "A brief list of your s3 resources",
	Long:    "A complete list information of your s3 resources in your CloudAvenue account.",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if time flag is set
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// init client
		s3Client := c.V1.S3()

		// Get the list of buckets
		output, err := s3Client.ListBuckets(&s3.ListBucketsInput{})
		if err != nil {
			fmt.Println("Error from S3 List", err)
			return
		}

		// Print the result
		flag := cmd.Flag("output").Value
		w := print.New()
		switch flag.String() {
		case "wide":
			w.SetHeader("name", "owner", "creation date", "owner id")
			for _, b := range output.Buckets {
				w.AddFields(*b.Name, *output.Owner.DisplayName, *b.CreationDate, *output.Owner.ID)
			}
		default:
			w.SetHeader("name", "owner")
			for _, b := range output.Buckets {
				w.AddFields(*b.Name, *output.Owner.DisplayName)
			}
		}
		w.PrintTable()
	},
}

// getEdgeGatewayCmd represents the list command
var getEdgeGatewayCmd = &cobra.Command{
	Use:     "edgegateway",
	Aliases: []string{"gw", "egw"},
	Short:   "A brief list of your edgegateway resources",
	Long:    "A conplete list information of your EdgeGateway resources in your CloudAvenue account.",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		edgeGateways, err := c.V1.EdgeGateway.List()
		if err != nil {
			fmt.Println("Error from EdgeGateway", err)
		}

		// Print the result
		flag := cmd.Flag("output").Value
		w := print.New()
		switch flag.String() {
		case "wide":
			w.SetHeader("name", "id", "owner name", "owner type", "ratelimit (mb/s)", "description", "tier0 vrf name")
			for _, e := range *edgeGateways {
				w.AddFields(e.EdgeName, e.EdgeID, e.OwnerName, e.OwnerType, e.Bandwidth, e.Description, e.Tier0VrfName)
			}
		default:
			w.SetHeader("name", "owner")
			for _, e := range *edgeGateways {
				w.AddFields(e.EdgeName, e.OwnerName)
			}
		}
		w.PrintTable()
	},
}

// listCmd represents the list command
var getVDCCmd = &cobra.Command{
	Use: "vdc",
	// Aliases: []string{"ls", "get"},
	Short: "A brief list of your vdc resources",
	Long:  "A complete list information of your s3 resources in your CloudAvenue account.",
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

		flag := cmd.Flag("output").Value.String()
		w := print.New()
		switch flag {
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
