package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	data_displayPkg "github.com/raitucarp/gomix/components/data_display"
)

func RenderBadgeDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default badge component.", `data_displayPkg.Badge(text("New"))`, data_displayPkg.Badge(text("New")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
