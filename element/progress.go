package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type progress struct {
	el *HtmlElement
}

func (p *progress) Max(max float64) *progress {
	p.el.AddAttribute("max", strconv.FormatFloat(max, 'f', -1, 32))
	return p
}

func (p *progress) Value(value float64) *progress {
	p.el.AddAttribute("value", strconv.FormatFloat(value, 'f', -1, 32))
	return p
}

func (p *progress) Element() *HtmlElement { return p.el }
func (p *progress) Render() string        { return p.el.Render() }

func Progress(children IsElement) *progress {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "progress",
		},
	}

	el.Children(children)
	p := &progress{el: el}

	return p
}
