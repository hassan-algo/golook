package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hassan-algo/golook/debugger"
)

type Page struct {
	Content debugger.Debugger
}

func CreatePage(content debugger.Debugger) *Page {
	return &Page{
		Content: content,
	}
}

func (p *Page) PageHandler(c *gin.Context) {
	pageData := p.Content

	data := debugger.ConvertDom(pageData)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(data))
	return
}
