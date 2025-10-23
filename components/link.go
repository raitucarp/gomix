package components

import (
	"github.com/raitucarp/gomix/element"
)

type link struct {
	component *Comp
}

func Link(children ...IsComponent) *link {
	l := &link{
		component: &Comp{
			El: element.A().Element(),
		},
	}

	l.Component().Children(children...)

	return l
}

func (l *link) Element() *element.HtmlElement {
	return l.component.El
}

func (l *link) Component() *Comp {
	return l.component
}

func (l *link) Href(url string) *link {
	l.Element().AddAttribute("href", url)
	return l
}

func (l *link) External() *link {
	l.Element().AddAttribute("_target", "blank")
	return l
}
