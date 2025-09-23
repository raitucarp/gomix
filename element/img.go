package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type img struct {
	el *HtmlElement
}

func (i *img) Alt(text string) *img {
	i.el.AddAttribute("alt", text)
	return i
}

func (i *img) CrossOrigin(origin CrossOrigin) *img {
	i.el.AddAttribute("crossorigin", string(origin))
	return i
}

func (i *img) Height(height int) *img {
	i.el.AddAttribute("height", strconv.Itoa(height))
	return i
}

func (i *img) IsMap() *img {
	i.el.AddAttributeBool("ismap")
	return i
}

func (i *img) Loading(loading Loading) *img {
	i.el.AddAttribute("loading", string(loading))
	return i
}

func (i *img) LongDesc(url string) *img {
	i.el.AddAttribute("longdesc", string(url))
	return i
}

type ImageReferrerPolicy string

const (
	ImageNoReferrer                    ImageReferrerPolicy = "no-referrer"                //	No referrer information is sent
	ImageNoReferrerWhenDowngrade       ImageReferrerPolicy = "no-referrer-when-downgrade" //	Default. The referrer header will not be sent to origins without HTTPS
	ImageReferrerOrigin                ImageReferrerPolicy = "origin"                     // Sends the origin (scheme, host, and port) of the document
	ImageReferrerOriginWhenCrossOrigin ImageReferrerPolicy = "origin-when-cross-origin"   //	For cross-origin requests: Send only scheme, host, and port. For same-origin requests: Also include the path
	ImageUnsafeUrl                     ImageReferrerPolicy = "unsafe-url"                 //	Send origin, path and query string (but not fragment, password, or username). This value is considered unsafe

)

func (i *img) Referrer(referrer ImageReferrerPolicy) *img {
	i.el.AddAttribute("referrer", string(referrer))
	return i
}

func (i *img) Sizes(sizes string) *img {
	i.el.AddAttribute("sizes", sizes)
	return i
}

func (i *img) Src(url string) *img {
	i.el.AddAttribute("src", url)
	return i
}

func (i *img) SrcSet(srcset string) *img {
	i.el.AddAttribute("srcset", srcset)
	return i
}

func (i *img) UseMap(mapname string) *img {
	i.el.AddAttribute("usemap", mapname)
	return i
}

func (i *img) Width(width int) *img {
	i.el.AddAttribute("width", strconv.Itoa(width))
	return i
}

func (i *img) Element() *HtmlElement { return i.el }
func (i *img) Render() string        { return i.el.Render() }

func Img(children IsElement) *img {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "img",
		},
	}

	i := &img{el: el}

	return i
}
