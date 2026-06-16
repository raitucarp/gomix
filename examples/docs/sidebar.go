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
		label: "Buttons",
		links: []ComponentLink{
			{label: "Button", url: "/gomix/docs/components/button"},
			{label: "Close Button", url: "/gomix/docs/components/close_button"},
			{label: "Icon Button", url: "/gomix/docs/components/icon_button"},
		},
	},
	{
		label: "Data Display",
		links: []ComponentLink{
			{label: "Badge", url: "/gomix/docs/components/badge"},
			{label: "Divider", url: "/gomix/docs/components/divider"},
			{label: "Stat", url: "/gomix/docs/components/stat"},
			{label: "Table", url: "/gomix/docs/components/table"},
			{label: "Tag", url: "/gomix/docs/components/tag"},
		},
	},
	{
		label: "Disclosure",
		links: []ComponentLink{
			{label: "Accordion", url: "/gomix/docs/components/accordion"},
			{label: "Disclosure", url: "/gomix/docs/components/disclosure"},
			{label: "Tabs", url: "/gomix/docs/components/tabs"},
			{label: "Visually Hidden", url: "/gomix/docs/components/visually_hidden"},
		},
	},
	{
		label: "Feedback",
		links: []ComponentLink{
			{label: "Alert", url: "/gomix/docs/components/alert"},
			{label: "Circular Progress", url: "/gomix/docs/components/circular_progress"},
			{label: "Progress", url: "/gomix/docs/components/progress"},
			{label: "Skeleton", url: "/gomix/docs/components/skeleton"},
			{label: "Spinner", url: "/gomix/docs/components/spinner"},
			{label: "Toast", url: "/gomix/docs/components/toast"},
		},
	},
	{
		label: "Forms",
		links: []ComponentLink{
			{label: "Checkbox", url: "/gomix/docs/components/checkbox"},
			{label: "Input", url: "/gomix/docs/components/input"},
			{label: "Number Input", url: "/gomix/docs/components/number_input"},
			{label: "Pin Input", url: "/gomix/docs/components/pin_input"},
			{label: "Radio", url: "/gomix/docs/components/radio"},
			{label: "Select Native", url: "/gomix/docs/components/select_native"},
			{label: "Slider", url: "/gomix/docs/components/slider"},
			{label: "Switch Cmp", url: "/gomix/docs/components/switch_cmp"},
			{label: "Textarea", url: "/gomix/docs/components/textarea"},
		},
	},
	{
		label: "Layout",
		links: []ComponentLink{
			{label: "Center", url: "/gomix/docs/components/center"},
			{label: "Grid", url: "/gomix/docs/components/grid"},
			{label: "Simple Grid", url: "/gomix/docs/components/simple_grid"},
			{label: "Spacer", url: "/gomix/docs/components/spacer"},
			{label: "Stack", url: "/gomix/docs/components/stack"},
			{label: "Wrap", url: "/gomix/docs/components/wrap"},
		},
	},
	{
		label: "Media And Icons",
		links: []ComponentLink{
			{label: "Avatar", url: "/gomix/docs/components/avatar"},
			{label: "Icon", url: "/gomix/docs/components/icon"},
			{label: "Image", url: "/gomix/docs/components/image"},
		},
	},
	{
		label: "Overlays",
		links: []ComponentLink{
			{label: "Alert Dialog", url: "/gomix/docs/components/alert_dialog"},
			{label: "Drawer", url: "/gomix/docs/components/drawer"},
			{label: "Menu", url: "/gomix/docs/components/menu"},
			{label: "Modal", url: "/gomix/docs/components/modal"},
			{label: "Popover", url: "/gomix/docs/components/popover"},
			{label: "Tooltip", url: "/gomix/docs/components/tooltip"},
		},
	},
	{
		label: "Typography",
		links: []ComponentLink{
			{label: "Code", url: "/gomix/docs/components/code"},
			{label: "Heading", url: "/gomix/docs/components/heading"},
			{label: "Kbd", url: "/gomix/docs/components/kbd"},
			{label: "Text Cmp", url: "/gomix/docs/components/text_cmp"},
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
