//go:build exclude

package main

import (
	"os"
	"path/filepath"

	"github.com/orange-cloudavenue/cloudavenue-cli/cmd"
	"github.com/spf13/cobra/doc"
)

const pathDoc = "./docs/command"

var err error

func main() {
	if err = deleteDirectory(pathDoc); err != nil {
		panic(err)
	}
	if err = doc.GenMarkdownTree(cmd.RootCmd, pathDoc); err != nil {
		panic(err)
	}
	return
}

// Function to delete directory (recursive)
func deleteDirectory(path string) error {
	// Convert relative path to absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Delete directory and sub directories
	err = os.RemoveAll(absPath + "/*")
	if err != nil {
		return err
	}
	return nil
}
