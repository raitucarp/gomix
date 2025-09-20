package element

import "golang.org/x/net/html"

type dd struct {
	el *HtmlElement
}

func (d *dd) Component()            {}
func (d *dd) isHtml()               {}
func (d *dd) element() *HtmlElement { return d.el }
func (d *dd) Render() string        { return d.el.Render() }

func Dd(child *HtmlElement) *dd {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "dd",
		},
	}

	el.Children(child)
	d := &dd{el: el}
	return d
}
