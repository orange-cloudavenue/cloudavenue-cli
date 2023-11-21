/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
	jsontmpl "github.com/orange-cloudavenue/cloudavenue-cli/pkg/templates/json"
	"github.com/spf13/cobra"
)

// s3Cmd represents the vdc command
var s3Cmd = &cobra.Command{
	Use:     "s3",
	Example: "s3 <list | create | delete>",
	Short:   "Option to manage your s3 (Object Storage) on CloudAvenue.",
}

func init() {

	rootCmd.AddCommand(s3Cmd)

	// ? List command
	s3Cmd.Args = cobra.NoArgs
	s3Cmd.AddCommand(s3ListCmd)

	// ? Delete command
	s3Cmd.AddCommand(s3DelCmd)
	s3DelCmd.Args = cobra.MinimumNArgs(1)

	// ? Create command
	s3Cmd.AddCommand(s3CreateCmd)
	s3CreateCmd.PersistentFlags().String("name", "", "s3 bucket name")
	if err := s3CreateCmd.MarkPersistentFlagRequired("name"); err != nil {
		fmt.Println("Error from Flag name, is require.", err)
		return
	}
}

// listCmd represents the list command
var s3ListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief list of your s3 resources",
	Run: func(cmd *cobra.Command, args []string) {

		// init client
		s3Client := c.V1.S3()

		// Get the list of buckets
		output, err := s3Client.ListBuckets(&s3.ListBucketsInput{})
		if err != nil {
			fmt.Println("Error from S3 List", err)
			return
		}

		// Struct to print a basic view
		type basicBucket = struct {
			BucketName string `json:"bucket_name"`
			Owner      string `json:"owner"`
		}
		basicBuckets := []*basicBucket{}

		// Set the struct
		for _, buck := range output.Buckets {
			x := &basicBucket{
				BucketName: *buck.Name,
			}
			x.Owner = *output.Owner.DisplayName
			basicBuckets = append(basicBuckets, x)
		}

		// Print the result
		jsontmpl.Format(jsontmpl.JsonTemplate{
			Fields: []string{"bucket_name", "owner"},
			Data:   basicBuckets,
		})
	},
}

// deleteCmd represents the delete command
var s3DelCmd = &cobra.Command{
	Use:     "delete",
	Example: "vdc delete <name> [<name>] [<name>] ...",
	Short:   "Delete a vdc",

	Run: func(cmd *cobra.Command, args []string) {

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
		s3ListCmd.Run(cmd, []string{})

	},
}

// createCmd represents the create command
var s3CreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "Ceate an bucket",
	Example: "vdc create --name <bucket name>",

	Run: func(cmd *cobra.Command, args []string) {

		// init client
		s3Client := c.V1.S3()

		bucketName, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Malformed bucket name ", err)
			return
		}

		// Create the bucket
		fmt.Println("create a bucket resource (with basic value)")
		fmt.Println("bucket name: " + bucketName)

		_, err = s3Client.CreateBucket(&s3.CreateBucketInput{Bucket: &bucketName})
		if err != nil {
			fmt.Println("Error from S3 Create", err)
			return
		}

		fmt.Println("Bucket resource created successfully !")
		fmt.Println("\nBucket resource list after creation:")
		s3ListCmd.Run(cmd, []string{})

	},
}
