package element

import (
	"time"

	"golang.org/x/net/html"
)

type ins struct {
	el *HtmlElement
}

func (i *ins) Cite(url string) *ins {
	i.el.AddAttribute("cite", url)
	return i
}

func (i *ins) DateTime(date time.Time) *ins {
	i.el.AddAttribute("datetime", date.Format(time.RFC3339))
	return i
}

func (i *ins) Element() *HtmlElement { return i.el }
func (i *ins) Render() string        { return i.el.Render() }

func Ins(children IsElement) *ins {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "ins",
		},
	}

	el.Children(children)

	i := &ins{el: el}

	return i
}
