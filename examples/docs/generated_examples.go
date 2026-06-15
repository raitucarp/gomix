package main

import (
	"github.com/raitucarp/gomix/icons/fontawesome-5"

)

func renderComponentExample(comp string) isComponent {
	switch comp {
	case "button":
		return div(
			h2(text("Example for button")),
			Button(text("Click me")).Component(),
		)
	case "close_button":
		return div(
			h2(text("Example for close_button")),
			CloseButton().Component(),
		)
	case "icon_button":
		return div(
			h2(text("Example for icon_button")),
			IconButton(text("X")).Component(),
		)
	case "badge":
		return div(
			h2(text("Example for badge")),
			Badge(text("New")).Component(),
		)
	case "divider":
		return div(
			h2(text("Example for divider")),
			Divider().Component(),
		)
	case "stat":
		return div(
			h2(text("Example for stat")),
			Stat().Component(),
		)
	case "table":
		return div(
			h2(text("Example for table")),
			Table().Component(),
		)
	case "tag":
		return div(
			h2(text("Example for tag")),
			Tag(text("Tag Content")).Component(),
		)
	case "accordion":
		return div(
			h2(text("Example for accordion")),
			Accordion().Component(),
		)
	case "disclosure":
		return div(
			h2(text("Example for disclosure")),
			text("Not implemented yet"),
		)
	case "tabs":
		return div(
			h2(text("Example for tabs")),
			Tabs().Component(),
		)
	case "visually_hidden":
		return div(
			h2(text("Example for visually_hidden")),
			VisuallyHidden(text("hidden")).Component(),
		)
	case "alert":
		return div(
			h2(text("Example for alert")),
			Alert().Component(),
		)
	case "circular_progress":
		return div(
			h2(text("Example for circular_progress")),
			CircularProgress().Component(),
		)
	case "progress":
		return div(
			h2(text("Example for progress")),
			Progress().Component(),
		)
	case "skeleton":
		return div(
			h2(text("Example for skeleton")),
			Skeleton().Component(),
		)
	case "spinner":
		return div(
			h2(text("Example for spinner")),
			Spinner().Component(),
		)
	case "toast":
		return div(
			h2(text("Example for toast")),
			Toast().Component(),
		)
	case "checkbox":
		return div(
			h2(text("Example for checkbox")),
			Checkbox().Component(),
		)
	case "input":
		return div(
			h2(text("Example for input")),
			Input().Component(),
		)
	case "number_input":
		return div(
			h2(text("Example for number_input")),
			NumberInput().Component(),
		)
	case "pin_input":
		return div(
			h2(text("Example for pin_input")),
			PinInput().Component(),
		)
	case "radio":
		return div(
			h2(text("Example for radio")),
			Radio().Component(),
		)
	case "select_native":
		return div(
			h2(text("Example for select_native")),
			SelectNative().Component(),
		)
	case "slider":
		return div(
			h2(text("Example for slider")),
			Slider().Component(),
		)
	case "switch_cmp":
		return div(
			h2(text("Example for switch_cmp")),
			SwitchCmp().Component(),
		)
	case "textarea":
		return div(
			h2(text("Example for textarea")),
			Textarea().Component(),
		)
	case "center":
		return div(
			h2(text("Example for center")),
			Center(text("Center")).Component(),
		)
	case "grid":
		return div(
			h2(text("Example for grid")),
			Grid().Component(),
		)
	case "simple_grid":
		return div(
			h2(text("Example for simple_grid")),
			SimpleGrid().Component(),
		)
	case "spacer":
		return div(
			h2(text("Example for spacer")),
			Spacer().Component(),
		)
	case "stack":
		return div(
			h2(text("Example for stack")),
			Stack().Component(),
		)
	case "wrap":
		return div(
			h2(text("Example for wrap")),
			Wrap().Component(),
		)
	case "avatar":
		return div(
			h2(text("Example for avatar")),
			Avatar().Component(),
		)
	case "icon":
		return div(
			h2(text("Example for icon")),
			Icon(fa5.FaBrand500px()).Component(),
		)
	case "image":
		return div(
			h2(text("Example for image")),
			Image().Component(),
		)
	case "alert_dialog":
		return div(
			h2(text("Example for alert_dialog")),
			AlertDialog().Component(),
		)
	case "drawer":
		return div(
			h2(text("Example for drawer")),
			Drawer().Component(),
		)
	case "menu":
		return div(
			h2(text("Example for menu")),
			Menu().Component(),
		)
	case "modal":
		return div(
			h2(text("Example for modal")),
			Modal().Component(),
		)
	case "popover":
		return div(
			h2(text("Example for popover")),
			Popover().Component(),
		)
	case "tooltip":
		return div(
			h2(text("Example for tooltip")),
			Tooltip().Component(),
		)
	case "code":
		return div(
			h2(text("Example for code")),
			Code(text("code snippets")).Component(),
		)
	case "heading":
		return div(
			h2(text("Example for heading")),
			Heading(text("Some heading")).Component(),
		)
	case "kbd":
		return div(
			h2(text("Example for kbd")),
			Kbd(text("Ctrl")).Component(),
		)
	case "text_cmp":
		return div(
			h2(text("Example for text_cmp")),
			TextCmp(text("Some text")).Component(),
		)
	default:
		return div(text("Component not found"))
	}
}
