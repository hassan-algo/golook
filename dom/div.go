package dom

import (
	"fmt"

	"github.com/hassan-algo/golook/attributes"
	"github.com/hassan-algo/golook/debugger"
)

type Div struct {
	ClassName string
	Id        string
	Style     []attributes.Style
	Title     string
	Inner     []debugger.Debugger
}

func (d Div) Convert() string {
	myDom := fmt.Sprintf("<div id=\"%s\" class=\"%s\" style=\"%s\">", d.Id, d.ClassName, d.GetStyles())
	if d.Inner != nil {
		for _, myInner := range d.Inner {
			myDom += fmt.Sprintf("%s", myInner.Convert())
		}
	}
	myDom += fmt.Sprintf("</div>")
	return myDom
}

func (d Div) GetStyles() string {
	style := ""
	for _, s := range d.Style {
		style += s.GetStyles()
	}
	if style == "" {
		return style
	}
	return style[:len(style)-1]
}
