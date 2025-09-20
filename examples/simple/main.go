package main

import (
	"log"
	"net/http"
	"time"

	"github.com/raitucarp/gomix"
	"github.com/raitucarp/gomix/addons/htmx"
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	homePath      gomix.LocationPath = "/"
	blogPath      gomix.LocationPath = "/blog"
	blogAboutPath gomix.LocationPath = "/blog/about/{name}"
	shopPath      gomix.LocationPath = "/shop"
)

const (
	blogArticle gomix.FragmentPath = "/blog/{permalink}"
)

func MainLayout() components.IsComponent {
	return element.Div(
		element.Text("This is main layout"),
		components.Slot(),
	)
}

func HomeLayout(page *gomix.Page) components.IsComponent {
	return element.Div(
		components.Slot(),
		element.Text("Home layout"),
	)
}

func BlogHomePage(page *gomix.Page) components.IsComponent {
	http.SetCookie(page.Response(), &http.Cookie{
		Name:     "my_cookie",
		Value:    "some_value",
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	return element.Div(
		components.Slot(),
		element.Element(element.Div()).Id("response-div"),
		element.Text("Welcome to the blog!"),
		htmx.Htmx(
			element.Element(
				element.Button(element.Text("testa")).Autofocus(),
			).Id("button-good"),
		).
			Get("/blog/doyouwant").
			Target("#response-div"),
	)
}

func BlogArticle(fragment *gomix.Fragment) components.IsComponent {
	bb := fragment.Request().PathValue("permalink")

	toAa := element.A(element.Text("rere " + bb)).
		Href("http://aa.com").
		Target(element.TargetBlank)

	return element.Div(
		element.Text("what is "+bb),
		toAa,
	)
}

func ShopLayout(page *gomix.Page) components.IsComponent {
	return element.Div(element.Text("Shop layout"))
}

func main() {

	app := gomix.New("Sample Test")
	app.Addons(
		htmx.Addon("2.0.7"),
	)

	app.EnableLogger()
	app.Layout(MainLayout())

	app.Fragment(blogArticle, BlogArticle)

	app.Page(homePath, HomePage)
	app.Page(shopPath, ShopPage)

	app.ServeAtPort(8080)

}
