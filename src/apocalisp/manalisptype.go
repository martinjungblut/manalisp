package apocalisp

import (
	"fmt"
	"strings"
)

type ApocalispType struct {
	Integer        *int64
	Symbol         *string
	List           *[]ApocalispType
	Vector         *[]ApocalispType
	Hashmap        *[]ApocalispType
	NativeFunction *(func(...ApocalispType) ApocalispType)
}

func (node *ApocalispType) ToString() string {
	wrapSequence := func(sequence *[]ApocalispType, lWrap string, rWrap string) string {
		tokens := []string{}
		for _, element := range *sequence {
			if token := element.ToString(); len(token) > 0 {
				tokens = append(tokens, token)
			}
		}
		return fmt.Sprintf("%s%s%s", lWrap, strings.Join(tokens, " "), rWrap)
	}

	if node != nil {
		if node.IsInteger() {
			return fmt.Sprintf("%d", node.AsInteger())
		} else if node.IsSymbol() {
			return node.AsSymbol()
		} else if node.IsList() {
			return wrapSequence(node.List, "(", ")")
		} else if node.IsVector() {
			return wrapSequence(node.Vector, "[", "]")
		} else if node.IsHashmap() {
			return wrapSequence(node.Hashmap, "{", "}")
		} else if node.IsNativeFunction() {
			return node.AsSymbol()
		} else {
			return ""
		}
	} else {
		return ""
	}
}

// integer
func (node *ApocalispType) IsInteger() bool {
	return node.Integer != nil
}

func (node *ApocalispType) AsInteger() int64 {
	return *node.Integer
}

// symbol
func (node *ApocalispType) IsSymbol() bool {
	return node.Symbol != nil
}

func (node *ApocalispType) AsSymbol() string {
	return *node.Symbol
}

// list
func NewList() *ApocalispType {
	l := make([]ApocalispType, 1)
	return &ApocalispType{List: &l}
}

func (node *ApocalispType) AddToList(t ApocalispType) {
	*node.List = append(*node.List, t)
}

func (node *ApocalispType) AsList() []ApocalispType {
	return *node.List
}

func (node *ApocalispType) IsList() bool {
	return node.List != nil
}

func (node *ApocalispType) IsEmptyList() bool {
	return node.IsList() && (len(*node.List) == 0)
}

// vector
func NewVector() *ApocalispType {
	l := make([]ApocalispType, 1)
	return &ApocalispType{Vector: &l}
}

func (node *ApocalispType) AddToVector(t ApocalispType) {
	*node.Vector = append(*node.Vector, t)
}

func (node *ApocalispType) AsVector() []ApocalispType {
	return *node.Vector
}

func (node *ApocalispType) IsVector() bool {
	return node.Vector != nil
}

func (node *ApocalispType) IsEmptyVector() bool {
	return node.IsVector() && (len(*node.Vector) == 0)
}

// hashmap
func NewHashmap() *ApocalispType {
	l := make([]ApocalispType, 1)
	return &ApocalispType{Hashmap: &l}
}

func (node *ApocalispType) AddToHashmap(t ApocalispType) {
	*node.Hashmap = append(*node.Hashmap, t)
}

func (node *ApocalispType) AsHashmap() []ApocalispType {
	return *node.Hashmap
}

func (node *ApocalispType) IsHashmap() bool {
	return node.Hashmap != nil
}

func (node *ApocalispType) IsEmptyHashmap() bool {
	return node.IsHashmap() && (len(*node.Hashmap) == 0)
}

// native function
func (node *ApocalispType) IsNativeFunction() bool {
	return node.NativeFunction != nil
}

func (node *ApocalispType) CallNativeFunction(parameters ...ApocalispType) ApocalispType {
	return (*node.NativeFunction)(parameters...)
}
