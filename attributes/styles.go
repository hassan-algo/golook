package attributes

import "fmt"

type Style struct {
	Name  string
	Value string
}

func (s Style) GetStyles() string {
	return fmt.Sprintf("%s:%s,", s.Name, s.Value)
}
