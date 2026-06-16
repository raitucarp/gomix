package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderAccordionItem() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/disclosure"

var AccordionItem = disclosure.AccordionItem`, nil),
		Playground("Basic Usage", "Standard implementation of the AccordionItem component.", `AccordionItem(text("Title"), text("Content"))`, AccordionItem(text("Title"), text("Content")).Component()),
		PropsTable([]map[string]string{
			{"name": "CollapseIcon", "type": "icon icons.IsIcon", "description": "Chainable method for CollapseIcon setting Gomix attributes/styles."},
			{"name": "ExpandIcon", "type": "icon icons.IsIcon", "description": "Chainable method for ExpandIcon setting Gomix attributes/styles."},
		}),
	).Component().GapBy(value.Rem(2))
}
