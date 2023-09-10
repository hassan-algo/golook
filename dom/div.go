package dom

import (
	"fmt"

	debugger "github.com/hassan-algo/golook/Debugger"
)

type Div struct {
	ClassName string
	Id        string
	Style     string
	Title     string
	Inner     debugger.Debugger
}

func (d Div) Convert() string {
	myDom := fmt.Sprintf("<div id=\"%s\" class=\"%s\" style=\"%s\">", d.Id, d.ClassName, d.Style)
	if d.Inner != nil {
		myDom += fmt.Sprintf("%s</div>", d.Inner.Convert())
	}
	return myDom
}
