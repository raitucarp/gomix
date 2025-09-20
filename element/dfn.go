package element

import "golang.org/x/net/html"

type dfn struct {
	el *HtmlElement
}

func (d *dfn) Component()            {}
func (d *dfn) isHtml()               {}
func (d *dfn) element() *HtmlElement { return d.el }
func (d *dfn) Render() string        { return d.el.Render() }

func Dfn(child *HtmlElement) *dfn {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "dfn",
		},
	}

	el.Children(child)
	d := &dfn{el: el}
	return d
}
