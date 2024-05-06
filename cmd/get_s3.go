package cmd

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/customErrors"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output"
	"github.com/orange-cloudavenue/common-go/print"
	"github.com/spf13/cobra"
)

// getS3Cmd return a list of your s3 (bucket) resource(s)
var getS3Cmd = &cobra.Command{
	Use:               argS3,
	Aliases:           []string{argS3Alias},
	Short:             "A brief list of your s3 resources",
	Long:              "A complete list information of your s3 resources in your CloudAvenue account." + description,
	Example:           "get s3",
	DisableAutoGenTag: true,
	SilenceErrors:     true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		// init Config File & Client
		if err := initConfig(); err != nil {
			return fmt.Errorf("Unable to initialize: %w", err)
		}

		// Check if time flag is set and print time elapsed
		if cmd.Flag(flagTime).Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the list of buckets or a specific bucket
		var s3s *s3.ListBucketsOutput
		s3s, err = c.V1.S3().ListBuckets(&s3.ListBucketsInput{})
		if err != nil {
			return fmt.Errorf("CloudAvenue Error from S3 List: %w", err)
		}
		if cmd.Flag(flagName) != nil && cmd.Flag(flagName).Value.String() != "" {
			for _, b := range s3s.Buckets {
				if *b.Name == cmd.Flag(flagName).Value.String() {
					s3s.Buckets = []*s3.Bucket{b}
				}
			}
		}

		// Print the result
		s.Stop()
		flag := cmd.Flag(flagOutput).Value.String()
		w := print.New()
		// var format output.Formatter
		switch flag {
		case flagOutputValueWide:
			w.SetHeader("name", "owner", "creation date", "owner id")
			for _, b := range s3s.Buckets {
				w.AddFields(*b.Name, *s3s.Owner.DisplayName, *b.CreationDate, *s3s.Owner.ID)
			}
		case flagOutputValueJSON, flagOutputValueYAML:
			x, err := output.New(stringToTypeFormat(flag), s3s)
			if err != nil {
				return fmt.Errorf("Impossible to format output: %w", err)
			}
			x.ToOutput()
		case "":
			w.SetHeader("name", "owner")
			for _, b := range s3s.Buckets {
				w.AddFields(*b.Name, *s3s.Owner.DisplayName)
			}
		default:
			return fmt.Errorf("Output format %v: %w", flag, customErrors.ErrNotValidOutput)
		}
		w.PrintTable()
		return nil
	},
}
