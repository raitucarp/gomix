package element

import "golang.org/x/net/html"

type div struct {
	el *HtmlElement
}

func (d *div) Element() *HtmlElement { return d.el }
func (d *div) Render() string        { return d.el.Render() }

func Div(children ...IsElement) *div {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "div",
		},
	}

	el.Children(children...)

	d := &div{el: el}

	return d
}
