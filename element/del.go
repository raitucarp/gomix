package element

import (
	"time"

	"golang.org/x/net/html"
)

type del struct {
	el *HtmlElement
}

func (d *del) Cite(url string) *HtmlElement {
	d.el.AddAttribute("url", url)
	return d.el
}

func (d *del) DateTime(datetime time.Time) *HtmlElement {
	d.el.AddAttribute("datetime", datetime.Format("2006-01-02T15:04:05-07:00Z"))
	return d.el
}

func (d *del) Component()            {}
func (d *del) isHtml()               {}
func (d *del) element() *HtmlElement { return d.el }
func (d *del) Render() string        { return d.el.Render() }

func Del(child *HtmlElement) *del {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "del",
		},
	}

	el.Children(child)
	d := &del{el: el}
	return d
}
