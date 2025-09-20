package element

import "golang.org/x/net/html"

type data struct {
	el *HtmlElement
}

func (d *data) Value(value string) *HtmlElement {
	d.el.AddAttribute("value", value)
	return d.el
}

func (d *data) Component()            {}
func (d *data) isHtml()               {}
func (d *data) element() *HtmlElement { return d.el }
func (d *data) Render() string        { return d.el.Render() }

func Data(child *HtmlElement) *data {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "data",
		},
	}

	el.Children(child)
	d := &data{el: el}
	return d
}
