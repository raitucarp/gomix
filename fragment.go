package gomix

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/raitucarp/gomix/components"
	"golang.org/x/net/html"
)

type FragmentPath string

type Fragment struct {
	path      FragmentPath
	component fragmentComponent

	request  *http.Request
	response http.ResponseWriter
	handler  func(w http.ResponseWriter, req *http.Request)
}

type FragmentParam func(fragment *Fragment)
type fragmentComponent func(fragment *Fragment) components.IsComponent

// func (fragment *Fragment) Write() {
// 	io.WriteString(fragment.response, fragment.Render())
// }

func (fragment *Fragment) Render() {
	html.Render(fragment.response, fragment.component(fragment).Element().GetNode())
}

func (fragment *Fragment) Request() *http.Request {
	return fragment.request
}

func newFragment(path FragmentPath) *Fragment {
	fragment := &Fragment{path: path}
	fragment.handler = func(w http.ResponseWriter, req *http.Request) {
		if req.Header.Get("HX-Request") != "true" {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "fragment can only requested by hx-request")
			return
		}

		fragment.request = req
		fragment.response = w

		log.Println(
			fmt.Sprintf("Fragment [%s]:", fragment.path),
			fmt.Sprintf("by %s#%s", req.Header.Get("HX-Trigger-Name"), req.Header.Get("HX-Trigger-Id")),
			fmt.Sprintf("for %s", req.Header.Get("Hx-Target")),
			fmt.Sprintf("from %s", req.Header.Get("Referer")),
		)
		fragment.Render()
	}

	return fragment
}

func FragmentComponent(component fragmentComponent) AppParam {
	return func(app *Application) (scope Scope, fn func(params ...any)) {
		return FragmentScope, func(params ...any) {
			if len(params) > 0 {
				if fragment, ok := params[0].(*Fragment); ok {
					fragment.component = component
				}
			}

		}
	}
}
