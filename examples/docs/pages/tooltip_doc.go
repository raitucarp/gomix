package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	overlaysPkg "github.com/raitucarp/gomix/components/overlays"
)

func RenderTooltipDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default tooltip component.", `overlaysPkg.Tooltip()`, overlaysPkg.Tooltip().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
