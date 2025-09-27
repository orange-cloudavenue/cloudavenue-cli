package main

import (
	"encoding/json"
	"fmt"

	_ "embed"
)

//go:embed selected_fields.json
var fieldsData []byte

var fC fieldsConfig

type (
	fieldsConfig struct {
		Commands []field `json:"Commands"`
	}

	field struct {
		Namespace          string         `json:"namespace"`
		Resource           string         `json:"resource"`
		Verb               string         `json:"verb"`
		SelectedFields     []fieldDetail  `json:"selected_fields"`
		SelectedFieldsWide []fieldDetail  `json:"selected_fields_wide"`
		BuildedFields      []fieldBuilded `json:"builded_fields"`
	}

	fieldDetail struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Index   int    `json:"index,omitempty"`
	}

	fieldBuilded struct {
		Bindings   []string `json:"bindings"`
		Display    string   `json:"display"`
		Expression string   `json:"expression"`
		ID         string   `json:"id"`
		Name       string   `json:"name"`
	}
)

func init() {
	err := decoderFields()
	if err != nil {
		panic(err)
	}
}

func decoderFields() error {
	// Decode the JSON data into the struct
	err := json.Unmarshal(fieldsData, &fC)
	if err != nil {
		return err
	}

	return nil
}

func getCliFieldCommands(ns string, resource string, verb string) (f field, err error) {
	for _, cmd := range fC.Commands {
		if cmd.Namespace == ns && cmd.Resource == resource && cmd.Verb == verb {
			return cmd, nil
		}
	}

	// if not found and verb is update, try to find get command
	if verb == "Update" {
		// try to find create command
		for _, cmd := range fC.Commands {
			if cmd.Namespace == ns && cmd.Resource == resource && cmd.Verb == "Get" {
				return cmd, nil
			}
		}
	}

	return field{}, fmt.Errorf("no command found for namespace: %s, resource: %s, verb: %s", ns, resource, verb)
}
