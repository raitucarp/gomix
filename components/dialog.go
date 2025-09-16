package components

import "golang.org/x/net/html"

type dialog struct {
	comp *Component
}

func (d *dialog) Open() *Component {
	d.comp.addAttributeBool("open")
	return d.comp
}

func (d *dialog) Component() *Component {
	return d.comp
}

func Dialog(child *Component) *dialog {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "dialog",
		},
	}

	comp.Children(child)
	d := &dialog{comp: comp}
	return d
}
