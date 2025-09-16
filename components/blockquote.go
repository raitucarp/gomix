package components

import "golang.org/x/net/html"

type blockquote struct {
	comp *Component
}

func (b *blockquote) Cite(url string) *blockquote {
	b.comp.addAttribute("cite", url)
	return b
}

func (b *blockquote) Component() *Component {
	return b.comp
}

func Blockquote(child *Component) *blockquote {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "blockquote",
		},
	}

	comp.Children(child)
	b := &blockquote{comp: comp}
	return b
}
