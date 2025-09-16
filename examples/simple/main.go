package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/raitucarp/gomix"
	"github.com/raitucarp/gomix/addons"
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/components/styles"
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

func MainLayout() *components.Component {
	return components.Div().
		Children(
			components.TextRaw("This is main layout"),
			components.Slot(),
		)
}

func HomeLayout(page *gomix.Page) *components.Component {
	return components.Div().
		Children(
			components.Slot(),
			components.TextRaw("Home layout"),
		)
}

func BlogHomePage(page *gomix.Page) *components.Component {
	http.SetCookie(page.Response(), &http.Cookie{
		Name:     "my_cookie",
		Value:    "some_value",
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	return components.Div(
		components.Slot(),
		components.Div().Id("response-div"),
		components.TextRaw("Welcome to the blog!"),
		components.Button(components.TextRaw("testa")).
			Autofocus().
			Component().
			Id("button-good").
			HxGet("/blog/doyouwant").
			HxTarget("#response-div"),
	)
}

func BlogArticle(fragment *gomix.Fragment) *components.Component {
	bb := fragment.Request().PathValue("permalink")

	toAa := components.AsComponent(
		components.
			A(components.TextRaw("rere " + bb)).
			Href("http://aa.com").
			Target(components.AnchorBlank),
	)

	return components.Div(
		components.TextRaw("what is "+bb),
		toAa,
	)
}

func ShopLayout(page *gomix.Page) *components.Component {
	return components.Div().
		Styles(styles.New().Aspect("3/4")).
		Children(
			components.TextRaw("Shop layout"),
		)
}

func main() {

	ctx := context.Background()

	app := gomix.New(ctx, "Sample Test")
	app.Addons(addons.Htmx)

	app.EnableLogger()
	app.Layout(MainLayout())

	app.Fragment(blogArticle, BlogArticle)

	app.Page(homePath, HomePage)
	app.Page(shopPath, ShopPage)

	app.ServeAtPort(8080)

}
