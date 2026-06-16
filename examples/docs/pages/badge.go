package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	data_displayPkg "github.com/raitucarp/gomix/components/data_display"
)

func RenderBadge() components.IsComponent {
	// Alias simulation for documentation
	var Badge = data_displayPkg.Badge

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/data_display"

var Badge = data_display.Badge`, nil),
		Playground("Basic Usage", "Standard implementation of the Badge component.", `Badge(text("New"))`, Badge(text("New")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
