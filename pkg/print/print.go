package print

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// var fs string

type Writer struct {
	tw *tabwriter.Writer
}

// New Writer
func New() Writer {
	return Writer{
		tw: tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0),
	}
}

// Create format string
func format(fields ...any) (fs string) {
	fs = ""
	for _, field := range fields {
		switch field.(type) {
		case string:
			fs += "%s\t"
		case int:
			fs += "%d\t"
		case float64:
			fs += "%f\t"
		case bool:
			fs += "%t\t"
		default:
			fs += "%v\t"
		}
	}
	fs += "\n"
	return fs
}

// Set Header fieds into upper case
func (w Writer) SetHeader(fields ...any) {
	if len(fields) == 0 {
		return
	}
	fs := format(fields...)
	for i, field := range fields {
		switch field.(type) {
		case string:
			fields[i] = strings.ToUpper(field.(string))
		}
	}
	fmt.Fprintf(w.tw, fs, fields...)
}

// AddFields to the table
func (w Writer) AddFields(fields ...any) {
	if len(fields) == 0 {
		return
	}
	fs := format(fields...)
	fmt.Fprintf(w.tw, fs, fields...)
}

// Print a line
func (w Writer) PrintTable() {
	w.tw.Flush()
}
