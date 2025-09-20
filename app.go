package gomix

import (
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type App struct {
	name      string
	port      int
	layout    components.IsComponent
	pages     []*Page
	fragments []*Fragment

	scripts      []string
	stylesheets  []string
	enableLogger bool
}

type LocationPath string

func New(name string) *App {
	return &App{
		name: name,
		layout: components.Component(
			element.Body(
				element.Element(components.Slot()),
			),
		),
	}
}

func (app *App) flattenPages() (pages []*Page) {
	for _, page := range app.pages {
		page.flattened = true
		page.AddLayouts(func(page *Page) components.IsComponent { return app.layout })
		page.AddScripts(app.scripts...)

		pages = append(pages, page)
		if len(page.children) > 0 {
			pages = append(pages, page.flattenPages()...)
		}

	}
	return
}

func (app *App) Scripts(url ...string) {
	app.scripts = append(app.scripts, url...)
	app.scripts = slices.Compact(app.scripts)
}

func (app *App) EnableLogger() {
	app.enableLogger = true
}

func (app *App) Layout(component components.IsComponent) {
	app.layout = components.Component(
		element.Body(
			element.Element(component),
		),
	)
}

func (app *App) Addons(addons ...func(app *App)) {
	for _, addon := range addons {
		addon(app)
	}
}

func (app *App) ServeAtPort(port int) {
	app.port = port

	mux := http.NewServeMux()

	portString := fmt.Sprintf(":%d", port)

	allPages := app.flattenPages()

	for _, page := range allPages {
		mux.HandleFunc(string(page.path), page.handler)
	}

	for _, fragment := range app.fragments {
		mux.HandleFunc(string(fragment.path), fragment.handler)
	}

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{})
	})

	logString := fmt.Sprintf("Server %s listening on %s", app.name, portString)

	log.Println(logString)
	http.ListenAndServe(portString, mux)
}
