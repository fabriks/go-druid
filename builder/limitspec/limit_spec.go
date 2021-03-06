package limitspec

import (
	"encoding/json"

	"github.com/fabriks/go-druid/builder"
)

type Base struct {
	Typ builder.ComponentType
}

func (b *Base) SetType(typ builder.ComponentType) *Base {
	b.Typ = typ
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
	}
	return d, json.Unmarshal(data, &d)
}
