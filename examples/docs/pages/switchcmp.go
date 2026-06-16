package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderSwitchCmp() components.IsComponent {
	// Alias simulation for documentation
	var SwitchCmp = formsPkg.SwitchCmp

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/forms"

var SwitchCmp = forms.SwitchCmp`, nil),
		Playground("Basic Usage", "Standard implementation of the SwitchCmp component.", `SwitchCmp()`, SwitchCmp().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
