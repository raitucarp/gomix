package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	overlaysPkg "github.com/raitucarp/gomix/components/overlays"
)

func RenderMenuDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default menu component.", `overlaysPkg.Menu()`, overlaysPkg.Menu().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
