package hyperscript

import "strings"

type IdRef struct {
	id string
}

func (ref *IdRef) IsRef() {}
func (ref *IdRef) String() string {
	s := []string{"#", ref.id}
	return strings.Join(s, "")
}

func Id(id string) *IdRef {
	return &IdRef{id: id}
}
