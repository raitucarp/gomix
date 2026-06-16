package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderSelectNativeDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default select_native component.", `formsPkg.SelectNative()`, formsPkg.SelectNative().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
