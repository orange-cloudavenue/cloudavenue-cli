package model

type Formatter interface {
	ToOutput()
}

type TypeFormat string

const (
	TypeJSON TypeFormat = "json"
	TypeYAML TypeFormat = "yaml"
)

var ListTypeFormat = []TypeFormat{TypeJSON, TypeYAML}
