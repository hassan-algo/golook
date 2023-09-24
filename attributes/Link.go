package attributes

import "fmt"

type Link struct {
	Rel  string
	Href string
	Type string
}

func (l Link) Convert() string {
	return fmt.Sprintf("<link rel=\"%s\" href=\"%s\" type=\"%s\">", l.Rel, l.Href, l.Type)
}
