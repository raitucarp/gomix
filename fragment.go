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
	component FragmentComponent

	request  *http.Request
	response http.ResponseWriter
	handler  func(w http.ResponseWriter, req *http.Request)
}

type FragmentComponent func(fragment *Fragment) *components.Component

func (app *App) Fragment(path FragmentPath, component FragmentComponent) *Fragment {
	fragment := NewFragment(path, component)
	app.fragments = append(app.fragments, fragment)

	return fragment
}

// func (fragment *Fragment) Write() {
// 	io.WriteString(fragment.response, fragment.Render())
// }

func (fragment *Fragment) Render() {
	component := fragment.component(fragment)

	html.Render(fragment.response, component.GetNode())
}

func (fragment *Fragment) Request() *http.Request {
	return fragment.request
}

func NewFragment(path FragmentPath, component FragmentComponent) *Fragment {
	fragment := &Fragment{path: path, component: component}
	fragment.handler = func(w http.ResponseWriter, req *http.Request) {
		if req.Header.Get("HX-Request") != "true" {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "no fragment")
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
