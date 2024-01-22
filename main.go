package main

import (
	"fmt"

	"github.com/orange-cloudavenue/cloudavenue-cli/cmd"
)

//go:generate go run tools/tools.go

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error in Execute main Command", err)
		fmt.Println("Please check your configuration (https://github.com/orange-cloudavenue/cloudavenue-cli/blob/main/docs/index.md)")
	}

}
