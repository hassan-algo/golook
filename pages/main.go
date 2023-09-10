package pages

import debugger "github.com/hassan-algo/golook/Debugger"

type Page struct {
	Content debugger.Debugger
}

func CreatePages(content debugger.Debugger) *Page {
	return &Page{
		Content: content,
	}
}
