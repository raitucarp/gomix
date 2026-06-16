package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	typographyPkg "github.com/raitucarp/gomix/components/typography"
)

func RenderTextCmp() components.IsComponent {
	// Alias simulation for documentation
	var TextCmp = typographyPkg.TextCmp

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/typography"

var TextCmp = typography.TextCmp`, nil),
		Playground("Basic Usage", "Standard implementation of the TextCmp component.", `TextCmp(text("Some text"))`, TextCmp(text("Some text")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
