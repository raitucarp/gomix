package main

import (
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	homePath               = page_path("/")
	blogPath               = page_path("/blog")
	blogAboutPath          = page_path("/blog/about/{name}")
	shopPath               = page_path("/shop")
	blogArticleFramentPath = fragment_path("/blog/{permalink}")
)

func main() {
	app(
		name("Simple App"),
		features(
			logger(),
		),
		web(
			layout(
				AppLayout(),
			),
			web_addons(
				_htmx("2.0.7"),
				_modern_normalize(),
			),

			fragment_at(blogArticleFramentPath,
				fragment(BlogArticleFragment),
			),

			style_global(css{
				"html": style(
					bg_slate(900),
					dark(
						bg_slate(200),
					),
				),
				"div": style(
					dark(
						text_red(400),
						hover(
							text_blue(400),
						),
					),
				),
			}),

			// page
			page_at(homePath,
				title_("Home Page"),
				component(HomeLayout),

				page_at(blogPath,
					title_("The New Blog"),
					component(BlogHomePage),

					page_at(blogAboutPath,
						title_("Blog About"),
						component(BlogAboutPage),
					),
				),
			),
			page_at(shopPath,
				title_("Shop Page"),
				component(ShopLayout),
			),
		),
		port(8080),
	)

}
