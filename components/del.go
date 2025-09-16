package components

import (
	"time"

	"golang.org/x/net/html"
)

type del struct {
	comp *Component
}

func (d *del) Cite(url string) *Component {
	d.comp.addAttribute("url", url)
	return d.comp
}

func (d *del) DateTime(datetime time.Time) *Component {
	d.comp.addAttribute("datetime", datetime.Format("2006-01-02T15:04:05-07:00Z"))
	return d.comp
}

func (d *del) Component() *Component {
	return d.comp
}

func Del(child *Component) *del {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "del",
		},
	}

	comp.Children(child)
	d := &del{comp: comp}
	return d
}
