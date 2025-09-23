package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type video struct {
	el *HtmlElement
}

func (v *video) Autoplay() *video {
	v.el.AddAttributeBool("autoplay")
	return v
}

func (v *video) Controls() *video {
	v.el.AddAttributeBool("controls")
	return v
}

func (v *video) Height(height int) *video {
	v.el.AddAttribute("height", strconv.Itoa(height))
	return v
}

func (v *video) Loop() *video {
	v.el.AddAttributeBool("loop")
	return v
}

func (v *video) Muted() *video {
	v.el.AddAttributeBool("muted")
	return v
}

func (v *video) Poster(url string) *video {
	v.el.AddAttribute("poster", url)
	return v
}

type VideoPreload string

const (
	VideoPreloadAuto     VideoPreload = "auto"
	VideoPreloadMetadata VideoPreload = "metadata"
	VideoPreloadNone     VideoPreload = "none"
)

func (v *video) Preload(preload VideoPreload) *video {
	v.el.AddAttribute("preload", string(preload))
	return v
}

func (v *video) Src(url string) *video {
	v.el.AddAttribute("src", url)
	return v
}

func (v *video) Width(width int) *video {
	v.el.AddAttribute("width", strconv.Itoa(width))
	return v
}

func (v *video) Element() *HtmlElement { return v.el }
func (v *video) Render() string        { return v.el.Render() }

func Video(children ...IsElement) *video {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "video",
		},
	}

	el.Children(children...)
	v := &video{el: el}

	return v
}
