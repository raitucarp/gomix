package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type th struct {
	el *HtmlElement
}

func (t *th) Abbr(abbr string) *th {
	t.el.AddAttribute("abbr", abbr)
	return t
}

func (t *th) ColSpan(col int) *th {
	t.el.AddAttribute("colspan", strconv.Itoa(col))
	return t
}

func (t *th) RowSpan(row int) *th {
	t.el.AddAttribute("rowspan", strconv.Itoa(row))
	return t
}

func (t *th) Headers(headerId string) *th {
	t.el.AddAttribute("headers", headerId)
	return t
}

type ThScope string

const (
	ScopeCol    ThScope = "col"
	ScopeRow    ThScope = "row"
	ScopeColGrp ThScope = "colgroup"
	ScopeRowGrp ThScope = "rowgroup"
)

func (t *th) Scope(scope ThScope) *th {
	t.el.AddAttribute("scope", string(scope))
	return t
}

func (t *th) Element() *HtmlElement { return t.el }
func (t *th) Render() string        { return t.el.Render() }

func Th(children ...IsElement) *th {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "th",
		},
	}

	el.Children(children...)
	t := &th{el: el}

	return t
}
