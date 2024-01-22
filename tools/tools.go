//go:build exclude

package main

import (
	"log"

	"github.com/orange-cloudavenue/cloudavenue-cli/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	log.Default().Println("Generating documentation")
	if err := doc.GenMarkdownTree(cmd.RootCmd, "./docs/command"); err != nil {
		panic(err)
	}
	return
}
