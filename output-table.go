package main

import (
	"strings"

	"github.com/orange-cloudavenue/cloudavenue-sdk-go-v2/commands"
	"github.com/orange-cloudavenue/common-go/print"
	"github.com/urfave/cli/v3"
)

var _ output = (*outputTable)(nil)

type outputTable struct{}

func OutputTable() output {
	return &outputTable{}
}

func (o *outputTable) Print(f *field, result any, cmd *cli.Command, cavcmd commands.Command) {

	// Initialize the printer
	x := print.New(print.WithOutput(cmd.Root().Writer))

	// Build headers
	x.SetHeader(func() (headers []any) {
		for _, fi := range f.SelectedFields {
			switch {
			// If ID is set, try to find the field in BuildedFields to get its Display or Name
			case fi.ID != "":
				for _, bf := range f.BuildedFields {
					if bf.ID == fi.ID {
						if bf.Display != "" {
							headers = append(headers, bf.Display)
						} else {
							headers = append(headers, bf.Name)
						}
						break
					}
				}

			case fi.Display != "":
				headers = append(headers, fi.Display)
			case fi.Name != "":
				headers = append(headers, fi.Name)
			}
		}
		return headers
	}()...)

	fieldsLines := [][]any{}

	switch strings.ToLower(cavcmd.GetVerb()) {
	case "create", "update", "get":
		// Single object
		line := buildLine(*f, result)
		fieldsLines = append(fieldsLines, line)
	case "list":
		// List of objects
		fieldsLines = buildListLine(*f, result)
	default:
		// Fallback to single object
		line := buildLine(*f, result)
		fieldsLines = append(fieldsLines, line)
	}

	for _, line := range fieldsLines {
		x.AddFields(line...)
	}
	x.PrintTable()
}
