package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hassan-algo/golook/components"
	"github.com/hassan-algo/golook/page"
)

func main() {
	r := gin.Default()

	home := page.CreatePage(components.Home())
	r.GET("/", home.PageHandler)

	about := page.CreatePage(components.About())
	r.GET("/about", about.PageHandler)

	r.Run(":8080")
}
