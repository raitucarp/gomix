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
		webpage(
			layout(
				AppLayout(),
			),
			web_addons(
				_htmx("2.0.7"),
			),
			fragment(blogArticleFramentPath, BlogArticle),

			// page
			at(homePath, HomePage),
			at(shopPath, ShopPage),
		),
		port(8080),
	)

}
