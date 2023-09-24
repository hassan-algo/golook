package attributes

import "fmt"

type Script struct {
	Src string
}

func (s Script) Convert() string {
	return fmt.Sprintf("<script src=\"%s\"></script>", s.Src)
}
