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
	DisableAutoGenTag: true,
	SilenceErrors:     true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// init Config File & Client
		if err := initConfig(); err != nil {
			return fmt.Errorf("Unable to initialize: %w", err)
		}

		// Check if time flag is set and print time elapsed
		if cmd.Flag(flagTime).Value.String() == "true" {
			defer timeTrack(time.Now(), cmd.CommandPath())
		}

		// Get the list of buckets
		s3, err := c.V1.S3().ListBuckets(&s3.ListBucketsInput{})
		if err != nil {
			return fmt.Errorf("CloudAvenue Error from S3 List: %w", err)
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
				return fmt.Errorf("Impossible to format output: %w", err)
			}
			x.ToOutput()
		case "":
			w.SetHeader("name", "owner")
			for _, b := range s3.Buckets {
				w.AddFields(*b.Name, *s3.Owner.DisplayName)
			}
		default:
			return fmt.Errorf("Output format %v: %w", flag, customErrors.ErrNotValidOutput)
		}
		w.PrintTable()
		return nil
	},
}
