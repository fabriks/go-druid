package extractionfn

import (
	"encoding/json"

	"github.com/fabriks/go-druid/builder"
	"github.com/fabriks/go-druid/builder/searchqueryspec"
)

type SearchQuery struct {
	Base
	Query builder.SearchQuerySpec `json:"query"`
}

func NewSearchQuery() *SearchQuery {
	s := &SearchQuery{}
	s.SetType("searchQuery")
	return s
}

func (s *SearchQuery) SetQuery(q builder.SearchQuerySpec) *SearchQuery {
	s.Query = q
	return s
}

func (s *SearchQuery) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Query json.RawMessage `json:"query"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	q, err := searchqueryspec.Load(tmp.Query)
	if err != nil {
		return err
	}
	s.Base = tmp.Base
	s.Query = q
	return nil
}
