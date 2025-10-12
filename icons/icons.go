package icons

import "github.com/raitucarp/gomix/element"

type IsIcon interface {
	IsIcon()
	Element() *element.HtmlElement
}
