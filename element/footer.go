package element

import "golang.org/x/net/html"

type footer struct {
	el *HtmlElement
}

func (f *footer) Element() *HtmlElement { return f.el }
func (f *footer) Render() string        { return f.el.Render() }

func Footer(children ...IsElement) *footer {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "footer",
		},
	}

	el.Children(children...)

	f := &footer{el: el}

	return f
}
