package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderTextareaDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default textarea component.", `formsPkg.Textarea()`, formsPkg.Textarea().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
