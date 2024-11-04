package cmd

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

// addS3Cmd add a s3 bucket resource(s)
var addS3Cmd = &cobra.Command{
	Use:               argS3,
	Short:             "Add an S3 bucket",
	Example:           "add s3 --name <bucket name>",
	DisableAutoGenTag: true,
	SilenceErrors:     true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// init Config File & Client
		if err := initConfig(); err != nil {
			return fmt.Errorf("unable to initialize: %w", err)
		}

		// Check if time flag is set
		if cmd.Flag(flagTime).Value.String() == trueValue {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		bucketName, err := cmd.Flags().GetString(flagName)
		if err != nil {
			return fmt.Errorf("unable to retrieve flag %v: %w", flagName, err)
		}

		// Add the bucket
		s.Stop()
		fmt.Println("add a bucket resource (with basic value)")
		fmt.Println("bucket name: " + bucketName)
		s.Restart()

		_, err = c.V1.S3().CreateBucket(&s3.CreateBucketInput{Bucket: &bucketName})
		if err != nil {
			return fmt.Errorf("unable to add bucket S3 %v: %w", bucketName, err)
		}
		s.FinalMSG = "Bucket resource added successfully !!"
		s.Stop()
		return nil
	},
}
