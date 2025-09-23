package element

import (
	"golang.org/x/net/html"
)

type param struct {
	el *HtmlElement
}

func (p *param) Name(name string) *param {
	p.el.AddAttribute("name", name)
	return p
}

func (p *param) Value(value string) *param {
	p.el.AddAttribute("value", value)
	return p
}

func (p *param) Element() *HtmlElement { return p.el }
func (p *param) Render() string        { return p.el.Render() }

func Param() *param {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "param",
		},
	}

	p := &param{el: el}

	return p
}
