package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	overlaysPkg "github.com/raitucarp/gomix/components/overlays"
)

func RenderMenu() components.IsComponent {
	// Alias simulation for documentation
	var Menu = overlaysPkg.Menu

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/overlays"

var Menu = overlays.Menu`, nil),
		Playground("Basic Usage", "Standard implementation of the Menu component.", `Menu()`, Menu().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
