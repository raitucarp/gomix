package element

import (
	"golang.org/x/net/html"
)

type track struct {
	el *HtmlElement
}

func (t *track) Default() *track {
	t.el.AddAttributeBool("default")
	return t
}

type TrackKind string

const (
	TrackKindSubtitles    TrackKind = "subtitles"
	TrackKindCaptions     TrackKind = "captions"
	TrackKindDescriptions TrackKind = "descriptions"
	TrackKindChapters     TrackKind = "chapters"
	TrackKindMetadata     TrackKind = "metadata"
)

func (t *track) Kind(kind TrackKind) *track {
	t.el.AddAttribute("kind", string(kind))
	return t
}

func (t *track) Label(label string) *track {
	t.el.AddAttribute("label", label)
	return t
}

func (t *track) Src(url string) *track {
	t.el.AddAttribute("src", url)
	return t
}

func (t *track) SrcLang(code LanguageCode) *track {
	t.el.AddAttribute("srclang", string(code))
	return t
}

func (t *track) Element() *HtmlElement { return t.el }
func (t *track) Render() string        { return t.el.Render() }

func Track(children ...IsElement) *track {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "track",
		},
	}

	el.Children(children...)
	t := &track{el: el}

	return t
}
