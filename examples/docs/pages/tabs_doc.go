package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	disclosurePkg "github.com/raitucarp/gomix/components/disclosure"
)

func RenderTabsDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default tabs component.", `disclosurePkg.Tabs()`, disclosurePkg.Tabs().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
