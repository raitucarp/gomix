package components

import "golang.org/x/net/html"

type details struct {
	comp *Component
}

func (d *details) Open() *Component {
	d.comp.addAttributeBool("open")
	return d.comp
}

func (d *details) Name(groupname string) *Component {
	d.comp.addAttribute("name", groupname)
	return d.comp
}

func (d *details) Component() *Component {
	return d.comp
}

func Details(children ...*Component) *details {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "details",
		},
	}

	comp.Children(children...)
	d := &details{comp: comp}
	return d
}
