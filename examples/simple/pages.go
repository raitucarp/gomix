package main

import "github.com/raitucarp/gomix"

func ShopPage(page *gomix.Page) {
	page.Component(ShopLayout)
	page.Title("Shope page")
}

func HomePage(page *gomix.Page) {
	page.Component(HomeLayout)
	page.Title("This is Home Page")

	page.Children(gomix.Pages{
		page.New(blogPath, BlogPage),
	})
}

func BlogAboutPage(page *gomix.Page) {
	bb := page.Request().PathValue("name")
	page.Title("About Blog" + bb)
}

func BlogPage(blogPage *gomix.Page) {
	blogPage.Component(BlogHomePage)
	blogPage.Title("The New Blog")

	
	blogPage.Children(gomix.Pages{
		blogPage.New(blogAboutPath, BlogAboutPage),
	})
}
