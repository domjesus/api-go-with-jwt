package models

type Occurrence struct {
	Name string `validate:"nonzero" json:"name"`
	Type string `validate:"nonzero" json:"type"`
	Icon string `validate:"nonzero" json:"icon"`
}

var Ocurrences []Occurrence
