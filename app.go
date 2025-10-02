package gomix

import (
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type web struct {
	prefixTitle string
	suffixTitle string
}

type Application struct {
	name string
	port int

	features []string
	addons   []string

	layout    components.IsComponent
	pages     []*Page
	fragments []*Fragment

	scripts      []string
	stylesheets  []string
	enableLogger bool
}

type LocationPath string
type AppParam func(app *Application)

func App(params ...AppParam) {
	app := &Application{
		port: -1,
		layout: components.Component(
			element.Body(
				element.Element(components.Slot()),
			),
		),
	}
	for _, param := range params {
		param(app)
	}

	if app.port > 0 {
		app.serve()
	}
}

func Name(name string) AppParam {
	return func(app *Application) {
		app.name = name
	}
}

func Addons(addons ...AppParam) AppParam {
	return func(app *Application) {
		for _, addon := range addons {
			addon(app)
		}
	}
}

func Features(features ...AppParam) AppParam {
	return func(app *Application) {
		for _, feature := range features {
			feature(app)
		}
	}
}

func Webpage(features ...AppParam) AppParam {
	return func(app *Application) {
		for _, feature := range features {
			feature(app)
		}
	}
}

func Port(port int) AppParam {
	return func(app *Application) {
		app.port = port
	}
}

func Logger() AppParam {
	return func(app *Application) {
		app.features = append(app.features, "logger")
	}
}

func Rest(apis ...AppParam) AppParam {
	return func(app *Application) {
		for _, aa := range apis {
			aa(app)
		}
	}
}

func Scripts(url ...string) AppParam {
	return func(app *Application) {
		app.scripts = append(app.scripts, url...)
		app.scripts = slices.Compact(app.scripts)
	}
}

func PageAt(path LocationPath, setup PageSetup) AppParam {
	return func(app *Application) {
		newPage := NewPage(path, setup)
		app.pages = append(app.pages, newPage)
	}
}

func FragmentAt(path FragmentPath, component FragmentComponent) AppParam {
	return func(app *Application) {
		fragment := NewFragment(path, component)
		app.fragments = append(app.fragments, fragment)
	}
}

func Layout(component components.IsComponent) AppParam {
	return func(app *Application) {
		app.layout = components.Component(
			element.Body(
				element.Element(component),
			),
		)
	}
}

func (app *Application) Apply(param AppParam) {
	param(app)
}

func (app *Application) InstallAddon(addonName string) {
	log.Println("Install", addonName)
	app.addons = append(app.addons, addonName)
}

func (app *Application) IsAddonActivated(addonName string) bool {
	return slices.Contains(app.addons, addonName)
}

func (app *Application) flattenPages() (pages []*Page) {
	for _, p := range app.pages {
		p.flattened = true
		p.AddLayouts(func(page *Page) components.IsComponent { return app.layout })
		p.AddScripts(app.scripts...)

		pages = append(pages, p)
		if len(p.children) > 0 {
			pages = append(pages, p.flattenPages()...)
		}

	}
	return
}

func (app *Application) serve() {
	mux := http.NewServeMux()
	portString := fmt.Sprintf(":%d", app.port)

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
