package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	typographyPkg "github.com/raitucarp/gomix/components/typography"
)

func RenderHeadingDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default heading component.", `typographyPkg.Heading(text("Some Heading"))`, typographyPkg.Heading(text("Some Heading")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
