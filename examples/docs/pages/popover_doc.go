package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	overlaysPkg "github.com/raitucarp/gomix/components/overlays"
)

func RenderPopoverDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default popover component.", `overlaysPkg.Popover()`, overlaysPkg.Popover().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
