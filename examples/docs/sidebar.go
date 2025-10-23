package main

type ComponentGroup struct {
	label string
	links []ComponentLink
}

type ComponentLink struct {
	label string
	url   string
}

var componentsGroup = []ComponentGroup{
	{
		label: "Typography",
		links: []ComponentLink{
			{label: "List", url: "/components/list"},
			{label: "Link", url: "/components/link"},
		},
	},
	{
		label: "Layout",
		links: []ComponentLink{
			{label: "Stack", url: "/components/stack"},
			{label: "Card", url: "/components/card"},
		},
	},
}

func SidebarLink(link ComponentLink) isComponent {
	return ListItem(
		Link(text(link.label)).
			Href(link.url).
			Component().
			WFull().
			Block().
			StyleHover(bg_slate_alpha(200)).
			NoUnderline().
			P(2),
	)
}

func ComponentNav(item ComponentLink, index int) isComponent {
	return SidebarLink(item)
}

func ToComponentGroup(group ComponentGroup, index int) isComponent {
	return AccordionItem(
		Text(group.label).Component().FontBold().P(2),
		List().
			Ordered(
				ForEachComponentLink(group.links...).
					Map(ComponentNav),
			).
			Unstyled(),
	)
}

func Sidebar() isComponent {
	return Accordion(
		ForComponentGroup(componentsGroup...).
			Map(ToComponentGroup),
	).
		Component().
		// As(aside()).
		StyleHover(
			overflowY_scroll(),
		).
		MinWBy(rem(13)).
		OverflowHidden().
		HFull()
}
