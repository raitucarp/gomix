package element

import "golang.org/x/net/html"

type dialog struct {
	el *HtmlElement
}

func (d *dialog) Open() *HtmlElement {
	d.el.AddAttributeBool("open")
	return d.el
}

func (d *dialog) Component()            {}
func (d *dialog) isHtml()               {}
func (d *dialog) Element() *HtmlElement { return d.el }
func (d *dialog) Render() string        { return d.el.Render() }

func Dialog(child *HtmlElement) *dialog {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "dialog",
		},
	}

	el.Children(child)
	d := &dialog{el: el}
	return d
}
