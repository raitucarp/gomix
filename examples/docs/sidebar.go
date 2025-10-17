package main

import "github.com/raitucarp/gomix/value"

func SidebarLink(textContent string) isComponent {
	return ListItem(text(textContent)).Component().StyleHover(bg_slate(500))
}

func Links() []isComponent {
	return []isComponent{
		SidebarLink("test 1"),
		SidebarLink("test 3"),
	}
}

func Sidebar() isComponent {
	return Vstack(
		List().
			Ordered(
				Links()...,
			).
			TypeUpperAlpha().
			Unstyled().
			Outside().Component().MinWBy(value.Rem(13)).MaxWBy(value.Rem(15)),
	)
}
