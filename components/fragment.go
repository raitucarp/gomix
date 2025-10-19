package components

import (
	"golang.org/x/net/html"
)

func removeFragment(n *html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		removeFragment(c)
	}

	if n.Type == html.ElementNode && n.Data == "fragment" {
		parent := n.Parent
		if parent == nil {
			return
		}

		for child := n.FirstChild; child != nil; {
			next := child.NextSibling
			n.RemoveChild(child)
			parent.InsertBefore(child, n)
			child = next
		}

		parent.RemoveChild(n)
	}
}
