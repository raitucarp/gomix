package element

import "golang.org/x/net/html"

type dl struct {
	el *HtmlElement
}

func (d *dl) Element() *HtmlElement { return d.el }
func (d *dl) Render() string        { return d.el.Render() }

func Dl(children ...IsElement) *dl {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "dl",
		},
	}

	el.Children(children...)

	d := &dl{el: el}

	return d
}
