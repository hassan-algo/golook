package main

import (
	"github.com/hassan-algo/golook/components"
	"github.com/hassan-algo/golook/page"
	"github.com/hassan-algo/golook/server"
)

func main() {
	s := server.CreateNewServer()

	s.CreatePage("/", page.CreatePage(components.Home()))
	s.CreatePage("/about", page.CreatePage(components.About()))

	s.Start(":8080")
}
