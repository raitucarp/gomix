package main

import (
	"github.com/raitucarp/gomix"
	"github.com/raitucarp/gomix/value"
)

func AppLayout() isComponent {
	return div(slot())
}

type DropdownData struct {
	Open bool `json:"open"`
}

func ComponentContent(page *gomix.Page) isComponent {
	return alpine(
		div(
			alpine(
				button(text("Toggle Dropdown")),
			).On("click", "open = ! open"),

			alpine(
				div(text("Dropdown Contents...")),
			).Show("open"),
		),
	).Data(&DropdownData{Open: false})

}

func ComponentsGallery(page *gomix.Page) isComponent {
	return div(text("Component Gallery"))
}

func ComponentsLayout(page *gomix.Page) isComponent {
	return Hstack(
		Sidebar(),
		div(
			slot(),
		).
			Element().
			Style(flex_1),
	).
		Component().
		PBy(value.Rem(.8)).
		GapBy(value.Rem(0.5)).
		WFull().
		HScreen().
		JustifyCenter()
}

func HomeLayout(page *gomix.Page) isComponent {
	return div(slot())
}

func TutorialPage(page *gomix.Page) isComponent {
	return div(
		text("tutorial"),
	)
}

func NotFoundPage(page *gomix.Page) isComponent {
	return div(
		text("Custom 404 Page Not Found"),
	)
}
