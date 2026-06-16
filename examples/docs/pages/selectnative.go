package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderSelectNative() components.IsComponent {
	// Alias simulation for documentation
	var SelectNative = formsPkg.SelectNative

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/forms"

var SelectNative = forms.SelectNative`, nil),
		Playground("Basic Usage", "Standard implementation of the SelectNative component.", `SelectNative()`, SelectNative().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
