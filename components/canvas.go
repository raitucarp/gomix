package components

import "golang.org/x/net/html"

type canvas struct {
	comp *Component
}

func (c *canvas) Component() *Component {
	return c.comp
}

func Canvas(message string) *canvas {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "canvas",
		},
	}

	comp.Children(TextRaw(message))
	c := &canvas{comp: comp}
	return c
}
