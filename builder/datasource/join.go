package datasource

import (
	"encoding/json"

	"github.com/fabriks/go-druid/builder"
	"github.com/fabriks/go-druid/builder/types"
)

type Join struct {
	Base
	Left        builder.DataSource `json:"left"`
	Right       builder.DataSource `json:"right"`
	RightPrefix string             `json:"rightPrefix"`
	Condition   string             `json:"condition"`
	JoinType    types.JoinType     `json:"joinType"`
}

func NewJoin() *Join {
	j := &Join{}
	j.SetType("join")
	return j
}

func (j *Join) SetLeft(left builder.DataSource) *Join {
	j.Left = left
	return j
}

func (j *Join) SetRight(right builder.DataSource) *Join {
	j.Right = right
	return j
}

func (j *Join) SetRightPrefix(rightPrefix string) *Join {
	j.RightPrefix = rightPrefix
	return j
}

func (j *Join) SetCondition(condition string) *Join {
	j.Condition = condition
	return j
}

func (j *Join) SetJoinType(joinType types.JoinType) *Join {
	j.JoinType = joinType
	return j
}

func (j *Join) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Left        json.RawMessage `json:"left"`
		Right       json.RawMessage `json:"right"`
		RightPrefix string          `json:"rightPrefix"`
		Condition   string          `json:"condition"`
		JoinType    types.JoinType  `json:"joinType"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	l, err := Load(tmp.Left)
	if err != nil {
		return err
	}
	r, err := Load(tmp.Right)
	if err != nil {
		return err
	}
	j.Base = tmp.Base
	j.Left = l
	j.Right = r
	j.RightPrefix = tmp.RightPrefix
	j.Condition = tmp.Condition
	j.JoinType = tmp.JoinType
	return nil
}
