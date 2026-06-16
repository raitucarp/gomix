package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	data_displayPkg "github.com/raitucarp/gomix/components/data_display"
)

func RenderTagDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default tag component.", `data_displayPkg.Tag(text("Tag"))`, data_displayPkg.Tag(text("Tag")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
