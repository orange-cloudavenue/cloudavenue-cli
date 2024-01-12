package cmd

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go/v1/infrapi"
	"github.com/spf13/cobra"
)

var (
	exampleCreate1 = `
	#List all T0
	cav create vdc --name myvdc`
)

// getCmd represents the t0 command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "add"},
	Example: exampleCreate1,
	Short:   "Create resource to CloudAvenue.",
}

func init() {
	rootCmd.AddCommand(createCmd)

	// ? Create command
	createCmd.Args = cobra.NoArgs
	// ? Create for public ip
	createCmd.AddCommand(createPublicIPCmd)
	// ? Create for s3
	createCmd.AddCommand(createS3Cmd)
	// ? Create for edgegateway
	createCmd.AddCommand(createEdgeGatewayCmd)
	// ? Create for vdc
	createCmd.AddCommand(createVDCCmd)

	// ? Options for edgegateway
	createEdgeGatewayCmd.Flags().String("t0", "", "t0 name")
	createEdgeGatewayCmd.Flags().String("vdc", "", "vdc name")
	if err := createEdgeGatewayCmd.MarkFlagRequired("vdc"); err != nil {
		fmt.Println("Error from Flag VDC, is require.", err)
		return
	}

	// ? Options for publicip
	createPublicIPCmd.Flags().String("name", "", "public ip address")
	if err := createPublicIPCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println("Error from Flag name, is require.", err)
		return
	}

	// ? Options for s3
	createS3Cmd.Flags().String("name", "", "s3 bucket name")
	if err := createS3Cmd.MarkFlagRequired("name"); err != nil {
		fmt.Println("Error from Flag name, is require.", err)
		return
	}

	// ? Options for vdc
	createVDCCmd.Flags().String("name", "", "vdc name")
	if err := createVDCCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println("Error from Flag Name, is require.", err)
		return
	}
}

// createCmd represents the create command
var createPublicIPCmd = &cobra.Command{
	Use:     "publicip",
	Short:   "Create an ip",
	Example: "ip create --name <>",

	Run: func(cmd *cobra.Command, args []string) {
		// Check if time flag is set
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

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
	},
}

// createCmd represents the create command
var createVDCCmd = &cobra.Command{
	Use:     "vdc",
	Short:   "Create an vdc",
	Example: "vdc create --name <vdc name>",

	Run: func(cmd *cobra.Command, args []string) {
		// Check if time flag is set
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the vdc name from the command line
		vdcName, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Malformed VDC name", err)
			return
		}

		// Create the vdc
		fmt.Println("create vdc resource (with basic value)")
		fmt.Println("vdc name: " + vdcName)

		_, err = c.V1.VDC().New(&infrapi.CAVVirtualDataCenter{VDC: infrapi.CAVVirtualDataCenterVDC{
			Name:                vdcName,
			ServiceClass:        "STD",
			BillingModel:        "PAYG",
			CPUAllocated:        22000,
			VCPUInMhz:           2200,
			Description:         "vdc created by cloudavenue-cli",
			MemoryAllocated:     30,
			DisponibilityClass:  "ONE-ROOM",
			StorageBillingModel: "PAYG",
			StorageProfiles: []infrapi.StorageProfile{
				{
					Class:   "gold",
					Limit:   500,
					Default: true,
				},
			},
		},
		})

		if err != nil {
			fmt.Println("Error from vdc", err)
			return
		}

		fmt.Println("vdc resource created successfully !")
	},
}

// createCmd represents the create command
var createEdgeGatewayCmd = &cobra.Command{
	Use:     "edgegateway",
	Short:   "Create an edgeGateway",
	Aliases: []string{"gw", "egw"},
	Example: "edgegateway create --vdc <vdc name> [--t0 <t0 name>]",

	Run: func(cmd *cobra.Command, args []string) {
		// Check if time flag is set
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the vdc name from the command line
		vdc, err := cmd.Flags().GetString("vdc")
		if err != nil {
			fmt.Println("Error from VDC", err)
			return
		}

		// Get the t0 name
		t0s, err := c.V1.T0.GetT0s()
		if err != nil {
			fmt.Println("Error from T0 List", err)
			return
		}
		if len(*t0s) == 0 {
			fmt.Println("No T0 found, please create one before")
			return
		}
		if len(*t0s) > 1 && cmd.Flag("t0").Value.String() == "" {
			fmt.Println("More than one T0 found, please specify one")
			return
		}
		var t0 string
		if len(*t0s) == 1 {
			t0 = (*t0s)[0].Tier0Vrf
		} else {
			t0, err = cmd.Flags().GetString("t0")
			if err != nil {
				fmt.Println("Error to retrieve T0", err)
				return
			}
		}

		// Create the edgeGateway
		fmt.Println("Creating EdgeGateway resource")
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
	},
}

// createCmd represents the create command
var createS3Cmd = &cobra.Command{
	Use:     "s3",
	Short:   "Create an S3 bucket",
	Example: "create s3 --name <bucket name>",

	Run: func(cmd *cobra.Command, args []string) {
		// Check if time flag is set
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		bucketName, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Malformed bucket name ", err)
			return
		}

		// Create the bucket
		fmt.Println("create a bucket resource (with basic value)")
		fmt.Println("bucket name: " + bucketName)

		_, err = c.V1.S3().CreateBucket(&s3.CreateBucketInput{Bucket: &bucketName})
		if err != nil {
			fmt.Println("Error from S3 Create", err)
			return
		}

		fmt.Println("Bucket resource created successfully !")
	},
}
