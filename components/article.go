package components

import "golang.org/x/net/html"

type article struct {
	comp *Component
}

func (a *article) Component() *Component {
	return a.comp
}

func Article(child ...*Component) *article {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "article",
		},
	}

	comp.Children(child...)
	art := &article{comp: comp}
	return art
}
