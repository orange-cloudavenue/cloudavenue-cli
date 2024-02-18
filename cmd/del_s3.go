package cmd

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

// deleteCmd delete a s3 bucket resource(s)
var delS3Cmd = &cobra.Command{
	Use:               argS3,
	Example:           "delete s3 <name> [<name>] [<name>] ...",
	Short:             "Delete a s3 bucket",
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

		for i, arg := range args {
			s.Stop()
			fmt.Println("delete bucket resource... " + arg)
			s.Restart()
			// Del the bucket
			_, err := c.V1.S3().DeleteBucket(&s3.DeleteBucketInput{Bucket: &args[i]})
			if err != nil {
				return fmt.Errorf("Error from S3 Delete: %w", err)
			}
			s.FinalMSG = "Bucket resource deleted " + arg + " successfully !!\n"
			s.Stop()
		}
		return nil
	},
}
