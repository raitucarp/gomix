package main

import "github.com/raitucarp/gomix/value"

func SidebarLink(textContent string) isComponent {
	return ListItem(text(textContent)).Component().StyleHover(bg_slate(500))
}

var componentList = []string{
	"stack",
	"accordion",
	"list",
}

func Links(item string, index int) isComponent {
	return SidebarLink(item)
}

func Sidebar() isComponent {
	return Vstack(
		List().
			Ordered(
				ForEachComponentLink(componentList...).Map(Links),
			).
			TypeUpperAlpha().
			Unstyled().
			Outside().Component().MinWBy(value.Rem(13)).MaxWBy(value.Rem(15)),
	)
}
