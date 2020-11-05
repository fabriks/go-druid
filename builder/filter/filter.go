package filter

import (
	"encoding/json"

	"github.com/fabriks/go-druid/builder"
)

type Base struct {
	Typ builder.ComponentType `json:"type"`
}

func (b *Base) SetType(typ builder.ComponentType) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

func Load(data []byte) (builder.Filter, error) {
	var t struct {
		Typ builder.ComponentType `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var f builder.Filter
	switch t.Typ {
	case "and":
		f = NewAnd()
	case "bound":
		f = NewBound()
	case "columnComparison":
		f = NewColumnComparison()
	case "expression":
		f = NewExpression()
	case "extraction":
		f = NewExtraction()
	case "false":
		f = NewFalse()
	case "filterTuning":
		f = NewFilterTuning()
	case "in":
		f = NewIn()
	case "interval":
		f = NewInterval()
	case "javascript":
		f = NewJavascript()
	case "like":
		f = NewLike()
	case "not":
		f = NewNot()
	case "or":
		f = NewOr()
	case "regex":
		f = NewRegex()
	case "search":
		f = NewSearch()
	case "selector":
		f = NewSelector()
	case "spatial":
		f = NewSpatial()
	case "true":
		f = NewTrue()
	}
	return f, json.Unmarshal(data, &f)
}
