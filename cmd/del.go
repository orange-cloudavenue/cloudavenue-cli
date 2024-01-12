package cmd

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go/pkg/uuid"
	v1 "github.com/orange-cloudavenue/cloudavenue-sdk-go/v1"
	"github.com/spf13/cobra"
)

var (
	exdel1 = "#Delete a Public IP\ncav del ip 192.168.0.2\n\n"
	exdel2 = "#Delete several vdc named xxxx and yyyy\ncav del vdc --name xxxx yyyy\n\n"
)

// getCmd represents the t0 command
var delCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm"},
	Example: exdel1 + exdel2,
	Short:   "Delete resource from CloudAvenue.",
}

func init() {
	rootCmd.AddCommand(delCmd)

	// ? Delete command
	delCmd.Args = cobra.NoArgs

	// ? Delete for public ip
	delCmd.AddCommand(delPublicIPCmd)
	delPublicIPCmd.Args = cobra.MinimumNArgs(1)

	// ? Delete for s3
	delCmd.AddCommand(delS3Cmd)
	delS3Cmd.Args = cobra.MinimumNArgs(1)

	// ? Delete for edgegateway
	delCmd.AddCommand(delEdgeGatewayCmd)
	delEdgeGatewayCmd.Args = cobra.MinimumNArgs(1)

	// ? Delete for vdc
	delCmd.AddCommand(delVDCCmd)
	delVDCCmd.Args = cobra.MinimumNArgs(1)

}

// delVDCCmd represents the delete command
var delVDCCmd = &cobra.Command{
	Use:     "vdc",
	Example: "del vdc <name> [<name>] [<name>] ...",
	Short:   "Delete a vdc",

	Run: func(cmd *cobra.Command, args []string) {
		// Check if time flag is set
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		for _, arg := range args {
			fmt.Println("delete vdc resource " + arg)
			// vdc, err := c.V1.VDC.Get(arg)
			vdc, err := c.V1.VDC().GetVDC(arg)
			if err != nil {
				fmt.Println("Error from vdc", err)
				return
			}
			job, err := vdc.Delete()
			if err != nil {
				fmt.Println("Unable to delete vdc", err)
				return
			}
			err = job.Wait(3, 300)
			if err != nil {
				fmt.Println("Error during vdc Deletion !!", err)
				return
			}
			fmt.Println("vdc resource deleted " + arg + " successfully !!")
		}

	},
}

// deleteCmd represents the delete command
var delS3Cmd = &cobra.Command{
	Use:     "s3",
	Example: "delete s3 <name> [<name>] [<name>] ...",
	Short:   "Delete a s3 bucket",

	Run: func(cmd *cobra.Command, args []string) {
		// Check if time flag is set
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// init client
		s3Client := c.V1.S3()

		for i, arg := range args {
			fmt.Println("delete bucket resource " + arg)
			// Del the bucket
			_, err := s3Client.DeleteBucket(&s3.DeleteBucketInput{Bucket: &args[i]})
			if err != nil {
				fmt.Println("Error from S3 Delete", err)
				return
			}
			fmt.Println("Bucket resource deleted " + arg + " successfully !!")
			fmt.Println("\nBucket resource list after deletion:")
		}

	},
}

// deleteCmd represents the delete command
var delEdgeGatewayCmd = &cobra.Command{
	Use:     "edgegateway",
	Aliases: []string{"egw", "gw"},
	Example: "delete edgegateway <id or name> [<id or name>] [<id or name>] ...",
	Short:   "Delete an edgeGateway (name or id)",

	Run: func(cmd *cobra.Command, args []string) {
		// Check if time flag is set
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}
		var (
			gw  *v1.EdgeGw
			err error
		)

		for _, arg := range args {
			fmt.Println("delete EdgeGateway resource " + arg)
			if uuid.IsUUIDV4(arg) {
				gw, err = c.V1.EdgeGateway.GetByID(arg)
			} else {
				gw, err = c.V1.EdgeGateway.GetByName(arg)
			}
			if err != nil {
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
		}

	},
}

// deleteCmd represents the delete command
var delPublicIPCmd = &cobra.Command{
	Use:     "publicip",
	Example: "delete publicip <ip> [<ip>] [<ip>] ...",
	Short:   "Delete public ip resource(s)",

	Run: func(cmd *cobra.Command, args []string) {
		// Check if time flag is set
		if cmd.Flag("time").Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

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
		}

	},
}
