package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/orange-cloudavenue/cloudavenue-sdk-go-v2/commands"
)

// TODO factorise
func buildLine(fi field, item any) (line []any) {
	for _, f := range fi.SelectedFields {
		// Builded field
		if f.ID != "" {
			for _, bf := range fi.BuildedFields {
				// expression is template like ${.b0}/${.b1}
				if bf.ID == f.ID {
					// Expression is defined in field-ui tool
					expr := bf.Expression
					for _, key := range bf.Bindings {
						v, err := commands.GetValueAtPath(item, key)
						if err != nil {
							v = key
						}
						expr = strings.ReplaceAll(expr, "${"+key+"}", fmt.Sprintf("%v", v))
					}
					line = append(line, expr)
					break
				}
			}
			continue
		}

		// Simple field
		v, err := commands.GetValueAtPath(item, f.Name)
		if err != nil {
			continue
		}
		line = append(line, v)
	}
	return line
}

func buildListLine(field field, result any) (fieldsLines [][]any) {
	if result == nil {
		return fieldsLines
	}

	// Find in the selected fields the key of slices
	// ex vdcs.{index}.name -> vdcs
	sliceKey := ""
	for _, f := range field.SelectedFields {
		parts := strings.Split(f.Name, ".")
		if len(parts) >= 2 && strings.EqualFold(parts[1], "{index}") {
			sliceKey = parts[0]
			break
		}
	}

	if sliceKey == "" {
		return fieldsLines
	}

	// Determine the length of the slice
	sliceLen := 0
	if sliceKey != "" {
		v, err := commands.GetAllValuesAtTarget(result, sliceKey)
		if err != nil {
			return fieldsLines
		}

		// GetAllValuesAtTarget resturn []interface{[]models.ModelXX{}}
		// Get the len of []models.ModelXX{}
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			s := reflect.ValueOf(v)
			if s.Len() > 0 {
				first := s.Index(0).Interface()
				if reflect.TypeOf(first).Kind() == reflect.Slice {
					s2 := reflect.ValueOf(first)
					sliceLen = s2.Len()
				}
			}
		}

	}
	for i := 0; i < sliceLen; i++ {
		// Clone SelectedFields to avoid overwriting the original
		clonedFields := make([]fieldDetail, len(field.SelectedFields))
		for ii, sf := range field.SelectedFields {
			nSf := sf
			nSf.Name = strings.ReplaceAll(sf.Name, "{index}", fmt.Sprintf("%d", i))
			clonedFields[ii] = nSf
		}

		clonedFieldsWide := make([]fieldDetail, len(field.SelectedFieldsWide))
		for ii, sf := range field.SelectedFieldsWide {
			nSf := sf
			nSf.Name = strings.ReplaceAll(sf.Name, "{index}", fmt.Sprintf("%d", i))
			clonedFieldsWide[ii] = nSf
		}

		clonedBuildedFields := make([]fieldBuilded, len(field.BuildedFields))
		for ii, fb := range field.BuildedFields {
			clonedBindings := make([]string, len(fb.Bindings))
			copy(clonedBindings, fb.Bindings)
			fb.Bindings = clonedBindings
			// replace {index} in bindings and expression
			for ii, b := range fb.Bindings {
				nB := strings.ReplaceAll(b, "{index}", fmt.Sprintf("%d", i))
				clonedBindings[ii] = nB
			}
			fb.Expression = strings.ReplaceAll(fb.Expression, "{index}", fmt.Sprintf("%d", i))
			fb.Bindings = clonedBindings
			clonedBuildedFields[ii] = fb
		}

		ff := field
		ff.SelectedFields = clonedFields
		ff.SelectedFieldsWide = clonedFieldsWide
		ff.BuildedFields = clonedBuildedFields
		line := buildLine(ff, result)
		fieldsLines = append(fieldsLines, line)
	}
	return fieldsLines
}
