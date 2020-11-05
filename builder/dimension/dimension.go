package dimension

import (
	"encoding/json"

	"github.com/fabriks/go-druid/builder"
	"github.com/fabriks/go-druid/builder/types"
)

type Base struct {
	Typ        builder.ComponentType `json:"type"`
	Dimension  string                `json:"dimension"`
	OutputName string                `json:"outputName"`
	OutputType types.OutputType      `json:"outputType,omitempty"`
}

func (b *Base) SetType(typ builder.ComponentType) *Base {
	b.Typ = typ
	return b
}

func (b *Base) SetDimension(dimension string) *Base {
	b.Dimension = dimension
	return b
}

func (b *Base) SetOutputName(outputName string) *Base {
	b.OutputName = outputName
	return b
}

func (b *Base) SetOutputType(outputType types.OutputType) *Base {
	b.OutputType = outputType
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

func Load(data []byte) (builder.Dimension, error) {
	var t struct {
		Typ builder.ComponentType `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var d builder.Dimension
	switch t.Typ {
	case "default":
		d = NewDefault()
	case "extraction":
		d = NewExtraction()
	case "listFiltered":
		d = NewListFiltered()
	case "lookup":
		d = NewLookup()
	case "prefixFiltered":
		d = NewPrefixFiltered()
	case "regexFiltered":
		d = NewRegexFiltered()
	}
	return d, json.Unmarshal(data, &d)
}
