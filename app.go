package gomix

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/gorilla/mux"
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/theme"
)

type Application struct {
	name string
	port int
	host string

	features []string
	addons   []string

	web *webPage

	enableLogger bool
}

type LocationPath string
type AppParam func(app *Application) (scope Scope, fn func(params ...any))

func App(params ...AppParam) {
	app := &Application{
		host: "0.0.0.0",
		port: 3000,
		web: &webPage{
			layout: components.Component(
				element.Body(
					element.Element(components.Slot()),
				),
			),
			theme: theme.Default,
			pages: []*Page{},
		},
	}

	app.web.pages = append(app.web.pages, app.notFoundPage())

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

		p.applyTitle(app.web.title, app.web.titleTemplate)

		layouts := []pageComponent{}
		layouts = append(layouts, app.defaultLayout())
		if p.layout != nil {
			layouts = append(layouts, p.layout)
		}
		layouts = append(layouts, p.component)

		p.addLayouts(layouts...)
		p.addStylesheets(app.web.stylesheets...)
		p.addScripts(app.web.scripts...)
		p.css = app.web.css + p.css
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

func (app *Application) notFoundPage() *Page {
	notFoundPage := newPage("/404")
	notFoundPage.kind = pageNotFound
	notFoundPage.title = "404 Page Not Found"
	notFoundPage.component = notFoundPageComponent
	return notFoundPage
}

func (app *Application) favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}

func (app *Application) serve() {
	r := mux.NewRouter()

	allPages := app.flattenPages()

	for _, page := range allPages {
		switch page.kind {
		case pageNotFound:
			r.NotFoundHandler = http.HandlerFunc(page.handler)
			continue
		case pageError:
			r.NotFoundHandler = http.HandlerFunc(page.handler)
			continue
		case pageNormal:
			r.HandleFunc(string(page.path), page.handler)
			continue
		}
	}

	for _, fragment := range app.web.fragments {
		r.HandleFunc(string(fragment.path), fragment.handler)
	}

	r.HandleFunc("/favicon.ico", app.favicon)

	logString := fmt.Sprintf("Server %s listening on %d", app.name, app.port)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", app.host, app.port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println(logString)
	log.Fatal(srv.ListenAndServe())
}
