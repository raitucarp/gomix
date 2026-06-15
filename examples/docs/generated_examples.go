package main

import (
	"github.com/raitucarp/gomix/value"
	"github.com/raitucarp/gomix/icons/fontawesome-5"
)

// A playground wrapper that shows the component, its code, and maybe variants/props.
func componentPlayground(title string, description string, children ...isComponent) isComponent {
    return VStack(
        Heading(text(title)).Component().TextXl().FontBold(),
        TextCmp(text(description)).Component().TextSm().TextColor(value.Hex("#666")),
        VStack(children...).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1)).WFull(),
    ).Component().GapBy(value.Rem(0.5)).WFull()
}

func renderComponentExample(comp string) isComponent {
	switch comp {
	case "accordion":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the accordion component.", Accordion().Component()),
		).Component().GapBy(value.Rem(2))
	case "tabs":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the tabs component.", Tabs().Component()),
		).Component().GapBy(value.Rem(2))
	case "visually_hidden":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the visually_hidden component.", VisuallyHidden(text("hidden")).Component()),
		).Component().GapBy(value.Rem(2))
	case "disclosure":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the disclosure component.", text("Not implemented yet")),
		).Component().GapBy(value.Rem(2))
	case "tooltip":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the tooltip component.", Tooltip().Component()),
		).Component().GapBy(value.Rem(2))
	case "menu":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the menu component.", Menu().Component()),
		).Component().GapBy(value.Rem(2))
	case "modal":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the modal component.", Modal().Component()),
		).Component().GapBy(value.Rem(2))
	case "popover":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the popover component.", Popover().Component()),
		).Component().GapBy(value.Rem(2))
	case "drawer":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the drawer component.", Drawer().Component()),
		).Component().GapBy(value.Rem(2))
	case "alert_dialog":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the alert_dialog component.", AlertDialog().Component()),
		).Component().GapBy(value.Rem(2))
	case "icon_button":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the icon_button component.", IconButton(text("X")).Component()),
		).Component().GapBy(value.Rem(2))
	case "button":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the button component.", Button(text("Click me")).Component()),
			componentPlayground("Variants", "Different variants available.",
				HStack(Button(text("Solid")).Variant("solid").Component(), Button(text("Outline")).Variant("outline").Component(), Button(text("Ghost")).Variant("ghost").Component()).Component().GapBy(value.Rem(1))),
			componentPlayground("Sizes", "Different sizes available.",
				HStack(Button(text("sm")).Size("sm").Component(), Button(text("md")).Size("md").Component(), Button(text("lg")).Size("lg").Component()).Component().GapBy(value.Rem(1))),
		).Component().GapBy(value.Rem(2))
	case "close_button":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the close_button component.", CloseButton().Component()),
		).Component().GapBy(value.Rem(2))
	case "kbd":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the kbd component.", HStack(Kbd(text("Ctrl")).Component(), text("+"), Kbd(text("C")).Component()).Component().GapBy(value.Rem(0.5))),
		).Component().GapBy(value.Rem(2))
	case "text_cmp":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the text_cmp component.", TextCmp(text("Some text")).Component()),
		).Component().GapBy(value.Rem(2))
	case "code":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the code component.", Code(text("fmt.Println(\"Hello, World!\")")).Component()),
		).Component().GapBy(value.Rem(2))
	case "heading":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the heading component.", Heading(text("Some heading")).Component()),
		).Component().GapBy(value.Rem(2))
	case "skeleton":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the skeleton component.", Skeleton().Component()),
		).Component().GapBy(value.Rem(2))
	case "toast":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the toast component.", Toast().Component()),
		).Component().GapBy(value.Rem(2))
	case "progress":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the progress component.", Progress().Component()),
		).Component().GapBy(value.Rem(2))
	case "circular_progress":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the circular_progress component.", CircularProgress().Component()),
		).Component().GapBy(value.Rem(2))
	case "spinner":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the spinner component.", Spinner().Component()),
		).Component().GapBy(value.Rem(2))
	case "alert":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the alert component.", Alert().Component()),
		).Component().GapBy(value.Rem(2))
	case "image":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the image component.", Image().Component()),
		).Component().GapBy(value.Rem(2))
	case "icon":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the icon component.", Icon(fa5.FaBrand500px()).Component()),
		).Component().GapBy(value.Rem(2))
	case "avatar":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the avatar component.", Avatar().Component()),
		).Component().GapBy(value.Rem(2))
	case "pin_input":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the pin_input component.", PinInput().Component()),
		).Component().GapBy(value.Rem(2))
	case "input":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the input component.", Input().Component()),
		).Component().GapBy(value.Rem(2))
	case "slider":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the slider component.", Slider().Component()),
		).Component().GapBy(value.Rem(2))
	case "radio":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the radio component.", Radio().Component()),
		).Component().GapBy(value.Rem(2))
	case "switch_cmp":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the switch_cmp component.", SwitchCmp().Component()),
		).Component().GapBy(value.Rem(2))
	case "checkbox":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the checkbox component.", Checkbox().Component()),
		).Component().GapBy(value.Rem(2))
	case "textarea":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the textarea component.", Textarea().Component()),
		).Component().GapBy(value.Rem(2))
	case "number_input":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the number_input component.", NumberInput().Component()),
		).Component().GapBy(value.Rem(2))
	case "select_native":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the select_native component.", SelectNative().Component()),
		).Component().GapBy(value.Rem(2))
	case "table":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the table component.", Table().Component()),
		).Component().GapBy(value.Rem(2))
	case "tag":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the tag component.", Tag(text("Tag Content")).Component()),
		).Component().GapBy(value.Rem(2))
	case "badge":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the badge component.", Badge(text("New")).Component()),
		).Component().GapBy(value.Rem(2))
	case "divider":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the divider component.", Divider().Component()),
		).Component().GapBy(value.Rem(2))
	case "stat":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the stat component.", Stat().Component()),
		).Component().GapBy(value.Rem(2))
	case "spacer":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the spacer component.", HStack(div(text("Left")), Spacer().Component(), div(text("Right"))).Component().WFull()),
		).Component().GapBy(value.Rem(2))
	case "stack":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the stack component.", Stack().Component()),
		).Component().GapBy(value.Rem(2))
	case "simple_grid":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the simple_grid component.", SimpleGrid().Component()),
		).Component().GapBy(value.Rem(2))
	case "grid":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the grid component.", Grid().Component()),
		).Component().GapBy(value.Rem(2))
	case "center":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the center component.", Center(text("Center")).Component()),
		).Component().GapBy(value.Rem(2))
	case "wrap":
		return VStack(
			componentPlayground("Basic Example", "The default usage of the wrap component.", Wrap().Component()),
		).Component().GapBy(value.Rem(2))
	default:
		return div(text("Component not found"))
	}
}
