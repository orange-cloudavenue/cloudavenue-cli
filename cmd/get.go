package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output"
	"github.com/orange-cloudavenue/common-go/print"
	"github.com/spf13/cobra"
)

const (
	cmdGet       = "get"
	cmdGetAlias1 = "ls"
	cmdGetAlias2 = "list"
	description  = `
					The default output format print the minimal necessary information like name, status or group.
					You can use the -o flag to specify the output format.
					"wide" will print some additional information.
					"json" or "yaml" will print the result in the specified format.`
)

var (
	exampleGet1 = `
	#List all T0
	cav get t0`
	exampleGet2 = `
	#List all T0 in wide format
	cav get t0 -o wide`
	exampleGet3 = `
	#List all Public IP
	cav get publicip`
	exampleGet4 = `
	#List all VDC in yaml format
	cav get vdc -o yaml`
	exampleGet5 = `
	#List all S3 in json format
	cav get s3 -o json`
)

// getCmd list a CAV resource
var getCmd = &cobra.Command{
	Use:               cmdGet,
	Aliases:           []string{cmdGetAlias1, cmdGetAlias2},
	Example:           exampleGet1 + "\n" + exampleGet2 + "\n" + exampleGet3 + "\n" + exampleGet4 + "\n" + exampleGet5,
	Short:             "Get resource to retrieve information from CloudAvenue.",
	DisableAutoGenTag: true,
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
	getCmd.PersistentFlags().StringP(flagOutput, "o", "", "Output format. One of: (wide, json, yaml)")
}

// getT0Cmd return a list of your t0 resource(s)
var getT0Cmd = &cobra.Command{
	Use:               argT0,
	Short:             "A brief list of your t0 resources",
	Long:              "A complete list information of your T0 resources in your CloudAvenue account." + description,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		// init Config File
		v, err := initConfig()
		if err != nil {
			log.Default().Println("Error from init config file", err)
			os.Exit(1)
		}

		// init Client
		err = initClient(v)
		if err != nil {
			log.Default().Println("Error from init client:", err)
			os.Exit(1)
		}

		// Check if time flag is set and print time elapsed
		if cmd.Flag(flagTime).Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the list of t0
		t0s, err := c.V1.T0.GetT0s()
		if err != nil {
			log.Default().Println("Error from t0 List", err)
			os.Exit(1)
		}

		// Print the result
		s.Stop()
		flag := cmd.Flag(flagOutput).Value.String()
		w := print.New()
		switch flag {
		case flagOutputValueWide:
			w.SetHeader("name", "t0 provider name", "t0 class service", "services", "class service")
			for _, t0 := range *t0s {
				w.AddFields(t0.Tier0Vrf, t0.Tier0Provider, t0.Tier0ClassService, t0.Services, t0.ClassService)
			}
		case flagOutputValueJSON, flagOutputValueYAML:
			x, err := output.New(stringToTypeFormat(flag), t0s)
			if err != nil {
				log.Default().Println("Error creating output", err)
				os.Exit(1)
			}
			x.ToOutput()
		case "":
			w.SetHeader("name", "t0 provider name")
			for _, t0 := range *t0s {
				w.AddFields(t0.Tier0Vrf, t0.Tier0Provider)
			}
		default:
			log.Default().Println("Error unknow output")
			return
		}
		w.PrintTable()
	},
}

// getPublicIPCmd return a list of your publicip
var getPublicIPCmd = &cobra.Command{
	Use:               argPublicIP,
	Aliases:           []string{argPublicIPAlias1},
	Short:             "A brief list of your public ip resources",
	Long:              "A complete list information of your Public IP resources in your CloudAvenue account." + description,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		// init Config File
		v, err := initConfig()
		if err != nil {
			log.Default().Println("Error from init config file", err)
			os.Exit(1)
		}

		// init Client
		err = initClient(v)
		if err != nil {
			log.Default().Println("Error from init client:", err)
			os.Exit(1)
		}

		// Check if time flag is set and print time elapsed
		if cmd.Flag(flagTime).Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the list of publicip
		ips, err := c.V1.PublicIP.GetIPs()
		if err != nil {
			fmt.Println("Error from IP List", err)
			os.Exit(1)
		}

		// Print the result
		s.Stop()
		flag := cmd.Flag(flagOutput).Value.String()
		w := print.New()
		switch flag {
		case flagOutputValueWide:
			w.SetHeader("public ip", "edge gateway name", "ip natted")
			for _, i := range ips.NetworkConfig {
				w.AddFields(i.UplinkIP, i.EdgeGatewayName, i.TranslatedIP)
			}
		case flagOutputValueJSON, flagOutputValueYAML:
			x, err := output.New(stringToTypeFormat(flag), ips)
			if err != nil {
				log.Default().Println("Error creating output", err)
				os.Exit(1)
			}
			x.ToOutput()
		case "":
			w.SetHeader("public ip", "edge gateway name")
			for _, i := range ips.NetworkConfig {
				w.AddFields(i.UplinkIP, i.EdgeGatewayName)
			}
		default:
			log.Default().Println("Error unknow output")
			return
		}
		w.PrintTable()
	},
}

// getS3Cmd return a list of your s3 (bucket) resource(s)
var getS3Cmd = &cobra.Command{
	Use:               argS3,
	Aliases:           []string{argS3Alias},
	Short:             "A brief list of your s3 resources",
	Long:              "A complete list information of your s3 resources in your CloudAvenue account." + description,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		// init Config File
		v, err := initConfig()
		if err != nil {
			log.Default().Println("Error from init config file", err)
			os.Exit(1)
		}

		// init Client
		err = initClient(v)
		if err != nil {
			log.Default().Println("Error from init client:", err)
			os.Exit(1)
		}

		// Check if time flag is set and print time elapsed
		if cmd.Flag(flagTime).Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the list of buckets
		s3, err := c.V1.S3().ListBuckets(&s3.ListBucketsInput{})
		if err != nil {
			fmt.Println("Error from S3 List", err)
			os.Exit(1)
		}

		// Print the result
		s.Stop()
		flag := cmd.Flag(flagOutput).Value.String()
		w := print.New()
		// var format output.Formatter
		switch flag {
		case flagOutputValueWide:
			w.SetHeader("name", "owner", "creation date", "owner id")
			for _, b := range s3.Buckets {
				w.AddFields(*b.Name, *s3.Owner.DisplayName, *b.CreationDate, *s3.Owner.ID)
			}
		case flagOutputValueJSON, flagOutputValueYAML:
			x, err := output.New(stringToTypeFormat(flag), s3)
			if err != nil {
				log.Default().Println("Error creating output", err)
				os.Exit(1)
			}
			x.ToOutput()
		case "":
			w.SetHeader("name", "owner")
			for _, b := range s3.Buckets {
				w.AddFields(*b.Name, *s3.Owner.DisplayName)
			}
		default:
			log.Default().Println("Error unknow output")
			return
		}
		w.PrintTable()
	},
}

// getEdgeGatewayCmd return a list of your edgegateway resource(s)
var getEdgeGatewayCmd = &cobra.Command{
	Use:               argEdgeGateway,
	Aliases:           []string{argEdgeGatewayAlias1, argEdgeGatewayAlias2},
	Short:             "A brief list of your edgegateway resources",
	Long:              "A complete list information of your EdgeGateway resources in your CloudAvenue account." + description,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		// init Config File
		v, err := initConfig()
		if err != nil {
			log.Default().Println("Error from init config file", err)
			os.Exit(1)
		}

		// init Client
		err = initClient(v)
		if err != nil {
			log.Default().Println("Error from init client:", err)
			os.Exit(1)
		}

		// Check if time flag is set and print time elapsed
		if cmd.Flag(flagTime).Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		edgeGateways, err := c.V1.EdgeGateway.List()
		if err != nil {
			fmt.Println("Error from EdgeGateway", err)
		}

		// Print the result
		s.Stop()
		flag := cmd.Flag(flagOutput).Value.String()
		w := print.New()
		switch flag {
		case flagOutputValueWide:
			w.SetHeader("name", "id", "owner name", "owner type", "ratelimit (mb/s)", "description", "tier0 vrf name")
			for _, e := range *edgeGateways {
				w.AddFields(e.EdgeName, e.EdgeID, e.OwnerName, e.OwnerType, e.Bandwidth, e.Description, e.Tier0VrfName)
			}
		case flagOutputValueJSON, flagOutputValueYAML:
			x, err := output.New(stringToTypeFormat(flag), edgeGateways)
			if err != nil {
				log.Default().Println("Error creating output", err)
				os.Exit(1)
			}
			x.ToOutput()
		case "":
			w.SetHeader("name", "owner")
			for _, e := range *edgeGateways {
				w.AddFields(e.EdgeName, e.OwnerName)
			}
		default:
			log.Default().Println("Error unknow output")
			return
		}
		w.PrintTable()
	},
}

// getVDCCmd return a list of your vdc resource(s)
var getVDCCmd = &cobra.Command{
	Use:               argVDC,
	Short:             "A brief list of your vdc resources",
	Long:              "A complete list information of your s3 resources in your CloudAvenue account." + description,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		// init Config File
		v, err := initConfig()
		if err != nil {
			log.Default().Println("Error from init config file", err)
			os.Exit(1)
		}

		// init Client
		err = initClient(v)
		if err != nil {
			log.Default().Println("Error from init client:", err)
			os.Exit(1)
		}

		// Check if time flag is set and print time elapsed
		if cmd.Flag(flagTime).Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the list of vdc
		vdcs, err := c.V1.Querier().List().VDC()
		if err != nil {
			fmt.Println("Error from VDC List", err)
			os.Exit(1)
		}

		// Print the result
		s.Stop()
		flag := cmd.Flag(flagOutput).Value.String()
		w := print.New()
		switch flag {
		case flagOutputValueWide:
			w.SetHeader("name", "status", "cpu used (mhz)", "memory used (mb)", "storage used (mb)", "number of vm(s)", "number of vapp(s)")
			for _, v := range vdcs {
				w.AddFields(v.Name, v.Status, *v.CpuUsedMhz, *v.MemoryUsedMB, *v.StorageUsedMB, *v.NumberOfVMs, *v.NumberOfVApps)
			}
		case flagOutputValueJSON, flagOutputValueYAML:
			x, err := output.New(stringToTypeFormat(flag), vdcs)
			if err != nil {
				log.Default().Println("Error creating output", err)
				os.Exit(1)
			}
			x.ToOutput()
		case "":
			w.SetHeader("name", "status")
			for _, v := range vdcs {
				w.AddFields(v.Name, v.Status)
			}
		default:
			log.Default().Println("Error unknow output")
			return
		}
		w.PrintTable()
	},
}
