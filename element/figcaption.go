package element

import "golang.org/x/net/html"

type figcaption struct {
	el *HtmlElement
}

func (f *figcaption) Element() *HtmlElement { return f.el }
func (f *figcaption) Render() string        { return f.el.Render() }

func FigCaption(children IsElement) *figcaption {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "figcaption",
		},
	}

	el.Children(children)

	f := &figcaption{el: el}

	return f
}
