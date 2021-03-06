package aggregation

type HyperUnique struct {
	Base
	FieldName          string `json:"fieldName"`
	IsInputHyperUnique bool   `json:"isInputHyperUnique"`
	Round              bool   `json:"round"`
}

func NewHyperUnique() *HyperUnique {
	h := &HyperUnique{}
	h.SetType("hyperUnique")
	return h
}

func (h *HyperUnique) SetName(name string) *HyperUnique {
	h.Base.SetName(name)
	return h
}

func (h *HyperUnique) SetFieldName(fieldName string) *HyperUnique {
	h.FieldName = fieldName
	return h
}

func (h *HyperUnique) SetIsInputHyperUnique(isInputHyperUnique bool) *HyperUnique {
	h.IsInputHyperUnique = isInputHyperUnique
	return h
}

func (h *HyperUnique) SetRound(round bool) *HyperUnique {
	h.Round = round
	return h
}
