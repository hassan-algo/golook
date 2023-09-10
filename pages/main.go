package pages

import (
	"fmt"
	"net/http"

	debugger "github.com/hassan-algo/golook/Debugger"
	"github.com/hassan-algo/golook/components"
)

type Page struct {
	Content debugger.Debugger
}

func CreatePages(content debugger.Debugger) *Page {
	return &Page{
		Content: content,
	}
}

func (p *Page) PageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homePage := components.Home()

	data := debugger.ConvertDom(homePage)
	fmt.Fprint(w, data)
}
