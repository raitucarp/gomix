package element

import "golang.org/x/net/html"

type audio struct {
	el *HtmlElement
}

func (a *audio) Autoplay() *HtmlElement {
	a.el.AddAttributeBool("autoplay")
	return a.el
}

func (a *audio) Controls() *HtmlElement {
	a.el.AddAttributeBool("controls")
	return a.el
}

func (a *audio) Loop() *HtmlElement {
	a.el.AddAttributeBool("loop")
	return a.el
}

func (a *audio) Muted() *HtmlElement {
	a.el.AddAttributeBool("muted")
	return a.el
}

type AudioPreload string

const (
	AudioPreloadAuto     AudioPreload = "auto"
	AudioPreloadMetadata AudioPreload = "metadata"
	AudioPreloadNone     AudioPreload = "none"
)

func (a *audio) Preload(preload AudioPreload) *HtmlElement {
	a.el.AddAttribute("preload", string(preload))
	return a.el
}

func (a *audio) Src(url string) *HtmlElement {
	a.el.AddAttribute("src", url)
	return a.el
}

func (a *audio) Children(children ...IsElement) *audio {
	a.el.Children(children...)
	return a
}
func (a *audio) Component()            {}
func (a *audio) isHtml()               {}
func (a *audio) Element() *HtmlElement { return a.el }
func (a *audio) Render() string        { return a.el.Render() }

func Audio(child ...IsElement) *audio {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "audio",
		},
	}

	el.Children(child...)
	a := &audio{el: el}
	return a
}
