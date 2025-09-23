package element

import "golang.org/x/net/html"

type figure struct {
	el *HtmlElement
}

func (f *figure) Element() *HtmlElement { return f.el }
func (f *figure) Render() string        { return f.el.Render() }

func Figure(children ...IsElement) *figure {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "figure",
		},
	}

	el.Children(children...)

	f := &figure{el: el}

	return f
}
