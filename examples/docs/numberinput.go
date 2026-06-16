package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderNumberInput() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/forms"

var NumberInput = forms.NumberInput`, nil),
		Playground("Basic Usage", "Standard implementation of the NumberInput component.", `NumberInput()`, NumberInput().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
