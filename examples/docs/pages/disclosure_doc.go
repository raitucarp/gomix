package pages

import (
	"github.com/raitucarp/gomix/element"

	textCmpPkg "github.com/raitucarp/gomix/components/typography"

    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderDisclosureDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default disclosure component.", `"Not implemented"`, textCmpPkg.TextCmp(element.Text("Not implemented yet")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
