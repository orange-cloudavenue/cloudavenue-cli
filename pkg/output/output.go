package output

import (
	"fmt"

	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output/model"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output/ojson"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/output/oyaml"
)

func New(format model.TypeFormat, data any) (model.Formatter, error) {
	switch format {
	case model.TypeJSON:
		return ojson.New(data)
	case model.TypeYAML:
		return oyaml.New(data)
	}
	return nil, fmt.Errorf("error creating output")
}
