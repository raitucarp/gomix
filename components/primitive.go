package components

import "golang.org/x/net/html"

func TextRaw(text string) *Component {
	return &Component{
		node: &html.Node{
			Type: html.TextNode,
			Data: text,
		},
	}
}

func Html(lang string) *Component {
	return &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "html",
			Attr: []html.Attribute{
				{Key: "lang", Val: lang},
			},
		},
	}
}

func Head() *Component {
	return &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "head",
		},
	}
}

func Title(title string) *Component {
	return &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "title",
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: title,
			},
		},
	}
}

func Meta(attribs map[string]string) *Component {
	attrs := []html.Attribute{}
	for k, v := range attribs {
		attrs = append(attrs, html.Attribute{Key: k, Val: v})
	}
	return &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "meta",
			Attr: attrs,
		},
	}
}

func Stylesheet(href string) *Component {
	return &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "link",
			Attr: []html.Attribute{
				{Key: "rel", Val: "stylesheet"},
				{Key: "href", Val: href},
			},
		},
	}
}

func Style(css string) *Component {
	return &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "style",
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: css,
			},
		},
	}
}

func Script(src string) *Component {
	return &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "script",
			Attr: []html.Attribute{
				{Key: "type", Val: "text/javascript"},
				{Key: "src", Val: src},
			},
		},
	}
}

func ScriptRaw(script string) *Component {
	return &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "script",
			Attr: []html.Attribute{
				{Key: "type", Val: "text/javascript"},
			},
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: script,
			},
		},
	}
}


func Div(children ...*Component) *Component {
	div := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "div",
		},
	}

	div.Children(children...)

	return div
}
