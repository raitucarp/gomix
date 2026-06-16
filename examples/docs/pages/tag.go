package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	data_displayPkg "github.com/raitucarp/gomix/components/data_display"
)

func RenderTag() components.IsComponent {
	// Alias simulation for documentation
	var Tag = data_displayPkg.Tag

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/data_display"

var Tag = data_display.Tag`, nil),
		Playground("Basic Usage", "Standard implementation of the Tag component.", `Tag(text("Tag"))`, Tag(text("Tag")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
