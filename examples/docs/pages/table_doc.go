package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	data_displayPkg "github.com/raitucarp/gomix/components/data_display"
)

func RenderTableDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default table component.", `data_displayPkg.Table()`, data_displayPkg.Table().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
