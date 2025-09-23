package element

import (
	"strings"

	"golang.org/x/net/html"
)

type RefererPolicy string

const (
	RefererPolicyNoReferer                   RefererPolicy = "no-referrer"
	RefererPolicyNoRefererWhenDowngrade      RefererPolicy = "no-referrer-when-downgrade"
	RefererPolicyOrigin                      RefererPolicy = "origin"
	RefererPolicyOriginWhenCrossOrigin       RefererPolicy = "origin-when-cross-origin"
	RefererPolicySameOrigin                  RefererPolicy = "same-origin"
	RefererPolicyStrictOriginWhenCrossOrigin RefererPolicy = "strict-origin-when-cross-origin"
	RefererPolicyUnsafeUrl                   RefererPolicy = "unsafe-url"
)

type AnchorRel string

const (
	Alternate  AnchorRel = "alternate"  // Provides a link to an alternate representation of the document (i.e. print page, translated or mirror)
	Author     AnchorRel = "author"     // Provides a link to the author of the document
	Bookmark   AnchorRel = "bookmark"   // Permanent URL used for bookmarking
	External   AnchorRel = "external"   // Indicates that the referenced document is not part of the same site as the current document
	Help       AnchorRel = "help"       // Provides a link to a help document
	License    AnchorRel = "license"    // Provides a link to licensing information for the document
	Next       AnchorRel = "next"       // Provides a link to the next document in the series
	Nofollow   AnchorRel = "nofollow"   // Links to an unendorsed document, like a paid link. ("nofollow" is used by Google, to specify that the Google search spider should not follow that link)
	Noopener   AnchorRel = "noopener"   // Requires that any browsing context created by following the hyperlink must not have an opener browsing context
	Noreferrer AnchorRel = "noreferrer" // Makes the referrer unknown. No referer header will be included when the user clicks the hyperlink
	Prev       AnchorRel = "prev"       // The previous document in a selection
	SearchRel  AnchorRel = "search"     // Links to a search tool for the document
	Tag        AnchorRel = "tag"        // A tag (keyword) for the current document
)

type AnchorTarget string

const (
	TargetBlank  AnchorTarget = "_blank"  //	Opens the linked document in a new window or tab
	TargetSelf   AnchorTarget = "_self"   //	Opens the linked document in the same frame as it was clicked (this is default)
	TargetParent AnchorTarget = "_parent" //	Opens the linked document in the parent frame
	TargetTop    AnchorTarget = "_top"    //	Opens the linked document in the full body of the window
)

type anchor struct {
	el *HtmlElement
}

func (a *anchor) Download(filename string) *anchor {
	a.el.AddAttribute("download", filename)
	return a
}

func (a *anchor) Href(url string) *anchor {
	a.el.AddAttribute("href", url)
	return a
}

func (a *anchor) HrefLang(code LanguageCode) *anchor {
	a.el.AddAttribute("hreflang", string(code))
	return a
}

func (a *anchor) Media(media string) *anchor {
	a.el.AddAttribute("media", media)
	return a
}

func (a *anchor) Ping(urls ...string) *anchor {
	a.el.AddAttribute("ping", strings.Join(urls, ", "))
	return a
}

func (a *anchor) ReferrerPolicy(policy RefererPolicy) *anchor {
	a.el.AddAttribute("referrerpolicy", string(policy))
	return a
}

func (a *anchor) Rel(relationship AnchorRel) *anchor {
	a.el.AddAttribute("rel", string(relationship))
	return a
}

func (a *anchor) Target(target AnchorTarget) *anchor {
	a.el.AddAttribute("target", string(target))
	return a
}
func (a *anchor) Type(mediatype string) *anchor {
	a.el.AddAttribute("type", mediatype)
	return a
}

func (a *anchor) Children(children ...IsElement) *anchor {
	a.el.Children(children...)
	return a
}

func (a *anchor) Component()            {}
func (a *anchor) Element() *HtmlElement { return a.el }
func (a *anchor) Render() string        { return a.el.Render() }

func A(children ...IsElement) *anchor {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "a",
		},
	}

	el.Children(children...)

	a := &anchor{el: el}

	return a
}
