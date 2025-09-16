package components

import "golang.org/x/net/html"

type base struct {
	comp *Component
}

func (b *base) Href(url string) *Component {
	b.comp.addAttribute("href", url)
	return b.comp
}

func (b *base) Target(target AnchorTarget) *Component {
	b.comp.addAttribute("target", string(target))
	return b.comp
}

func (b *base) Component() *Component {
	return b.comp
}

func Base() *base {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "base",
		},
	}

	b := &base{comp: comp}
	return b
}
