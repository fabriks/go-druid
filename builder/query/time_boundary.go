package query

import (
	"encoding/json"

	"github.com/fabriks/go-druid/builder"
	"github.com/fabriks/go-druid/builder/filter"
	"github.com/fabriks/go-druid/builder/types"
)

type TimeBoundary struct {
	Base
	Bound  string         `json:"bound"`
	Filter builder.Filter `json:"filter"`
}

func NewTimeBoundary() *TimeBoundary {
	t := &TimeBoundary{}
	t.SetQueryType("timeBoundary")
	return t
}

func (t *TimeBoundary) SetDataSource(dataSource builder.DataSource) *TimeBoundary {
	t.Base.SetDataSource(dataSource)
	return t
}

func (t *TimeBoundary) SetIntervals(intervals []types.Interval) *TimeBoundary {
	t.Base.SetIntervals(intervals)
	return t
}

func (t *TimeBoundary) SetContext(context map[string]interface{}) *TimeBoundary {
	t.Base.SetContext(context)
	return t
}

func (t *TimeBoundary) SetBound(bound string) *TimeBoundary {
	t.Bound = bound
	return t
}

func (t *TimeBoundary) SetFilter(filter builder.Filter) *TimeBoundary {
	t.Filter = filter
	return t
}

func (t *TimeBoundary) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Bound  string          `json:"bound"`
		Filter json.RawMessage `json:"filter"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	f, err := filter.Load(tmp.Filter)
	if err != nil {
		return err
	}
	t.Base = tmp.Base
	t.Bound = tmp.Bound
	t.Filter = f
	return nil
}
