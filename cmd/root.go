/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/adampresley/sigint"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() (err error) {
	defer s.Stop()
	// ctrl+c handler
	sigint.ListenForSIGINT(func() {
		s.Stop()
		fmt.Println("SIGINT received. Exiting...")
		os.Exit(0)
	})

	// Execute root command
	if err = rootCmd.Execute(); err != nil {
		return fmt.Errorf("Command return an Error: %w", err)
	}
	return nil
}

// NewRootCmd returns the root command
func NewRootCmd() *cobra.Command {
	return rootCmd
}

// init initializes the root command
func init() {
	s.Start()
	rootCmd.PersistentFlags().BoolP("time", "t", false, "time elapsed for command")
	rootCmd.InitDefaultCompletionCmd()
}
