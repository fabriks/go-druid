package dimension

import (
	"encoding/json"

	"github.com/fabriks/go-druid/builder"
	"github.com/fabriks/go-druid/builder/lookup"
)

type Lookup struct {
	Base
	Name                    string                  `json:"name"`
	ReplaceMissingValueWith string                  `json:"replaceMissingValueWith"`
	RetainMissingValue      bool                    `json:"retainMissingValue"`
	Lookup                  builder.LookupExtractor `json:"lookup"`
	Optimize                bool                    `json:"optimize"`
}

type RegisteredLookup struct {
	Base
	Name string `json:"name"`
}

func NewLookup() *Lookup {
	l := &Lookup{}
	l.SetType("lookup")
	return l
}

func (l *Lookup) SetName(name string) *Lookup {
	l.Name = name
	return l
}

func (l *Lookup) SetReplaceMissingValueWith(replaceMissingValueWith string) *Lookup {
	l.ReplaceMissingValueWith = replaceMissingValueWith
	return l
}

func (l *Lookup) SetRetainMissingValue(retainMissingValue bool) *Lookup {
	l.RetainMissingValue = retainMissingValue
	return l
}

func (l *Lookup) SetLookup(lookup builder.LookupExtractor) *Lookup {
	l.Lookup = lookup
	return l
}

func (l *Lookup) SetOptimize(optimize bool) *Lookup {
	l.Optimize = optimize
	return l
}

func (l *Lookup) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Name                    string          `json:"name"`
		ReplaceMissingValueWith string          `json:"replaceMissingValueWith"`
		RetainMissingValue      bool            `json:"retainMissingValue"`
		Lookup                  json.RawMessage `json:"lookup"`
		Optimize                bool            `json:"optimize"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	lu, err := lookup.Load(tmp.Lookup)
	if err != nil {
		return err
	}
	l.Base = tmp.Base
	l.Name = tmp.Name
	l.ReplaceMissingValueWith = tmp.ReplaceMissingValueWith
	l.RetainMissingValue = tmp.RetainMissingValue
	l.Lookup = lu
	l.Optimize = tmp.Optimize
	return nil
}
