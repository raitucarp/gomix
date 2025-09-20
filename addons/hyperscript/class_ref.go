package hyperscript

import "strings"

type ClassRef struct {
	name      string
	innerHtml bool
}

func (c *ClassRef) IsRef() {}
func (c *ClassRef) InnerHtml() *ClassRef {
	c.innerHtml = true
	return c
}

func (c *ClassRef) String() string {
	s := []string{".", c.name}

	if c.innerHtml {
		s = append(s, ".innerHTML")

	}

	return strings.Join(s, "")
}

func Class(name string) *ClassRef {
	return &ClassRef{name: name}
}
