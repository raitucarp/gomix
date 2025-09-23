package gomix

import (
	"io"
	"net/http"

	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type PageComponent func(page *Page) components.IsComponent
type PageSetup func(page *Page)

type Page struct {
	path      LocationPath
	title     string
	component PageComponent
	layouts   []PageComponent
	children  []*Page

	scripts     []string
	stylesheets []string

	flattened bool
	request   *http.Request
	response  http.ResponseWriter
	handler   func(w http.ResponseWriter, req *http.Request)
}

func (app *App) Page(path LocationPath, setup PageSetup) *Page {
	newPage := NewPage(path, setup)
	app.pages = append(app.pages, newPage)

	return newPage
}

func NewPage(path LocationPath, setup PageSetup) *Page {
	newPage := &Page{path: path,

		component: func(page *Page) components.IsComponent {
			return element.Element(element.Text(""))
		},
		request: &http.Request{},
	}

	newPage.handler = func(w http.ResponseWriter, req *http.Request) {
		newPage.setReq(req)
		newPage.setResp(w)

		setup(newPage)
		newPage.Write()
	}

	setup(newPage)

	return newPage
}

func (page *Page) AddLayouts(layout ...PageComponent) {
	page.layouts = append(page.layouts, layout...)
}

func (page *Page) AddScripts(urls ...string) {
	page.scripts = append(page.scripts, urls...)
}

func (page *Page) New(path LocationPath, callback PageSetup) *Page {
	newPage := NewPage(path, callback)

	return newPage
}

type Pages []*Page

func (page *Page) Title(title string) {
	page.title = title
}

func (page *Page) Component(component PageComponent) {
	page.component = component
}

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

func (page *Page) Write() {
	io.WriteString(page.response, page.Render("en"))
}

func (page *Page) flattenPages() (pages []*Page) {

	page.AddLayouts(page.component)

	for _, p := range page.children {
		p.flattened = true
		p.AddLayouts(page.layouts...)
		p.AddScripts(page.scripts...)

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
	head := []element.IsHeadElement{
		element.Meta().CharSet("UTF-8"),
		element.Meta().Name(element.MetaNameViewport).Content("width=device-width, initial-scale=1.0"),
		element.Title(page.title),
	}

	for _, script := range page.scripts {
		head = append(head, element.Script().Src(script))
	}

	layout := components.Component(
		element.Html(
			element.Head(head...),
			element.Body(
				element.Element(
					components.Slot(),
				),
			),
		).Lang(lang),
	)

	for _, pageLayout := range page.layouts {
		c := pageLayout(page)
		layout = components.ApplyLayout(layout, c)
	}

	return components.Render(layout)
}
