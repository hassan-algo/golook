package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hassan-algo/golook/components"
	debugger "github.com/hassan-algo/golook/debugger"
	"github.com/hassan-algo/golook/pages"
)

func main() {
	http.HandleFunc("/", handleHome)

	pages.CreatePages(components.Home())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homePage := components.Home()

	data := debugger.ConvertDom(homePage)
	fmt.Fprint(w, data)
}
