/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/orange-cloudavenue/cloudavenue-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error in Execute main Command", err)
		fmt.Println("Please check your configuration (https://github.com/orange-cloudavenue/cloudavenue-cli/blob/main/docs/index.md)")
	}

}
