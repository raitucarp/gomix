package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderKbd() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/typography"

var Kbd = typography.Kbd`, nil),
		Playground("Basic Usage", "Standard implementation of the Kbd component.", `HStack(Kbd(text("Ctrl")), text("+"), Kbd(text("C")))`, HStack(Kbd(text("Ctrl")), text("+"), Kbd(text("C"))).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
