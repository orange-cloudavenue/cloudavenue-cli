package oyaml

import (
	"fmt"

	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output/model"
	"gopkg.in/yaml.v3"
)

var _ model.Formatter = oyaml{}

type oyaml struct {
	data          any
	humanReadable string
}

func New(data any) (oyaml, error) {
	y := oyaml{data: data}
	err := y.toMarshall()
	if err != nil {
		return oyaml{}, err
	}
	return y, nil
}

func (y *oyaml) toMarshall() error {
	yamlData, err := yaml.Marshal(y.data)
	if err != nil {
		return err
	}
	y.humanReadable = string(yamlData)
	return nil
}

func (y oyaml) ToOutput() {
	fmt.Println(y.humanReadable)
}
