package components

import (
	"fmt"

	"github.com/hassan-algo/golook/attributes"
	"github.com/hassan-algo/golook/debugger"
	"github.com/hassan-algo/golook/dom"
)

func Home() debugger.Debugger {
	fmt.Println("HOME")
	return dom.Body{
		Style: []attributes.Style{
			{
				"background-color", "darkgray",
			},
		},
		Inner: []debugger.Debugger{
			&dom.Div{
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
						ClassName: "nais2",
						Style: []attributes.Style{
							{
								"background-color", "black",
							},
						},
						Inner: []debugger.Debugger{
							myName(),
						},
					},
					&dom.Div{
						ClassName: "links",
						Inner: []debugger.Debugger{
							&dom.A{
								Href: "./about",
								Text: "About",
							},
						},
					},
				},
			},
		},
	}
}

func myName() debugger.Debugger {
	return &dom.P{
		Text: "Hassan Anwar",
		Style: []attributes.Style{
			{
				"background-color", "yellow",
			},
		},
	}
}
