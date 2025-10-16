package main

import (
	"net/http"
	"time"

	"github.com/raitucarp/gomix"
	googlefonts "github.com/raitucarp/gomix/addons/google-fonts"
	"github.com/raitucarp/gomix/element"
)

func StyledText(message string) isComponent {
	return div(text(message)).
		Element().
		Style(
			text_2xl(),
		)
}

func MainLayout(message string) isComponent {
	return div(
		StyledText(message).
			Element().
			Style(
				text_blue(600),
				dark(
					text_red(400),
				),
			),
		slot(),
	)
}

func AppLayout() isComponent {
	return MainLayout("Welcome home")
}

func HomeLayout(page *gomix.Page) isComponent {
	return div(
		slot(),
		text("home layout"),
		Hstack(
			span(text("test")).Element().Style(
				text_3xl(),
				googlefonts.FontPlaywriteDeGrund(),
			),
			div(
				Icon(fa_brand_android()).Size2Xl().Component().FillRed(300),
				span(text("Font Awesome Brand Android")),
			).Element().Style(flex_1),
			div(
				span(text("ai acount book")),
				Icon(ai_account_book()).Size3Xl().Component().FillPurple(900),
			).Element().Style(flex_1),
		).
			Component().
			WFull().
			JustifyCenter(),
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
		div(
			text("this is response div"),
			div(text("inner div")),
		).Element().Id("response-div"),
		text("Welcome to the blog!"),
		Htmx(
			button(text("testa")).Autofocus().Element().Id("button-good"),
		).Get("/blog/doyouwant").Target("#response-div"),
		slot(),
		StyledText("whoopa").
			Element().
			PreferredClassName("thanks").
			Style(
				text_red(900),
				text_2xl(),

				dark(
					text_green(900),
					text_3xl(),

					hover(
						text_blue(900),
					),
				),

				hover(
					text_7xl(),
					text_orange(900),
					cursor_pointer(),
				),
			),
	)
}

func BlogAboutPage(page *gomix.Page) isComponent {
	name := page.Request().PathValue("name")

	return div(text(name)).Element().Style(text_green(900))
}

func BlogArticleFragment(fragment *gomix.Fragment) isComponent {
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
