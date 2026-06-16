package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	data_displayPkg "github.com/raitucarp/gomix/components/data_display"
)

func RenderStatDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default stat component.", `data_displayPkg.Stat()`, data_displayPkg.Stat().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
