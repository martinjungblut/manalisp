package manalisp

import (
	"fmt"
	"strings"
)

type ManalispType struct {
	Integer        *int64
	Symbol         *string
	List           *[]ManalispType
	Vector         *[]ManalispType
	Hashmap        *[]ManalispType
	NativeFunction *(func(...ManalispType) ManalispType)
}

func (m *ManalispType) IsInteger() bool {
	return m.Integer != nil
}

func (m *ManalispType) AsInteger() int64 {
	return *m.Integer
}

func (m *ManalispType) IsSymbol() bool {
	return m.Symbol != nil
}

func (m *ManalispType) AsSymbol() string {
	return *m.Symbol
}

func NewList() *ManalispType {
	l := make([]ManalispType, 1)
	return &ManalispType{List: &l}
}

func (m *ManalispType) AddToList(t ManalispType) {
	*m.List = append(*m.List, t)
}

func (m *ManalispType) AsList() []ManalispType {
	return *m.List
}

func (m *ManalispType) IsList() bool {
	return m.List != nil
}

func (m *ManalispType) IsEmptyList() bool {
	return m.IsList() && (len(*m.List) == 0)
}

func (m *ManalispType) IsVector() bool {
	return m.Vector != nil
}

func (m *ManalispType) IsHashmap() bool {
	return m.Hashmap != nil
}

func (m *ManalispType) ToString() string {
	wrapSequence := func(sequence *[]ManalispType, lWrap string, rWrap string) string {
		tokens := []string{}
		for _, element := range *sequence {
			if token := element.ToString(); len(token) > 0 {
				tokens = append(tokens, token)
			}
		}
		return fmt.Sprintf("%s%s%s", lWrap, strings.Join(tokens, " "), rWrap)
	}

	if m != nil {
		if m.IsInteger() {
			return fmt.Sprintf("%d", m.AsInteger())
		} else if m.IsSymbol() {
			return m.AsSymbol()
		} else if m.IsList() {
			return wrapSequence(m.List, "(", ")")
		} else if m.IsVector() {
			return wrapSequence(m.Vector, "[", "]")
		} else if m.IsHashmap() {
			return wrapSequence(m.Hashmap, "{", "}")
		} else {
			return ""
		}
	} else {
		return ""
	}
}
