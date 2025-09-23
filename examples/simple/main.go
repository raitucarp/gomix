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
	blogArticleFramentPath gomix.FragmentPath = "/blog/{permalink}"
)

type isComponent = components.IsComponent

func MainLayout() isComponent {
	return div(
		text("This is main layout"),
		slot(),
	)
}

func HomeLayout(page *gomix.Page) isComponent {
	return div(
		slot(),
		text("Home layout"),
	)
}

func BlogHomePage(page *gomix.Page) isComponent {
	http.SetCookie(page.Response(), &http.Cookie{
		Name:     "my_cookie",
		Value:    "some_value",
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	return div(
		slot(),
		div(
			text("this is response div"),
			div(text("inner div")),
		).Element().Id("response-div"),
		text("Welcome to the blog!"),
		Htmx(
			button(text("testa")).Autofocus().Element().Id("button-good"),
		).Get("/blog/doyouwant").Target("#response-div"),
	)
}

func BlogArticle(fragment *gomix.Fragment) isComponent {
	permalinkValue := fragment.Request().PathValue("permalink")

	aaLink := a(
		text("rere " + permalinkValue),
	).Href("http://aa.com").Target(element.TargetBlank)

	return div(
		text("what is "+permalinkValue),
		text(" "),
		aaLink,
	)
}

func ShopLayout(page *gomix.Page) isComponent {
	return div(
		text("Shop layout"),
	)
}

func main() {

	/*
		app(
			addons(
				htmx("2.0.7"),
			),
			features(
				logger(),
			),
			pages(
				page(homePath, HomePage),
				page(shopPath, ShopPage),
			),
			port(8080),
		)
	*/

	app := gomix.New("Sample Test")
	app.Addons(
		htmx.Addon("2.0.7"),
	)

	app.EnableLogger()
	app.Layout(MainLayout())

	app.Fragment(blogArticleFramentPath, BlogArticle)

	app.Page(homePath, HomePage)
	app.Page(shopPath, ShopPage)

	app.ServeAtPort(8080)

}
