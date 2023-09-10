package dom

import (
	"fmt"
)

type P struct {
	ClassName string
	Id        string
	Style     string
	Title     string
	Text      string
}

func (p P) Convert() string {
	myDom := fmt.Sprintf("<p class=\"%s\" style=\"%s\">%s</p>", p.ClassName, p.Style, p.Text)
	return myDom
}
