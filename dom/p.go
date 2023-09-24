package dom

import (
	"fmt"

	"github.com/hassan-algo/golook/attributes"
)

type P struct {
	ClassName string
	Id        string
	Style     []attributes.Style
	Title     string
	Text      string
}

func (p P) Convert() string {
	myDom := fmt.Sprintf("<p class=\"%s\" style=\"%s\">%s</p>", p.ClassName, p.GetStyles(), p.Text)
	return myDom
}

func (p P) GetStyles() string {
	style := ""
	for _, s := range p.Style {
		style += s.GetStyles()
	}
	if style == "" {
		return style
	}
	return style[:len(style)-1]
}
