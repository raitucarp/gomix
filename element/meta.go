package element

import "golang.org/x/net/html"

type meta struct {
	el *HtmlElement
}

func (m *meta) InsideHead() {}
func (m *meta) CharSet(charset string) *meta {
	m.el.AddAttribute("charset", charset)
	return m
}

func (m *meta) Content(content string) *meta {
	m.el.AddAttribute("content", content)
	return m
}

type HttpEquiv string

const (
	ContentSecurityPolicy HttpEquiv = "Content-Security-Policy"
	ContentType           HttpEquiv = "Content-Type"
	DefaultStyle          HttpEquiv = "Default-Style"
	Refresh               HttpEquiv = "Refresh"
	XUACompatible         HttpEquiv = "X-UA-Compatible"
	OriginTrial           HttpEquiv = "Origin-Trial"
)

func (m *meta) HttpEquiv(equiv HttpEquiv) *meta {
	m.el.AddAttribute("http-equiv", string(equiv))
	return m
}

type MetaName string

const (
	MetaNameApplicationName        MetaName = "application-name"
	MetaNameAuthor                 MetaName = "author"
	MetaNameDescription            MetaName = "description"
	MetaNameGenerator              MetaName = "generator"
	MetaNameKeywords               MetaName = "keywords"
	MetaNameReferrer               MetaName = "referrer"
	MetaNameViewport               MetaName = "viewport"
	MetaNameThemeColor             MetaName = "theme-color"
	MetaNameRobots                 MetaName = "robots"
	MetaNameGoogleSiteVerification MetaName = "google-site-verification"
)

func (m *meta) Name(name MetaName) *meta {
	m.el.AddAttribute("name", string(name))
	return m
}

func (m *meta) Element() *HtmlElement { return m.el }
func (m *meta) Render() string        { return m.el.Render() }

func Meta() *meta {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "meta",
		},
	}

	m := &meta{el: el}

	return m
}
