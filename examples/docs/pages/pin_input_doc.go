package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderPinInputDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default pin_input component.", `formsPkg.PinInput()`, formsPkg.PinInput().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
