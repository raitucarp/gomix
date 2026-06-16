package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	typographyPkg "github.com/raitucarp/gomix/components/typography"
)

func RenderKbdDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default kbd component.", `HStack(typographyPkg.Kbd(text("Ctrl")).Component(), text("+"), typographyPkg.Kbd(text("C")).Component())`, HStack(typographyPkg.Kbd(text("Ctrl")).Component(), text("+"), typographyPkg.Kbd(text("C")).Component()).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
