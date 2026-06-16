package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderCode() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/typography"

var Code = typography.Code`, nil),
		Playground("Basic Usage", "Standard implementation of the Code component.", `Code(text("fmt.Println(\"Hello\")"))`, Code(text("fmt.Println(\"Hello\")")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
