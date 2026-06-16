package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderSwitchCmpDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default switch_cmp component.", `formsPkg.SwitchCmp()`, formsPkg.SwitchCmp().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
