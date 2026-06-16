package main

import (
	"github.com/raitucarp/gomix/examples/docs/pages"

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
	compParam := page.Param("component")
	var compName string
	if compParam != nil {
		compName = compParam.(string)
	}

    rendered := pages.Render(compName)
    if rendered == nil {
        rendered = div(text("Not implemented fully mapped for " + compName))
    }

	return VStack(
		Heading(text(compName)).Component().FontBold().TextXl(),
		rendered,
	).Component().PBy(value.Rem(1)).GapBy(value.Rem(1)).WFull()
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
