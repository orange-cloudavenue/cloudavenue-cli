package ojson

import (
	"encoding/json"
	"fmt"

	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output/model"
)

var _ model.Formatter = Ojson{}

type Ojson struct {
	data          any
	humanReadable string
}

func New(data any) (Ojson, error) {
	j := Ojson{data: data}
	err := j.toMarshall()
	if err != nil {
		return Ojson{}, err
	}
	return j, nil
}

// ToOutput file in json format
func (j Ojson) ToOutput() {
	fmt.Println(j.humanReadable)
}

// toMarshall data in JSON Format
func (j *Ojson) toMarshall() error {
	jsonData, err := json.MarshalIndent(j.data, "", "  ")
	if err != nil {
		return err
	}
	j.humanReadable = string(jsonData)
	return nil
}
