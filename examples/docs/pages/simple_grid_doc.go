package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderSimpleGridDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default simple_grid component.", `layoutPkg.SimpleGrid()`, layoutPkg.SimpleGrid().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
