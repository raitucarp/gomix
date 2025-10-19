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
	homePath          = page_path("/")
	componentListPath = page_path("/components/")
	componentsPath    = page_path("/components/{component}")
	tutorialsPath     = page_path("/tutorials")
)

//go:embed alpine
var alpineGlobal embed.FS

func webSetup() gomix.AppParam {
	return web(
		app_layout(
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

		title_template("{title} | gomix Documentation"),

		// page
		page_at(homePath,
			title_("Getting Started"),
			component(HomeLayout),
		),
		page_at(componentListPath,
			title_("Components Gallery"),
			layout(ComponentsLayout),
			component(ComponentsGallery),

			page_at(componentsPath,
				title_("{component} Component"),
				component(ComponentContent),
			),
		),
		page_at(tutorialsPath,
			title_("Tutorials"),
			component(TutorialPage),
		),

		not_found_page(
			title_("404 Not Found"),
			component(NotFoundPage),
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
