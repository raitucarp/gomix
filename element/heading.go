package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type heading struct {
	el   *HtmlElement
	size int
}

func createHeading(size int, children ...IsElement) *heading {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "h" + strconv.Itoa(size),
		},
	}

	el.Children(children...)

	h := &heading{el: el, size: size}

	return h
}

func (h *heading) Element() *HtmlElement { return h.el }
func (h *heading) Render() string        { return h.el.Render() }

func H1(children ...IsElement) *heading { return createHeading(1, children...) }
func H2(children ...IsElement) *heading { return createHeading(2, children...) }
func H3(children ...IsElement) *heading { return createHeading(3, children...) }
func H4(children ...IsElement) *heading { return createHeading(4, children...) }
func H5(children ...IsElement) *heading { return createHeading(5, children...) }
func H6(children ...IsElement) *heading { return createHeading(6, children...) }
