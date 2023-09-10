package main

import (
	"log"
	"net/http"

	"github.com/hassan-algo/golook/components"
	"github.com/hassan-algo/golook/pages"
)

func main() {

	home := pages.CreatePages(components.Home())
	http.HandleFunc("/", home.PageHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
