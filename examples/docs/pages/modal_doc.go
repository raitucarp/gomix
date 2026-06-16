package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	overlaysPkg "github.com/raitucarp/gomix/components/overlays"
)

func RenderModalDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default modal component.", `overlaysPkg.Modal()`, overlaysPkg.Modal().Component()),
		PropsTable([]map[string]string{
			{"name": "IsOpen", "type": "isOpen bool", "description": "Chainable method for IsOpen."},
		}),
	).Component().GapBy(value.Rem(2))
}
