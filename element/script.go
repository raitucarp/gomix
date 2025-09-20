package element

import "golang.org/x/net/html"

type CrossOrigin string

const (
	CrossOriginAnonymous      CrossOrigin = "anonymous"
	CrossOriginUseCredentials CrossOrigin = "use-credentials"
)

type script struct {
	el *HtmlElement
}

func (s *script) InsideHead()           {}
func (s *script) Element() *HtmlElement { return s.el }
func (s *script) Render() string        { return s.el.Render() }

func (s *script) Async() *script {
	s.el.AddAttributeBool("async")
	return s
}

func (s *script) CrossOrigin(crossOrigin CrossOrigin) *script {
	s.el.AddAttribute("crossorigin", string(crossOrigin))
	return s
}

func (s *script) Defer() *script {
	s.el.AddAttributeBool("defer")
	return s
}

func (s *script) Integrity(filehash string) *script {
	s.el.AddAttribute("integrity", filehash)
	return s
}

func (s *script) Src(url string) *script {
	s.el.AddAttribute("src", url)
	return s
}

func (s *script) NoModule(nomodule bool) *script {
	if nomodule {
		s.el.AddAttribute("NoModule", "True")
	} else {
		s.el.AddAttribute("NoModule", "False")
	}
	return s
}

func (s *script) ReferrerPolicy(policy RefererPolicy) *script {
	s.el.AddAttribute("referrer", string(policy))
	return s
}

func (s *script) Content(content string) *script {
	s.el.node.FirstChild = &html.Node{
		Type: html.TextNode,
		Data: content,
	}
	return s
}

func (s *script) Type(scripttype string) *script {
	s.el.AddAttribute("type", scripttype)
	return s
}

func Script() *script {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "script",
		},
	}

	a := &script{el: el}

	return a
}
