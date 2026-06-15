package main

import (
	"github.com/raitucarp/gomix/value"

	"github.com/raitucarp/gomix/icons/fontawesome-5"

)

func renderComponentExample(comp string) isComponent {
	switch comp {
	case "button":
		return VStack(
			Heading(text("Example for button")).Component().TextLg().FontBold(),
			Button(text("Click me")).Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "close_button":
		return VStack(
			Heading(text("Example for close_button")).Component().TextLg().FontBold(),
			CloseButton().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "icon_button":
		return VStack(
			Heading(text("Example for icon_button")).Component().TextLg().FontBold(),
			IconButton(text("X")).Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "badge":
		return VStack(
			Heading(text("Example for badge")).Component().TextLg().FontBold(),
			Badge(text("New")).Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "divider":
		return VStack(
			Heading(text("Example for divider")).Component().TextLg().FontBold(),
			Divider().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "stat":
		return VStack(
			Heading(text("Example for stat")).Component().TextLg().FontBold(),
			Stat().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "table":
		return VStack(
			Heading(text("Example for table")).Component().TextLg().FontBold(),
			Table().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "tag":
		return VStack(
			Heading(text("Example for tag")).Component().TextLg().FontBold(),
			Tag(text("Tag Content")).Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "accordion":
		return VStack(
			Heading(text("Example for accordion")).Component().TextLg().FontBold(),
			Accordion().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "disclosure":
		return VStack(
			Heading(text("Example for disclosure")).Component().TextLg().FontBold(),
			text("Not implemented yet"),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "tabs":
		return VStack(
			Heading(text("Example for tabs")).Component().TextLg().FontBold(),
			Tabs().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "visually_hidden":
		return VStack(
			Heading(text("Example for visually_hidden")).Component().TextLg().FontBold(),
			VisuallyHidden(text("hidden")).Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "alert":
		return VStack(
			Heading(text("Example for alert")).Component().TextLg().FontBold(),
			Alert().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "circular_progress":
		return VStack(
			Heading(text("Example for circular_progress")).Component().TextLg().FontBold(),
			CircularProgress().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "progress":
		return VStack(
			Heading(text("Example for progress")).Component().TextLg().FontBold(),
			Progress().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "skeleton":
		return VStack(
			Heading(text("Example for skeleton")).Component().TextLg().FontBold(),
			Skeleton().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "spinner":
		return VStack(
			Heading(text("Example for spinner")).Component().TextLg().FontBold(),
			Spinner().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "toast":
		return VStack(
			Heading(text("Example for toast")).Component().TextLg().FontBold(),
			Toast().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "checkbox":
		return VStack(
			Heading(text("Example for checkbox")).Component().TextLg().FontBold(),
			Checkbox().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "input":
		return VStack(
			Heading(text("Example for input")).Component().TextLg().FontBold(),
			Input().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "number_input":
		return VStack(
			Heading(text("Example for number_input")).Component().TextLg().FontBold(),
			NumberInput().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "pin_input":
		return VStack(
			Heading(text("Example for pin_input")).Component().TextLg().FontBold(),
			PinInput().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "radio":
		return VStack(
			Heading(text("Example for radio")).Component().TextLg().FontBold(),
			Radio().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "select_native":
		return VStack(
			Heading(text("Example for select_native")).Component().TextLg().FontBold(),
			SelectNative().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "slider":
		return VStack(
			Heading(text("Example for slider")).Component().TextLg().FontBold(),
			Slider().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "switch_cmp":
		return VStack(
			Heading(text("Example for switch_cmp")).Component().TextLg().FontBold(),
			SwitchCmp().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "textarea":
		return VStack(
			Heading(text("Example for textarea")).Component().TextLg().FontBold(),
			Textarea().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "center":
		return VStack(
			Heading(text("Example for center")).Component().TextLg().FontBold(),
			Center(text("Center")).Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "grid":
		return VStack(
			Heading(text("Example for grid")).Component().TextLg().FontBold(),
			Grid().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "simple_grid":
		return VStack(
			Heading(text("Example for simple_grid")).Component().TextLg().FontBold(),
			SimpleGrid().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "spacer":
		return VStack(
			Heading(text("Example for spacer")).Component().TextLg().FontBold(),
			Spacer().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "stack":
		return VStack(
			Heading(text("Example for stack")).Component().TextLg().FontBold(),
			Stack().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "wrap":
		return VStack(
			Heading(text("Example for wrap")).Component().TextLg().FontBold(),
			Wrap().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "avatar":
		return VStack(
			Heading(text("Example for avatar")).Component().TextLg().FontBold(),
			Avatar().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "icon":
		return VStack(
			Heading(text("Example for icon")).Component().TextLg().FontBold(),
			Icon(fa5.FaBrand500px()).Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "image":
		return VStack(
			Heading(text("Example for image")).Component().TextLg().FontBold(),
			Image().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "alert_dialog":
		return VStack(
			Heading(text("Example for alert_dialog")).Component().TextLg().FontBold(),
			AlertDialog().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "drawer":
		return VStack(
			Heading(text("Example for drawer")).Component().TextLg().FontBold(),
			Drawer().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "menu":
		return VStack(
			Heading(text("Example for menu")).Component().TextLg().FontBold(),
			Menu().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "modal":
		return VStack(
			Heading(text("Example for modal")).Component().TextLg().FontBold(),
			Modal().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "popover":
		return VStack(
			Heading(text("Example for popover")).Component().TextLg().FontBold(),
			Popover().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "tooltip":
		return VStack(
			Heading(text("Example for tooltip")).Component().TextLg().FontBold(),
			Tooltip().Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "code":
		return VStack(
			Heading(text("Example for code")).Component().TextLg().FontBold(),
			Code(text("code snippets")).Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "heading":
		return VStack(
			Heading(text("Example for heading")).Component().TextLg().FontBold(),
			Heading(text("Some heading")).Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "kbd":
		return VStack(
			Heading(text("Example for kbd")).Component().TextLg().FontBold(),
			Kbd(text("Ctrl")).Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	case "text_cmp":
		return VStack(
			Heading(text("Example for text_cmp")).Component().TextLg().FontBold(),
			TextCmp(text("Some text")).Component(),
		).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1))
	default:
		return VStack(text("Component not found"))
	}
}
