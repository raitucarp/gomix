package element

import "golang.org/x/net/html"

type details struct {
	el *HtmlElement
}

func (d *details) Open() *HtmlElement {
	d.el.AddAttributeBool("open")
	return d.el
}

func (d *details) Name(groupname string) *HtmlElement {
	d.el.AddAttribute("name", groupname)
	return d.el
}

func (d *details) Component()            {}
func (d *details) isHtml()               {}
func (d *details) element() *HtmlElement { return d.el }
func (d *details) Render() string        { return d.el.Render() }

func Details(children ...IsElement) *details {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "details",
		},
	}

	el.Children(children...)
	d := &details{el: el}
	return d
}
