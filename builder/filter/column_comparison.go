package filter

import (
	"encoding/json"

	"github.com/fabriks/go-druid/builder"
	"github.com/fabriks/go-druid/builder/dimension"
)

type ColumnComparison struct {
	Base
	Dimensions []builder.Dimension `json:"dimensions"`
}

func NewColumnComparison() *ColumnComparison {
	c := &ColumnComparison{}
	c.SetType("columnComparison")
	return c
}

func (c *ColumnComparison) SetDimensions(dimensions []builder.Dimension) *ColumnComparison {
	c.Dimensions = dimensions
	return c
}

func (c *ColumnComparison) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Dimensions []json.RawMessage `json:"dimensions"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var err error
	var d builder.Dimension
	dd := make([]builder.Dimension, len(tmp.Dimensions))
	for i := range tmp.Dimensions {
		if d, err = dimension.Load(tmp.Dimensions[i]); err != nil {
			return err
		}
		dd[i] = d
	}
	c.Base = tmp.Base
	c.Dimensions = dd
	return nil
}
