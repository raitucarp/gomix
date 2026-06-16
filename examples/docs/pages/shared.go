package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/element"
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

func Playground(title string, description string, code string, children ...components.IsComponent) components.IsComponent {
    return VStack(
        Heading(element.Text(title)).Component().TextXl().FontBold(),
        TextCmp(element.Text(description)).Component().TextSm().TextColor(value.Hex("#666")),
        VStack(children...).Component().Border(value.Px(1)).BorderSlate(200).RoundedMd().PBy(value.Rem(1.5)).GapBy(value.Rem(1)).WFull(),
        Code(element.Text(code)).Component().WFull().WhitespacePreWrap().PBy(value.Rem(1)).BgSlate(100).RoundedMd(),
    ).Component().GapBy(value.Rem(0.5)).WFull()
}

func PropsTable(methods []map[string]string) components.IsComponent {
    rows := []element.IsElement{}
    for _, m := range methods {
        rows = append(rows, Tr(
            Td(Code(element.Text(m["name"])).Component()),
            Td(TextCmp(element.Text(m["type"])).Component()),
            Td(TextCmp(element.Text(m["description"])).Component()),
        ))
    }

    return VStack(
        Heading(element.Text("Props / Methods")).Component().TextXl().FontBold(),
        Table(
            Thead(Tr(Th(element.Text("Method")), Th(element.Text("Args")), Th(element.Text("Description")))),
            Tbody(rows...),
        ).Component().WFull().Border(value.Px(1)).BorderSlate(200).RoundedMd(),
    ).Component().GapBy(value.Rem(1)).WFull()
}

var text = element.Text
var div = element.Div
