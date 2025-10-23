package components

import (
	"github.com/raitucarp/gomix/element"
	"golang.org/x/net/html"
)

type slot struct {
	component *Comp
}

func Slot() *slot {
	slot := &slot{
		component: &Comp{
			El: element.CreateElementByNode(&html.Node{
				Type: html.ElementNode,
				Data: "slot",
			}),
		},
	}

	return slot
}

func (s *slot) Element() *element.HtmlElement {
	return s.component.El
}
