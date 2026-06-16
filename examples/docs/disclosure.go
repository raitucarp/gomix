package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderDisclosure() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/disclosure"

var Disclosure = disclosure.Disclosure`, nil),
		Playground("Basic Usage", "Standard implementation of the Disclosure component.", `text("Not implemented")`, text("Not implemented yet")),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
