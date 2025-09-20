package element

import "golang.org/x/net/html"

type datalist struct {
	el *HtmlElement
}

func (d *datalist) Component()            {}
func (d *datalist) isHtml()               {}
func (d *datalist) element() *HtmlElement { return d.el }
func (d *datalist) Render() string        { return d.el.Render() }

func DataList(child *HtmlElement) *datalist {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "datalist",
		},
	}

	el.Children(child)
	d := &datalist{el: el}
	return d
}
