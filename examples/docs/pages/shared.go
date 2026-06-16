package pages

import (
	"github.com/raitucarp/gomix/element"

    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
    "github.com/raitucarp/gomix/components/typography"
    layoutComponent "github.com/raitucarp/gomix/components/layout"
    "github.com/raitucarp/gomix/components/data_display"
)

var VStack = layoutComponent.VStack
var HStack = layoutComponent.HStack
var Heading = typography.Heading
var TextCmp = typography.TextCmp
var Code = typography.Code
var Table = data_display.Table
var Td = element.Td
var Th = element.Th
var Tr = element.Tr
var Thead = element.THead
var Tbody = element.TBody

func Playground(title string, description string, code string, child components.IsComponent) components.IsComponent {
    var preview components.IsComponent = div()
    if child != nil {
        preview = VStack(child).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1)).WFull()
    }

    return VStack(
        Heading(text(title)).Component().TextXl().FontBold(),
        TextCmp(text(description)).Component().TextSm().TextColor(value.Hex("#666")),
        preview,
        Code(text(code)).Component().WFull().WhitespacePreWrap().PBy(value.Rem(1)).BgSlate(100).RoundedMd(),
    ).Component().GapBy(value.Rem(0.5)).WFull()
}

func PropsTable(methods []map[string]string) components.IsComponent {
    rows := []element.IsElement{}
    for _, m := range methods {
        rows = append(rows, Tr(
            Td(Code(text(m["name"])).Component()),
            Td(TextCmp(text(m["type"])).Component()),
            Td(TextCmp(text(m["description"])).Component()),
        ))
    }

    return VStack(
        Heading(text("Props / Methods")).Component().TextXl().FontBold(),
        Table(
            Thead(Tr(Th(text("Method")), Th(text("Args")), Th(text("Description")))),
            Tbody(rows...),
        ).Component().WFull().Border(value.Px(1)).BorderSlate(200).RoundedMd(),
    ).Component().GapBy(value.Rem(1)).WFull()
}

var text = element.Text
var div = element.Div
