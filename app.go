package gomix

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/raitucarp/gomix/components"
)

type App struct {
	name      string
	port      int
	layout    *components.Component
	pages     []*Page
	fragments []*Fragment
	ctx       context.Context

	scripts      []string
	stylesheets  []string
	enableLogger bool
}

type LocationPath string

func New(ctx context.Context, name string) *App {
	return &App{
		name:   name,
		ctx:    ctx,
		layout: components.Body(components.Slot()).Component(),
	}
}

func (app *App) flattenPages() (pages []*Page) {
	for _, page := range app.pages {
		page.flattened = true
		page.AddLayouts(func(page *Page) *components.Component { return app.layout })
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

func (app *App) Layout(component *components.Component) {
	app.layout = components.Body(component).Component()
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
