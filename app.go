package gomix

import (
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/theme"
)

type Application struct {
	name string
	port int

	features []string
	addons   []string

	web *webPage

	enableLogger bool
}

type LocationPath string
type AppParam func(app *Application) (scope Scope, fn func(params ...any))

func App(params ...AppParam) {
	app := &Application{
		port: -1,
		web: &webPage{
			layout: components.Component(
				element.Body(
					element.Element(components.Slot()),
				),
			),
			theme: theme.Default,
		},
	}
	for _, param := range params {
		scope, runFn := param(app)
		if scope == AppScope {
			runFn()
		}
	}

	if app.port > 0 {
		app.serve()
	}
}

func Name(name string) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return AppScope, func(params ...any) {
			app.name = name
		}
	}
}

func Addons(addons ...AppParam) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return AppScope, func(params ...any) {
			for _, addon := range addons {
				scope, runFn := addon(app)
				if scope == AppScope {
					runFn()
				}
			}
		}
	}
}

func Features(features ...AppParam) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return AppScope, func(params ...any) {
			for _, feature := range features {
				scope, runFn := feature(app)
				if scope == AppScope {
					runFn()
				}
			}
		}
	}
}

func Web(features ...AppParam) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return AppScope, func(params ...any) {
			for _, fn := range features {
				scope, runFn := fn(app)
				if scope == WebScope {
					runFn()
				}
			}
		}
	}
}

func Port(port int) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return AppScope, func(params ...any) {
			app.port = port
		}
	}
}

func Logger() AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return AppScope, func(params ...any) {
			app.features = append(app.features, "logger")
		}
	}
}

func Rest(apis ...AppParam) AppParam {
	return func(app *Application) (Scope, func(params ...any)) {
		return AppScope, func(params ...any) {
			for _, aa := range apis {
				aa(app)
			}
		}
	}
}

func (app *Application) Apply(scope Scope, param AppParam) {
	paramScope, runFn := param(app)
	if scope == paramScope {
		runFn()
	}
}

func (app *Application) InstallAddon(addonName string) {
	log.Println("Install", addonName)
	app.addons = append(app.addons, addonName)
}

func (app *Application) IsAddonActivated(addonName string) bool {
	return slices.Contains(app.addons, addonName)
}

func (app *Application) defaultLayout() pageComponent {
	return func(page *Page) components.IsComponent { return app.web.layout }
}

func (app *Application) flattenPages() (pages []*Page) {
	for _, p := range app.web.pages {
		p.flattened = true
		p.addLayouts(app.defaultLayout(), p.component)
		p.addScripts(app.web.scripts...)
		if p.theme == nil {
			p.theme = app.web.theme
		}

		pages = append(pages, p)
		if len(p.children) > 0 {
			pages = append(pages, p.flattenPages()...)
		}
	}
	return
}

func (app *Application) findPageByPath(path LocationPath) *Page {
	for _, page := range app.web.pages {
		if page.path == path {
			return page
		}
	}
	return nil
}

func (app *Application) serve() {
	mux := http.NewServeMux()
	portString := fmt.Sprintf(":%d", app.port)

	allPages := app.flattenPages()

	for _, page := range allPages {
		mux.HandleFunc(string(page.path), page.handler)
	}

	for _, fragment := range app.web.fragments {
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
