package element

import (
	"golang.org/x/net/html"
)

type pre struct {
	el *HtmlElement
}

func (p *pre) Element() *HtmlElement { return p.el }
func (p *pre) Render() string        { return p.el.Render() }

func Pre(children IsElement) *pre {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "pre",
		},
	}

	el.Children(children)
	p := &pre{el: el}

	return p
}
