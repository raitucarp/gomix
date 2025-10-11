package gomix

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/theme"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
)

// type PageComponent func(page *Page) components.IsComponent
type PageSetup func(page *Page)
type PageParam func(page *Page)
type pageComponent func(page *Page) components.IsComponent
type PageParams []PageParam

type metadataField string

type Page struct {
	path      LocationPath
	title     string
	component pageComponent
	fragment  pageComponent
	layouts   []pageComponent
	children  []*Page
	metadata  map[metadataField]any

	scripts     []string
	stylesheets []string
	css         string
	theme       *theme.Theme

	flattened bool
	request   *http.Request
	response  http.ResponseWriter
	handler   func(w http.ResponseWriter, req *http.Request)
}

func newPage(path LocationPath) *Page {
	newPage := &Page{path: path,
		component: func(page *Page) components.IsComponent {
			return element.Element(element.Text(""))
		},
		request:  &http.Request{},
		metadata: map[metadataField]any{},
	}

	newPage.handler = func(w http.ResponseWriter, req *http.Request) {
		newPage.setReq(req)
		newPage.setResp(w)
		// setup(newPage)
		newPage.write()
	}

	return newPage
}

func (page *Page) addLayouts(layout ...pageComponent) {
	page.layouts = append(page.layouts, layout...)
}

func (page *Page) addScripts(urls ...string) {
	page.scripts = append(page.scripts, urls...)
}

func (page *Page) addStylesheets(urls ...string) {
	page.scripts = append(page.stylesheets, urls...)
}

func (page *Page) Theme(t *theme.Theme) {
	page.theme = t
}

func (page *Page) Inherit(path LocationPath, params ...PageParam) *Page {
	newPage := newPage(path)
	newPage.Theme(page.theme)

	return newPage
}

type Pages []*Page

func (page *Page) setReq(req *http.Request) {
	page.request = req
}

func (page *Page) setResp(res http.ResponseWriter) {
	page.response = res
}

func (page *Page) setup() {

}

func (page *Page) Request() *http.Request {
	return page.request
}

func (page *Page) Response() http.ResponseWriter {
	page.response.Header().Set("Content-Type", "text/html")
	return page.response
}

func (page *Page) write() {
	io.WriteString(page.response, page.Render("en"))
}

func (page *Page) flattenPages() (pages []*Page) {

	for _, p := range page.children {
		p.flattened = true
		p.addLayouts(page.layouts...)
		p.addLayouts(p.component)
		p.addStylesheets(page.stylesheets...)
		p.addScripts(page.scripts...)
		p.css = page.css + p.css
		if p.theme == nil {
			p.theme = page.theme
		}

		pages = append(pages, p)
		if len(p.children) > 0 {
			pages = append(pages, p.flattenPages()...)
		}
	}

	return
}

func (page *Page) Children(pages Pages) {
	if page.flattened {
		return
	}
	page.children = append(page.children, pages...)
}

func (page *Page) Render(lang element.LanguageCode) string {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)

	theLayout := components.Component(components.Slot())
	for _, pageLayout := range page.layouts {
		appliedLayout := pageLayout(page)
		theLayout = components.ApplyLayout(theLayout, appliedLayout)
	}

	head := []element.IsHeadElement{
		element.Meta().CharSet("UTF-8"),
		element.Meta().Name(element.MetaNameViewport).Content("width=device-width, initial-scale=1.0"),
		element.Title(page.title),
	}

	for _, stylesheet := range page.stylesheets {
		head = append(head, element.Link().Href(stylesheet).Rel("stylesheet"))
	}

	allStylesheet := []string{}
	themeVars, err := m.String("text/css", page.theme.RawCSS())
	if err == nil {
		allStylesheet = append(allStylesheet, themeVars)
	}

	globalCss, err := m.String("text/css", page.css)
	if err == nil {
		allStylesheet = append(allStylesheet, globalCss)
	}

	pageElementStyles := components.ExtractCSS(theLayout)
	pageElementStyles, err = m.String("text/css", pageElementStyles)
	if err == nil {
		allStylesheet = append(allStylesheet, pageElementStyles)
	}
	head = append(head, element.Style(strings.Join(allStylesheet, "\n")))

	for _, script := range page.scripts {
		head = append(head, element.Script().Src(script))
	}

	layout := components.Component(
		element.Html(
			element.Head(head...),
			element.Body(
				theLayout,
			),
		).Lang(lang),
	)

	finalHtml, _ := m.String("text/html", components.Render(layout))
	finalHtml = fmt.Sprintf("<!DOCTYPE html>%s", finalHtml)

	return finalHtml
}

func PageTitle(title string) AppParam {
	return func(app *Application) (scope Scope, fn func(params ...any)) {
		return PageScope, func(params ...any) {
			if len(params) > 0 {
				if page, ok := params[0].(*Page); ok {
					page.title = title
				}
			}

		}
	}
}

func Metadata(metaParams ...AppParam) AppParam {
	return func(app *Application) (scope Scope, fn func(params ...any)) {
		return PageScope, func(params ...any) {

		}
	}
}

func PageComponent(component pageComponent) AppParam {
	return func(app *Application) (scope Scope, fn func(params ...any)) {
		return PageScope, func(params ...any) {
			if len(params) > 0 {
				if page, ok := params[0].(*Page); ok {
					page.component = component
				}
			}

		}
	}
}

func PageFragmentComponent(component pageComponent) AppParam {
	return func(app *Application) (scope Scope, fn func(params ...any)) {
		return PageScope, func(params ...any) {
			if len(params) > 0 {
				if page, ok := params[0].(*Page); ok {
					page.fragment = component
				}
			}

		}
	}
}
