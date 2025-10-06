package main

import (
	"net/http"
	"time"

	"github.com/raitucarp/gomix"
	"github.com/raitucarp/gomix/element"
)

func WelcomeText(message string) isComponent {
	return div(text(message)).
		Element().
		Style(
			font_2xl(),
		)
}

func MainLayout(message string) isComponent {
	return div(
		WelcomeText(message),
		slot(),
	)
}

func AppLayout() isComponent {
	return MainLayout("Welcome home")
}

func HomeLayout(page *gomix.Page) isComponent {
	return div(
		slot(),
		WelcomeText("whoopa").
			Element().
			Style(
				hover(
					font_3xl(),
					text_red(500),
					cursor_pointer(),
				),
			),
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

	raitucarpLink := a(
		text("Raitucarp is " + permalinkValue),
	).Href("https://raitucarp.xyz").Target(element.TargetBlank)

	return div(
		text("what is "+permalinkValue),
		text(" "),
		raitucarpLink,
	)
}

func ShopLayout(page *gomix.Page) isComponent {
	return div(
		text("Shop layout"),
	)
}
