package element

import "golang.org/x/net/html"

type dt struct {
	el *HtmlElement
}

func (d *dt) Element() *HtmlElement { return d.el }
func (d *dt) Render() string        { return d.el.Render() }

func Dt(children IsElement) *dt {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "dt",
		},
	}

	el.Children(children)

	d := &dt{el: el}

	return d
}
