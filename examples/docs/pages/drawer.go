package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	overlaysPkg "github.com/raitucarp/gomix/components/overlays"
)

func RenderDrawer() components.IsComponent {
	// Alias simulation for documentation
	var Drawer = overlaysPkg.Drawer

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/overlays"

var Drawer = overlays.Drawer`, nil),
		Playground("Basic Usage", "Standard implementation of the Drawer component.", `Drawer()`, Drawer().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
