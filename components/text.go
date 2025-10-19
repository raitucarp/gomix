package components

import (
	"github.com/raitucarp/gomix/element"
)

type text struct {
	component *component
}

func Text(content string) *text {
	text := &text{
		component: &component{
			el: element.P(element.Text(content)).Element(),
		},
	}

	return text
}

func (s *text) Element() *element.HtmlElement {
	return s.component.el
}
