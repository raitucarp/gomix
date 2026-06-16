package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	overlaysPkg "github.com/raitucarp/gomix/components/overlays"
)

func RenderModal() components.IsComponent {
	// Alias simulation for documentation
	var Modal = overlaysPkg.Modal

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/overlays"

var Modal = overlays.Modal`, nil),
		Playground("Basic Usage", "Standard implementation of the Modal component.", `Modal()`, Modal().Component()),
		Playground("Prop: IsOpen", "Demonstrating the IsOpen method.", `HStack(
	Modal().IsOpen(true),
	Modal().IsOpen(false)
)`, HStack(Modal().IsOpen(true).Component(), Modal().IsOpen(false).Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		PropsTable([]map[string]string{
			{"name": "IsOpen", "type": "isOpen bool", "description": "Chainable method for IsOpen setting Gomix attributes/styles."},
		}),
	).Component().GapBy(value.Rem(2))
}
