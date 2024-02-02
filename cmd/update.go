package cmd

import (
	"context"
	"log"
	"os"
	"runtime"

	selfupdate "github.com/creativeprojects/go-selfupdate"
	"github.com/spf13/cobra"
)

const (
	CmdUpdate = "update"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

// updateCmd represents the version command.
var updateCmd = &cobra.Command{
	Use: CmdUpdate,
	// GroupID: "other",
	Short: "Check for updates and update the application",
	Long:  `Check if a new version is available and update the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		s.Stop()
		if version == "dev" {
			log.Printf("Cannot update a development version")
			return
		}
		log.Println("Checking for updates...")
		latest, found, err := selfupdate.DetectLatest(context.Background(), selfupdate.ParseSlug("orange-cloudavenue/cloudavenue-cli"))
		if err != nil {
			log.Default().Printf("error occurred while detecting version: %s", err)
			return
		}
		if !found {
			log.Default().Printf("latest version for %s/%s could not be found from github repository", runtime.GOOS, runtime.GOARCH)
			return
		}

		if latest.LessOrEqual(version) {
			log.Printf("Current version (%s) is the latest", version)
			return
		}

		exe, err := os.Executable()
		if err != nil {
			log.Default().Printf("could not locate executable path")
			return
		}

		log.Printf("Updating to version %s", latest.Version())
		if err := selfupdate.UpdateTo(context.Background(), latest.AssetURL, latest.AssetName, exe); err != nil {
			log.Default().Printf("error occurred while updating binary: %s", err)
			return
		}
		log.Printf("Successfully updated to version %s", latest.Version())
	},
}
