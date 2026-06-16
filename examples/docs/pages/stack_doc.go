package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderStackDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default stack component.", `layoutPkg.Stack()`, layoutPkg.Stack().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
