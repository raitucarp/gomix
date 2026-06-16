package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	typographyPkg "github.com/raitucarp/gomix/components/typography"
)

func RenderCodeDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default code component.", `typographyPkg.Code(text("Some Code"))`, typographyPkg.Code(text("Some Code")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
