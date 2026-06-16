package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	overlaysPkg "github.com/raitucarp/gomix/components/overlays"
)

func RenderDrawerDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default drawer component.", `overlaysPkg.Drawer()`, overlaysPkg.Drawer().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
