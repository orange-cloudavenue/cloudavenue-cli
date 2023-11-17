package tmpl_json

// package main

import (
	"encoding/json"
	"fmt"

	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
)

const (
	delimiter = "|"
)

type JsonTemplate struct {
	Fields []string
	Data   interface{}
}

func Format(j JsonTemplate) {
	var inInterface []interface{}
	inrec, err := json.Marshal(j.Data)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(inrec, &inInterface); err != nil {
		panic(err)
	}

	prettyPrintedTable, err := markdown.NewTableFormatterBuilder().
		WithPrettyPrint().
		Build(j.Fields...).
		Format(func() [][]string {
			var s [][]string
			for _, v := range inInterface {
				var row []string
				for _, v2 := range j.Fields {
					row = append(row, fmt.Sprintf("%s", v.(map[string]interface{})[v2]))
				}
				s = append(s, row)
			}
			return s
		}())

	if err != nil {
		panic(err)
	}
	fmt.Println(prettyPrintedTable)
}

func Format2(j JsonTemplate) {
	var inInterface interface{}
	inrec, err := json.Marshal(j.Data)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(inrec, &inInterface); err != nil {
		panic(err)
	}

	prettyPrintedTable, err := markdown.NewTableFormatterBuilder().
		WithPrettyPrint().
		Build(j.Fields...).
		Format(func() [][]string {
			var s [][]string
			var row []string
			for _, v2 := range j.Fields {
				row = append(row, fmt.Sprintf("%v", inInterface.(map[string]interface{})[v2]))
			}
			s = append(s, row)
			return s
		}())

	if err != nil {
		panic(err)
	}
	fmt.Println(prettyPrintedTable)
}
