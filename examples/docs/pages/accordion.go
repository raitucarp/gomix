package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	disclosurePkg "github.com/raitucarp/gomix/components/disclosure"
)

func RenderAccordion() components.IsComponent {
	// Alias simulation for documentation
	var Accordion = disclosurePkg.Accordion

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/disclosure"

var Accordion = disclosure.Accordion`, nil),
		Playground("Basic Usage", "Standard implementation of the Accordion component.", `Accordion()`, Accordion().Component()),
		PropsTable([]map[string]string{
			{"name": "Collapsible", "type": "", "description": "Chainable method for Collapsible setting Gomix attributes/styles."},
			{"name": "Multiple", "type": "", "description": "Chainable method for Multiple setting Gomix attributes/styles."},
		}),
	).Component().GapBy(value.Rem(2))
}
