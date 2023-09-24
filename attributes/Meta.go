package attributes

import "fmt"

type Meta struct {
	Name    string
	Content string
}

func (m Meta) Convert() string {
	return fmt.Sprintf("<meta name=\"%s\" content=\"%s\">", m.Name, m.Content)
}
