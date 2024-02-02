package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(versionCmd())
}

// func versionCmd() return the version of the CLI
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:               "version",
		Short:             "Print the version number of cav",
		Long:              `All software has versions. This is cav's`,
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			s.FinalMSG = "Version: " + version + "\nCommit: " + commit + "\nBuilt at: " + date + "\nBuilt by: " + builtBy
			s.Stop()
		},
	}
}
