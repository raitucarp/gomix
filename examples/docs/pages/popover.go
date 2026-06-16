package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	overlaysPkg "github.com/raitucarp/gomix/components/overlays"
)

func RenderPopover() components.IsComponent {
	// Alias simulation for documentation
	var Popover = overlaysPkg.Popover

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/overlays"

var Popover = overlays.Popover`, nil),
		Playground("Basic Usage", "Standard implementation of the Popover component.", `Popover()`, Popover().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
