package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderGridDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default grid component.", `layoutPkg.Grid()`, layoutPkg.Grid().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
