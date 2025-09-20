package element

import "golang.org/x/net/html"

type LinkRelationship string

const (
	RelAlternate   LinkRelationship = "alternate"
	RelAuthor      LinkRelationship = "author"
	RelDnsPrefetch LinkRelationship = "dns-prefetch"
	RelHelp        LinkRelationship = "help"
	RelIcon        LinkRelationship = "icon"
	RelLicense     LinkRelationship = "license"
	RelNext        LinkRelationship = "next"
	RelPingback    LinkRelationship = "pingback"
	RelPreconnect  LinkRelationship = "preconnect"
	RelPrefetch    LinkRelationship = "prefetch"
	RelPreload     LinkRelationship = "preload"
	RelPrerender   LinkRelationship = "prerender"
	RelPrev        LinkRelationship = "prev"
	RelSearch      LinkRelationship = "search"
	RelStylesheet  LinkRelationship = "stylesheet"
)

type link struct {
	el *HtmlElement
}

func (l *link) CrossOrigin(crossorigin CrossOrigin) *link {
	l.el.AddAttribute("crossorigin", string(crossorigin))
	return l
}

func (l *link) Href(url string) *link {
	l.el.AddAttribute("href", url)
	return l
}

func (l *link) HrefLang(code LanguageCode) *link {
	l.el.AddAttribute("hreflang", string(code))
	return l
}

func (l *link) Media(media_query string) *link {
	l.el.AddAttribute("media", media_query)
	return l
}

func (l *link) ReferrerPolicy(policy RefererPolicy) *link {
	l.el.AddAttribute("referrerpolicy", string(policy))
	return l
}

func (l *link) Rel(rel LinkRelationship) *link {
	l.el.AddAttribute("rel", string(rel))
	return l
}

func (l *link) Sizes(sizes string) *link {
	l.el.AddAttribute("sizes", sizes)
	return l
}

func (l *link) Type(media_type string) *link {
	l.el.AddAttribute("type", media_type)
	return l
}
func (l *link) InsideHead()           {}
func (l *link) Element() *HtmlElement { return l.el }
func (l *link) Render() string        { return l.el.Render() }

func Link() *link {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "link",
		},
	}

	l := &link{el: el}

	return l
}
