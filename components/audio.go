package components

import "golang.org/x/net/html"

type audio struct {
	comp *Component
}

func (a *audio) Autoplay() *Component {
	a.comp.addAttributeBool("autoplay")
	return a.comp
}

func (a *audio) Controls() *Component {
	a.comp.addAttributeBool("controls")
	return a.comp
}

func (a *audio) Loop() *Component {
	a.comp.addAttributeBool("loop")
	return a.comp
}

func (a *audio) Muted() *Component {
	a.comp.addAttributeBool("muted")
	return a.comp
}

type AudioPreload string

const (
	AudioPreloadAuto     AudioPreload = "auto"
	AudioPreloadMetadata AudioPreload = "metadata"
	AudioPreloadNone     AudioPreload = "none"
)

func (a *audio) Preload(preload AudioPreload) *Component {
	a.comp.addAttribute("preload", string(preload))
	return a.comp
}

func (a *audio) Src(url string) *Component {
	a.comp.addAttribute("src", url)
	return a.comp
}

func (a *audio) Component() *Component {
	return a.comp
}

func Audio(child ...*Component) *audio {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "audio",
		},
	}

	comp.Children(child...)
	a := &audio{comp: comp}
	return a
}
