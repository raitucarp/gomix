package gomix

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/theme"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/valyala/fasttemplate"
)

// type PageComponent func(page *Page) components.IsComponent
type PageSetup func(page *Page)
type PageParam func(page *Page)
type pageComponent func(page *Page) components.IsComponent
type PageParams []PageParam

type metadataField string

type pageKind string

const (
	pageNormal   pageKind = "normal"
	pageNotFound pageKind = "not_found"
	pageError    pageKind = "error"
)

type Page struct {
	path          LocationPath
	title         string
	titleTemplate string
	component     pageComponent
	fragment      pageComponent
	layouts       []pageComponent
	layout        pageComponent
	children      []*Page
	metadata      map[metadataField]any
	kind          pageKind

	scripts     []script
	stylesheets []string
	css         string
	theme       *theme.Theme

	flattened bool
	request   *http.Request
	response  http.ResponseWriter
	params    map[string]any
}

func newPage(path LocationPath) *Page {
	newPage := &Page{path: path,
		component: func(page *Page) components.IsComponent {
			return element.Element(element.Text(""))
		},
		request:  &http.Request{},
		metadata: map[metadataField]any{},
		kind:     pageNormal,
		params:   map[string]any{},
	}

	return newPage
}

func notFoundPageComponent(page *Page) components.IsComponent {
	return components.VStack(components.Text("Not Found")).Component().
		WScreen().
		HScreen().
		JustifyCenter().ItemsCenter()
}

func (page *Page) handler(w http.ResponseWriter, req *http.Request) {
	page.setReq(req)
	for key, value := range mux.Vars(req) {
		page.params[key] = value
	}
	page.setResp(w)
	page.write()
}

func (page *Page) addLayouts(layout ...pageComponent) {
	page.layouts = append(page.layouts, layout...)
}

func (page *Page) addScripts(scripts ...script) {
	page.scripts = append(page.scripts, scripts...)
}

func (page *Page) addStylesheets(urls ...string) {
	page.stylesheets = append(page.stylesheets, urls...)
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

func (page *Page) applyTitle(defaultTitle string, titleTemplate string) {
	if page.title == "" {
		page.title = defaultTitle
	}

	if titleTemplate != "" {
		t := fasttemplate.New(titleTemplate, "{", "}")
		page.title = t.ExecuteString(map[string]any{
			"title": page.title,
		})
	}
	page.titleTemplate = titleTemplate
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

		p.applyTitle(page.title, page.titleTemplate)

		layouts := []pageComponent{}
		if page.layout != nil && page.component != nil {
			layouts = append(layouts, page.layouts[:len(page.layouts)-1]...)
		} else {
			layouts = append(layouts, page.layouts...)
		}

		if p.layout != nil {
			layouts = append(layouts, p.layout)
		}
		layouts = append(layouts, p.component)

		p.addLayouts(layouts...)
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
		element.Title(
			fasttemplate.New(page.title, "{", "}").ExecuteString(page.params),
		),
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

	for _, s := range page.scripts {
		if s.kind == scriptUrl {
			if s.isDefer {
				head = append(head, element.Script().Src(s.data).Defer())
			} else {
				head = append(head, element.Script().Src(s.data))
			}
		}
	}

	for _, s := range page.scripts {
		if s.kind == scriptRaw {
			head = append(head, element.Script().Content(s.data))
		}
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

func PageLayout(layout pageComponent) AppParam {
	return func(app *Application) (scope Scope, fn func(params ...any)) {
		return PageScope, func(params ...any) {
			if len(params) > 0 {
				if page, ok := params[0].(*Page); ok {
					page.layout = layout
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
