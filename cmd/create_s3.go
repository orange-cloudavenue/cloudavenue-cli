package cmd

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

// createS3Cmd create a s3 bucket resource(s)
var createS3Cmd = &cobra.Command{
	Use:               argS3,
	Short:             "Create an S3 bucket",
	Example:           "create s3 --name <bucket name>",
	DisableAutoGenTag: true,
	SilenceErrors:     true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// init Config File & Client
		if err := initConfig(); err != nil {
			return fmt.Errorf("Unable to initialize: %w", err)
		}

		// Check if time flag is set
		if cmd.Flag(flagTime).Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		bucketName, err := cmd.Flags().GetString(flagName)
		if err != nil {
			return fmt.Errorf("Unable to retrieve flag %v: %w", flagName, err)
		}

		// Create the bucket
		s.Stop()
		fmt.Println("create a bucket resource (with basic value)")
		fmt.Println("bucket name: " + bucketName)
		s.Restart()

		_, err = c.V1.S3().CreateBucket(&s3.CreateBucketInput{Bucket: &bucketName})
		if err != nil {
			return fmt.Errorf("Unable to create bucket S3 %v: %w", bucketName, err)
		}
		s.FinalMSG = "Bucket resource created successfully !!"
		s.Stop()
		return nil
	},
}
