package main

import (
	"github.com/hassan-algo/golook/attributes"
	"github.com/hassan-algo/golook/components"
	"github.com/hassan-algo/golook/dom"
	"github.com/hassan-algo/golook/page"
	"github.com/hassan-algo/golook/server"
)

func main() {
	s := server.CreateNewServer()

	myHead := dom.Head{
		Charset: "UTF-8",
		Title:   "GoLook",
		Metas: []attributes.Meta{
			{
				Name:    "author",
				Content: "Hassan Anwar",
			},
		},
		Scripts: []attributes.Script{
			{
				Src: "https://cdn.tailwindcss.com",
			},
		},
	}

	s.CreatePage("/", page.CreatePage(myHead, components.Home()))
	s.CreatePage("/about", page.CreatePage(myHead, components.About()))
	s.CreatePage("/contact", page.CreatePage(myHead, components.Contact()))

	s.Start(":8080")
}
