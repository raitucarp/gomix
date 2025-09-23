package element

import "golang.org/x/net/html"

type em struct {
	el *HtmlElement
}

func (e *em) Element() *HtmlElement { return e.el }
func (e *em) Render() string        { return e.el.Render() }

func Em(children IsElement) *em {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "em",
		},
	}

	el.Children(children)

	d := &em{el: el}

	return d
}
