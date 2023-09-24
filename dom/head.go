package dom

import (
	"fmt"

	"github.com/hassan-algo/golook/attributes"
)

type Head struct {
	Charset string
	Title   string
	Base    string
	Links   []attributes.Link
	Scripts []attributes.Script
	Metas   []attributes.Meta
}

func (h Head) Convert() string {

	str := fmt.Sprintf("<head>")
	str += fmt.Sprintf("<meta charset=\"%s\">", h.Charset)
	str += fmt.Sprintf("<title>%s</title>", h.Title)
	str += fmt.Sprintf("%s", h.convertMetas())
	str += fmt.Sprintf("%s", h.convertLinks())
	str += fmt.Sprintf("%s", h.convertScripts())
	str += fmt.Sprintf("<base href=\"%s\">", h.Base)
	str += fmt.Sprintf("</head>")
	return str
}

func (h Head) convertMetas() string {
	str := ""
	for _, meta := range h.Metas {
		str += fmt.Sprintf("%s", meta.Convert())
	}
	return str
}

func (h Head) convertLinks() string {
	str := ""
	for _, link := range h.Links {
		str += fmt.Sprintf("%s", link.Convert())
	}
	return str
}

func (h Head) convertScripts() string {
	str := ""
	for _, script := range h.Scripts {
		str += fmt.Sprintf("%s", script.Convert())
	}
	return str
}
