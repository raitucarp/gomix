package htmx

import (
	"fmt"

	"github.com/raitucarp/gomix"
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

const AddonName = "htmx"

func Addon(version string) gomix.AppParam {
	scope := gomix.WebScope
	return func(app *gomix.Application) (gomix.Scope, func()) {
		return scope, func() {
			app.InstallAddon(AddonName)
			htmxUrl := fmt.Sprintf("https://cdn.jsdelivr.net/npm/htmx.org@%s/dist/htmx.min.js", version)

			app.Apply(scope, gomix.Scripts(htmxUrl))
		}
	}
}

type htmx struct {
	el *element.HtmlElement
}

func (h *htmx) Element() *element.HtmlElement { return h.el }
func (h *htmx) Render() string                { return h.el.Render() }

func (h *htmx) Get(url string) *htmx {
	h.el.AddAttribute("hx-get", url)
	return h
}

func (h *htmx) Target(id string) *htmx {
	h.el.AddAttribute("hx-target", id)
	return h
}

func Htmx(e components.IsComponent) *htmx {
	h := htmx{
		el: e.Element(),
	}
	return &h
}
