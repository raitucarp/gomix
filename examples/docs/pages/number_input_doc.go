package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderNumberInputDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default number_input component.", `formsPkg.NumberInput()`, formsPkg.NumberInput().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
