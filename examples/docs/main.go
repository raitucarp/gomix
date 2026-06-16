package main

import (
	"embed"
	"log"

	"github.com/raitucarp/gomix"
	"github.com/raitucarp/gomix/addons/alpinejs"
	"github.com/raitucarp/gomix/components/disclosure"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	homePath          = page_path("/gomix/docs/")
	componentListPath = page_path("/gomix/docs/components/")
	componentsPath    = page_path("/gomix/docs/components/{component}")
	tutorialsPath     = page_path("/gomix/docs/tutorials")
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
			alpinejs.CollapsePlugin("3.x.x"),
			alpinejs.Core("3.x.x"),
			alpinejs.GlobalsEmbed(alpineGlobal),

			disclosure.UseDisclosureComponents(),
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
				static_paths("/gomix/docs/components/button", "/gomix/docs/components/close_button", "/gomix/docs/components/icon_button", "/gomix/docs/components/badge", "/gomix/docs/components/divider", "/gomix/docs/components/stat", "/gomix/docs/components/table", "/gomix/docs/components/tag", "/gomix/docs/components/accordion", "/gomix/docs/components/disclosure", "/gomix/docs/components/tabs", "/gomix/docs/components/visually_hidden", "/gomix/docs/components/alert", "/gomix/docs/components/circular_progress", "/gomix/docs/components/progress", "/gomix/docs/components/skeleton", "/gomix/docs/components/spinner", "/gomix/docs/components/toast", "/gomix/docs/components/checkbox", "/gomix/docs/components/input", "/gomix/docs/components/number_input", "/gomix/docs/components/pin_input", "/gomix/docs/components/radio", "/gomix/docs/components/select_native", "/gomix/docs/components/slider", "/gomix/docs/components/switch_cmp", "/gomix/docs/components/textarea", "/gomix/docs/components/center", "/gomix/docs/components/grid", "/gomix/docs/components/simple_grid", "/gomix/docs/components/spacer", "/gomix/docs/components/stack", "/gomix/docs/components/wrap", "/gomix/docs/components/avatar", "/gomix/docs/components/icon", "/gomix/docs/components/image", "/gomix/docs/components/alert_dialog", "/gomix/docs/components/drawer", "/gomix/docs/components/menu", "/gomix/docs/components/modal", "/gomix/docs/components/popover", "/gomix/docs/components/tooltip", "/gomix/docs/components/code", "/gomix/docs/components/heading", "/gomix/docs/components/kbd", "/gomix/docs/components/text_cmp"),

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

	appParams := []gomix.AppParam{
		name("gomix Documentation"),
		features(
			logger(),
		),
		webSetup(),
	}

	appParams = append(appParams, static("./dist"))

	app(appParams...)
}
