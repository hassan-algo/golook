package components

import (
	"fmt"

	"github.com/hassan-algo/golook/attributes"
	"github.com/hassan-algo/golook/debugger"
	"github.com/hassan-algo/golook/dom"
)

func About() debugger.Debugger {
	fmt.Println("about")
	return dom.Body{
		Style: []attributes.Style{
			{
				"background-color", "darkgray",
			},
		},
		Inner: []debugger.Debugger{
			dom.Div{
				Id: "nais",
				Inner: []debugger.Debugger{
					&dom.Div{
						ClassName: "nais1",
						Style: []attributes.Style{
							{
								"background-color", "red",
							},
						},
						Inner: []debugger.Debugger{
							&dom.P{
								Text: "hello world",
							},
						},
					},
					&dom.Div{
						ClassName: "links",
						Inner: []debugger.Debugger{
							&dom.A{
								Href: "/",
								Text: "Home",
							},
						},
					},
				},
			},
		},
	}
}
