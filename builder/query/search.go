package query

import (
	"encoding/json"

	"github.com/fabriks/go-druid/builder"
	"github.com/fabriks/go-druid/builder/dimension"
	"github.com/fabriks/go-druid/builder/filter"
	"github.com/fabriks/go-druid/builder/granularity"
	"github.com/fabriks/go-druid/builder/searchqueryspec"
	"github.com/fabriks/go-druid/builder/types"
)

type SearchSortSpec struct {
	Type types.StringComparator `json:"type"`
}

type Search struct {
	Base
	Filter           builder.Filter          `json:"filter"`
	Granularity      builder.Granularity     `json:"granularity"`
	Limit            int64                   `json:"limit"`
	SearchDimensions []builder.Dimension     `json:"dimensions"`
	Query            builder.SearchQuerySpec `json:"query"`
	Sort             *SearchSortSpec         `json:"sort"`
}

func NewSearch() *Search {
	s := &Search{}
	s.SetQueryType("search")
	return s
}

func (s *Search) SetDataSource(dataSource builder.DataSource) *Search {
	s.Base.SetDataSource(dataSource)
	return s
}

func (s *Search) SetIntervals(intervals []types.Interval) *Search {
	s.Base.SetIntervals(intervals)
	return s
}

func (s *Search) SetContext(context map[string]interface{}) *Search {
	s.Base.SetContext(context)
	return s
}

func (s *Search) SetFilter(filter builder.Filter) *Search {
	s.Filter = filter
	return s
}

func (s *Search) SetGranularity(granularity builder.Granularity) *Search {
	s.Granularity = granularity
	return s
}

func (s *Search) SetLimit(limit int64) *Search {
	s.Limit = limit
	return s
}

func (s *Search) SetSearchDimensions(searchDimensions []builder.Dimension) *Search {
	s.SearchDimensions = searchDimensions
	return s
}

func (s *Search) SetQuery(q builder.SearchQuerySpec) *Search {
	s.Query = q
	return s
}

func (s *Search) SetSort(sort *SearchSortSpec) *Search {
	s.Sort = sort
	return s
}

func (s *Search) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Filter           json.RawMessage   `json:"filter"`
		Granularity      json.RawMessage   `json:"granularity"`
		Limit            int64             `json:"limit"`
		SearchDimensions []json.RawMessage `json:"dimensions"`
		Query            json.RawMessage   `json:"query"`
		Sort             *SearchSortSpec   `json:"sort"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var err error
	f, err := filter.Load(tmp.Filter)
	if err != nil {
		return err
	}
	gr, err := granularity.Load(tmp.Granularity)
	if err != nil {
		return err
	}
	var se builder.Dimension
	ss := make([]builder.Dimension, len(tmp.SearchDimensions))
	for i := range tmp.SearchDimensions {
		if se, err = dimension.Load(tmp.SearchDimensions[i]); err != nil {
			return err
		}
		ss[i] = se
	}
	q, err := searchqueryspec.Load(tmp.Query)
	if err != nil {
		return err
	}
	s.Base = tmp.Base
	s.Filter = f
	s.Granularity = gr
	s.SearchDimensions = ss
	s.Query = q
	s.Sort = tmp.Sort
	return nil
}
