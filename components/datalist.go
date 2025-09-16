package components

import "golang.org/x/net/html"

type datalist struct {
	comp *Component
}

func (d *datalist) Component() *Component {
	return d.comp
}

func DataList(child *Component) *datalist {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "datalist",
		},
	}

	comp.Children(child)
	d := &datalist{comp: comp}
	return d
}
