package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hassan-algo/golook/debugger"
	"github.com/hassan-algo/golook/dom"
)

type Page struct {
	Head    dom.Head
	Content debugger.Debugger
}

func CreatePage(head dom.Head, content debugger.Debugger) *Page {
	return &Page{
		Head:    head,
		Content: content,
	}
}

func (p *Page) PageHandler(c *gin.Context) {
	pageData := p.Content

	data := p.Head.Convert()
	data += debugger.ConvertDom(pageData)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(data))
	return
}
