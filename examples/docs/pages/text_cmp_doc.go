package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	typographyPkg "github.com/raitucarp/gomix/components/typography"
)

func RenderTextCmpDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default text_cmp component.", `typographyPkg.TextCmp(text("Some Text"))`, typographyPkg.TextCmp(text("Some Text")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
