package element

import "golang.org/x/net/html"

func Text(text string) *HtmlElement {
	return &HtmlElement{
		node: &html.Node{
			Type: html.TextNode,
			Data: text,
		},
	}
}
