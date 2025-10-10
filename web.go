package gomix

import (
	"fmt"
	"slices"
	"strings"

	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/styles"
	"github.com/raitucarp/gomix/theme"
)

type webPage struct {
	title       string
	layout      components.IsComponent
	pages       []*Page
	fragments   []*Fragment
	theme       *theme.Theme
	scripts     []string
	stylesheets []string
	css         string
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

type CSS map[string][]styles.ApplyProp

func Style(props ...styles.ApplyProp) []styles.ApplyProp {
	return props
}

func StyleGlobal(css CSS) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return WebScope, func(params ...any) {
			combinedCss := []string{}
			variantClassMap := map[string][]string{}
			variantPropsMap := map[string]map[string]string{}

			for rule, props := range css {
				style := styles.ApplyStyle(app.web.theme, props...)

				classes := strings.SplitSeq(rule, ",")

				for className := range classes {
					className = strings.Trim(className, " ")
					for variant, props := range style {
						variantStyleName := styles.VariantNameOfAttr(string(variant)).Name()

						if variantPropsMap[variantStyleName] == nil {
							variantPropsMap[variantStyleName] = make(map[string]string)
						}

						style := []string{}
						for key, p := range props {
							style = append(style, fmt.Sprintf("%s:%s", key, p))
						}

						if len(style) > 0 {
							variantPropsMap[variantStyleName][className] = strings.Join(style, ";")
							if !slices.Contains(variantClassMap[variantStyleName], className) {

								variantClassMap[variantStyleName] = append(
									variantClassMap[variantStyleName], className,
								)
							}
						}

					}
				}
			}

			combinedCss = styles.ExtractCSSFromStyle(variantClassMap, variantPropsMap)

			app.web.css = strings.Join(combinedCss, "\n")
		}
	}
}
