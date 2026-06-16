package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	data_displayPkg "github.com/raitucarp/gomix/components/data_display"
)

func RenderDividerDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default divider component.", `data_displayPkg.Divider()`, data_displayPkg.Divider().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
