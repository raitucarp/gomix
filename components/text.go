package components

import (
	"github.com/raitucarp/gomix/element"
)

type text struct {
	component *Comp
}

func Text(content string) *text {
	text := &text{
		component: &Comp{
			El: element.P(element.Text(content)).Element(),
		},
	}

	text.component.Me(0).Ms(0).My(0).Mx(0)

	return text
}

func (s *text) Element() *element.HtmlElement {
	return s.component.El
}

func (s *text) Component() *Comp {
	return s.component
}
