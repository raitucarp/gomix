package element

import (
	"golang.org/x/net/html"
)

type picture struct {
	el *HtmlElement
}

func (p *picture) Name(name string) *picture {
	p.el.AddAttribute("name", name)
	return p
}

func (p *picture) Value(value string) *picture {
	p.el.AddAttribute("value", value)
	return p
}

func (p *picture) Element() *HtmlElement { return p.el }
func (p *picture) Render() string        { return p.el.Render() }

func Picture(children ...IsElement) *picture {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "picture",
		},
	}

	el.Children(children...)
	p := &picture{el: el}

	return p
}
