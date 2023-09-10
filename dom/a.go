package dom

import (
	"fmt"

	"github.com/hassan-algo/golook/attributes"
)

type A struct {
	ClassName string
	Id        string
	Style     []attributes.Style
	Title     string
	Text      string
	Href      string
}

func (a A) Convert() string {
	myDom := fmt.Sprintf("<a class=\"%s\" href=\"%s\" style=\"%s\">%s</a>", a.ClassName, a.Href, a.Style, a.Text)
	return myDom
}

func (a A) GetStyles() string {
	style := ""
	for _, s := range a.Style {
		style += s.GetStyles()
	}
	return style[:len(style)-1]
}
