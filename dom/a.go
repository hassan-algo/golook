package dom

import (
	"fmt"

	"github.com/hassan-algo/golook/attributes"
	"github.com/hassan-algo/golook/debugger"
)

type A struct {
	ClassName string
	Id        string
	Style     []attributes.Style
	Title     string
	Href      string
	Inner     []debugger.Debugger
}

func (a A) Convert() string {
	myDom := fmt.Sprintf("<a class=\"%s\" href=\"%s\" style=\"%s\">", a.ClassName, a.Href, a.GetStyles())
	if a.Inner != nil {
		for _, myInner := range a.Inner {
			myDom += fmt.Sprintf("%s", myInner.Convert())
		}
	}
	myDom += fmt.Sprintf("</a>")
	return myDom
}

func (a A) GetStyles() string {
	style := ""
	for _, s := range a.Style {
		style += s.GetStyles()
	}
	if style != "" {
		return style[:len(style)-1]
	}
	return style
}
