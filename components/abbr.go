package components

import "golang.org/x/net/html"

type abbreviation struct {
	comp *Component
}

func (abbr *abbreviation) Component() *Component {
	return abbr.comp
}

func Abbr(child *Component) *abbreviation {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "abbr",
		},
	}

	comp.Children(child)
	abbr := &abbreviation{comp: comp}
	return abbr
}
