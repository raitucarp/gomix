package element

import (
	"time"

	"golang.org/x/net/html"
)

type timeEl struct {
	el *HtmlElement
}

func (t *timeEl) DateTime(date time.Time) *timeEl {
	t.el.AddAttribute("datetime", date.Format(time.RFC3339))
	return t
}
func (t *timeEl) Element() *HtmlElement { return t.el }
func (t *timeEl) Render() string        { return t.el.Render() }

func Time(children ...IsElement) *timeEl {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "time",
		},
	}

	el.Children(children...)
	t := &timeEl{el: el}

	return t
}
