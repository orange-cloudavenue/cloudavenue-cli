package main

import (
	"strings"

	"github.com/orange-cloudavenue/cloudavenue-sdk-go-v2/commands"
	"github.com/urfave/cli/v3"
)

type output interface {
	Print(f *field, result any, cmd *cli.Command, cavcmd commands.Command)
}

func Print(result any, cmd *cli.Command) error {
	cavcmd := cmd.Metadata["command"].(commands.Command)

	getField := func() (*field, error) {
		// getCliFieldCommands retrieves the fields configuration to display
		field, err := getCliFieldCommands(cavcmd.GetNamespace(), cavcmd.GetResource(), cavcmd.GetVerb())
		if err != nil {
			return nil, err
		}
		return &field, nil
	}

	switch strings.ToLower(cavcmd.GetVerb()) {
	case "get":
		f, err := getField()
		if err != nil {
			// If error occurs, fallback to table output without field configuration
			return cli.Exit(err.Error(), 1)
		}

		OutputTemplate().Print(f, result, cmd, cavcmd)
		// OutputTable().Print(f, result, cmd, cavcmd)
	case "delete", "add", "remove":
		// No output for these verbs
		return cli.Exit("Operation completed successfully.", 0)

	default:
		f, err := getField()
		if err != nil {
			// If error occurs, fallback to table output without field configuration
			return cli.Exit(err.Error(), 1)
		}
		OutputTable().Print(f, result, cmd, cavcmd)
	}

	return nil
}
