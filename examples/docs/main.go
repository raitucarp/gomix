package main

import (
	"embed"
	"log"

	"github.com/raitucarp/gomix"
	"github.com/raitucarp/gomix/addons/alpinejs"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	homePath      = page_path("/")
	tutorialsPath = page_path("/tutorials")
)

//go:embed alpine
var alpineGlobal embed.FS

func webSetup() gomix.AppParam {
	return web(
		layout(
			AppLayout(),
		),
		web_addons(
			// _htmx("2.0.7"),
			_modern_normalize(),
			_alpine_addon("3.x.x"),
			alpinejs.GlobalsEmbed(alpineGlobal),
		),

		style_global(css{
			"html": style(
				bg_slate(950),
				text_white(),
				dark(
					bg_slate(950),
					text_white(),
				),
			),
		}),

		// fragment_at(blogArticleFramentPath,
		// 	fragment(BlogArticleFragment),
		// ),

		// page
		page_at(homePath,
			title_("gomix Documentation"),
			component(HomeLayout),
		),
		page_at(tutorialsPath,
			title_("Tutorials"),
			component(TutorialPage),
		),
	)

}

func main() {
	app(
		name("gomix Documentation"),
		features(
			logger(),
		),
		webSetup(),
		port(8080),
	)

}
