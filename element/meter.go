package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type meter struct {
	el *HtmlElement
}

func (m *meter) Form(formId string) *meter {
	m.el.AddAttribute("form", formId)
	return m
}

func (m *meter) High(high int) *meter {
	m.el.AddAttribute("high", strconv.Itoa(high))
	return m
}

func (m *meter) Low(low int) *meter {
	m.el.AddAttribute("low", strconv.Itoa(low))
	return m
}

func (m *meter) Max(max int) *meter {
	m.el.AddAttribute("max", strconv.Itoa(max))
	return m
}

func (m *meter) Min(min int) *meter {
	m.el.AddAttribute("min", strconv.Itoa(min))
	return m
}

func (m *meter) Optimum(optimum int) *meter {
	m.el.AddAttribute("optimum", strconv.Itoa(optimum))
	return m
}

func (m *meter) Value(value int) *meter {
	m.el.AddAttribute("value", strconv.Itoa(value))
	return m
}

func (m *meter) Element() *HtmlElement { return m.el }
func (m *meter) Render() string        { return m.el.Render() }

func Meter(children IsElement) *meter {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "meter",
		},
	}

	el.Children(children)

	m := &meter{el: el}

	return m
}
