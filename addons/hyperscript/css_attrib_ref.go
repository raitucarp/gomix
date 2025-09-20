package hyperscript

import (
	"fmt"
	"strings"
)

type CssObjectRef struct {
	key   string
	value ValueParam
}

type ValueParam interface {
	String() string
}

func (ref *CssObjectRef) IsRef() {}

func (ref *CssObjectRef) String() string {
	s := []string{}
	if ref.value != nil {
		s = append(s, fmt.Sprintf("{ %s: ", ref.key))
		s = append(s, fmt.Sprintf("%s }", ref.value.String()))
		return strings.Join(s, "")
	}

	return strings.Join(s, fmt.Sprintf(".{'%s'}", ref.key))
}

func CSSObject(key string, value ...ValueParam) *CssObjectRef {
	if len(value) > 0 {
		return &CssObjectRef{key: key, value: value[0]}
	}

	return &CssObjectRef{key: key}
}
