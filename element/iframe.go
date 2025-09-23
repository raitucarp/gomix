package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type iframe struct {
	el *HtmlElement
}

func (i *iframe) Allow(v string) *iframe {
	i.el.AddAttribute("allow", v)
	return i
}

func (i *iframe) AllowFullscreen() *iframe {
	i.el.AddAttributeBool("allowfullscreen")
	return i
}

func (i *iframe) AllowPaymentRequest() *iframe {
	i.el.AddAttributeBool("allowpaymentrequest")
	return i
}

func (i *iframe) Height(height int) *iframe {
	i.el.AddAttribute("height", strconv.Itoa(height))
	return i
}

type Loading string

const (
	LoadingEager Loading = "eager"
	LoadingLazy  Loading = "lazy"
)

func (i *iframe) Loading(loading Loading) *iframe {
	i.el.AddAttribute("loading", string(loading))
	return i
}

func (i *iframe) Name(name string) *iframe {
	i.el.AddAttribute("name", name)
	return i
}

type IframeReferrerPolicy string

const (
	IframeReferrerPolicyNoReferrer                  IframeReferrerPolicy = "no-referrer"
	IframeReferrerPolicyNoReferrerWhenDowngrade     IframeReferrerPolicy = "no-referrer-when-downgrade"
	IframeReferrerPolicyOrigin                      IframeReferrerPolicy = "origin"
	IframeReferrerPolicyOriginWhenCrossOrigin       IframeReferrerPolicy = "origin-when-cross-origin"
	IframeReferrerPolicySameOrigin                  IframeReferrerPolicy = "same-origin"
	IframeReferrerPolicyStrictOrigin                IframeReferrerPolicy = "strict-origin"
	IframeReferrerPolicyStrictOriginWhenCrossOrigin IframeReferrerPolicy = "strict-origin-when-cross-origin"
	IframeReferrerPolicyUnsafeUrl                   IframeReferrerPolicy = "unsafe-url"
)

func (i *iframe) ReferrerPolicy(policy IframeReferrerPolicy) *iframe {
	i.el.AddAttribute("referrerpolicy", string(policy))
	return i
}

type IframeSandbox string

const (
	IframeSandboxAllowForms                         IframeSandbox = "allow-forms"
	IframeSandboxAllowModals                        IframeSandbox = "allow-modals"
	IframeSandboxAllowOrientationLock               IframeSandbox = "allow-orientation-lock"
	IframeSandboxAllowPointerLock                   IframeSandbox = "allow-pointer-lock"
	IframeSandboxAllowPopups                        IframeSandbox = "allow-popups"
	IframeSandboxAllowPopupsToEscapeSandbox         IframeSandbox = "allow-popups-to-escape-sandbox"
	IframeSandboxAllowPresentation                  IframeSandbox = "allow-presentation"
	IframeSandboxAllowSameOrigin                    IframeSandbox = "allow-same-origin"
	IframeSandboxAllowScripts                       IframeSandbox = "allow-scripts"
	IframeSandboxAllowTopNavigation                 IframeSandbox = "allow-top-navigation"
	IframeSandboxAllowTopNavigationByUserActivation IframeSandbox = "allow-top-navigation-by-user-activation"
)

func (i *iframe) Sandbox(sandbox IframeSandbox) *iframe {
	i.el.AddAttribute("sandbox", string(sandbox))
	return i
}

func (i *iframe) Src(url string) *iframe {
	i.el.AddAttribute("src", url)
	return i
}

func (i *iframe) SrcDoc(htmlSrc IsElement) *iframe {
	i.el.AddAttribute("srcdoc", htmlSrc.Element().Render())
	return i
}

func (i *iframe) Width(width int) *iframe {
	i.el.AddAttribute("width", strconv.Itoa(width))
	return i
}

func (i *iframe) Element() *HtmlElement { return i.el }
func (i *iframe) Render() string        { return i.el.Render() }

func Iframe() *iframe {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "iframe",
		},
	}

	i := &iframe{el: el}

	return i
}
