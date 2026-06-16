package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderWrapDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default wrap component.", `layoutPkg.Wrap()`, layoutPkg.Wrap().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
