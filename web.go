package gomix

import (
	"slices"

	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/theme"
)

type webPage struct {
	prefixTitle string
	suffixTitle string
	layout      components.IsComponent
	pages       []*Page
	fragments   []*Fragment
	theme       *theme.Theme
	scripts     []string
	stylesheets []string
}

func WebAddons(addons ...AppParam) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return WebScope, func(params ...any) {
			for _, addon := range addons {
				scope, runFn := addon(app)
				if scope == WebScope {
					runFn()
				}
			}
		}
	}
}

func Scripts(url ...string) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return WebScope, func(params ...any) {
			app.web.scripts = append(app.web.scripts, url...)
			app.web.scripts = slices.Compact(app.web.scripts)
		}
	}
}

func PageAt(path LocationPath, params ...AppParam) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {

		return WebScope, func(p ...any) {
			if len(p) <= 0 {
				newPage := newPage(path)
				app.web.pages = append(app.web.pages, newPage)

				for _, param := range params {
					scope, fn := param(app)
					if scope == PageScope {
						fn(newPage)
					}

					if scope == WebScope {
						fn(newPage)
					}
				}
				return
			}

			if len(p) > 0 {
				if page, ok := p[0].(*Page); ok {
					newPage := newPage(path)
					page.children = append(page.children, newPage)

					for _, param := range params {

						scope, fn := param(app)
						if scope == PageScope {
							fn(newPage)
						}

						if scope == WebScope {
							fn(newPage)
						}

					}

					return
				}
			}

		}
	}
}

func Theme(newTheme *theme.Theme) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return WebScope, func(params ...any) {
			app.web.theme = newTheme
		}
	}
}

func FragmentAt(path FragmentPath, params ...AppParam) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return WebScope, func(p ...any) {
			fragment := newFragment(path)
			app.web.fragments = append(app.web.fragments, fragment)
			for _, param := range params {
				scope, fn := param(app)
				if scope == FragmentScope {
					fn(fragment)
				}
			}
		}
	}
}

func Layout(component components.IsComponent) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return WebScope, func(params ...any) {
			app.web.layout = components.Component(
				element.Body(
					element.Element(component),
				),
			)
		}
	}
}
