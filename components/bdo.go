package components

import "golang.org/x/net/html"

type bdo struct {
	comp *Component
}

func (b *bdo) Ltr() *bdo {
	b.comp.addAttribute("dir", string(DirLtr))
	return b
}

func (b *bdo) Rtl() *bdo {
	b.comp.addAttribute("dir", string(DirRtl))
	return b
}

func (b *bdo) Dir(dir Dir) *bdo {
	b.comp.addAttribute("dir", string(dir))
	return b
}

func (b *bdo) Component() *Component {
	return b.comp
}

func Bdo() *bdo {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "bdo",
		},
	}

	b := &bdo{comp: comp}
	return b
}
