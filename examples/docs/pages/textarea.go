package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderTextarea() components.IsComponent {
	// Alias simulation for documentation
	var Textarea = formsPkg.Textarea

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/forms"

var Textarea = forms.Textarea`, nil),
		Playground("Basic Usage", "Standard implementation of the Textarea component.", `Textarea()`, Textarea().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
