package dom

import (
	"fmt"

	"github.com/hassan-algo/golook/attributes"
	"github.com/hassan-algo/golook/debugger"
)

type Body struct {
	ClassName string
	Id        string
	Style     []attributes.Style
	Title     string
	Inner     []debugger.Debugger
}

func (b Body) Convert() string {
	myDom := fmt.Sprintf("<body class=\"%s\" style=\"%s\">", b.ClassName, b.GetStyles())
	if b.Inner != nil {
		for _, myInner := range b.Inner {
			myDom += fmt.Sprintf("%s", myInner.Convert())
		}
	}
	myDom += fmt.Sprintf("</body>")
	return myDom
}

func (b Body) GetStyles() string {
	style := ""
	for _, s := range b.Style {
		style += s.GetStyles()
	}
	if style == "" {
		return style
	}
	return style[:len(style)-1]
}
