package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	overlaysPkg "github.com/raitucarp/gomix/components/overlays"
)

func RenderAlertDialog() components.IsComponent {
	// Alias simulation for documentation
	var AlertDialog = overlaysPkg.AlertDialog

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/overlays"

var AlertDialog = overlays.AlertDialog`, nil),
		Playground("Basic Usage", "Standard implementation of the AlertDialog component.", `AlertDialog()`, AlertDialog().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
