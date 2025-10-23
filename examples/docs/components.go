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
	return div(

		alpine(
			div(
				alpine(
					button(text("Toggle Dropdown")),
				).On("click", "open = ! open"),

				alpine(
					div(text("Dropdown Contents...")),
				).Show("open"),
			),
		).Data(&DropdownData{Open: false}),

		Accordion(
			AccordionItem(
				div(text("baraya")),
				div(text("Content 1")),
			),
			AccordionItem(
				div(text("baraya2")),
				div(text("Content 2")),
			),
			AccordionItem(
				div(text("babaa 3")),
				div(text("Content 3")),
			),
		),
	)

}

func ComponentsGallery(page *gomix.Page) isComponent {
	return div(text("Component Gallery"))
}

func ComponentsLayout(page *gomix.Page) isComponent {
	return HStack(
		Sidebar(),
		Main(
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
