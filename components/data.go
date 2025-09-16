package components

import "golang.org/x/net/html"

type data struct {
	comp *Component
}

func (d *data) Value(value string) *Component {
	d.comp.addAttribute("value", value)
	return d.comp
}

func (d *data) Component() *Component {
	return d.comp
}

func Data(child *Component) *data {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "data",
		},
	}

	comp.Children(child)
	d := &data{comp: comp}
	return d
}
