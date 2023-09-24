package components

import (
	"github.com/hassan-algo/golook/debugger"
	"github.com/hassan-algo/golook/dom"
)

func Contact() debugger.Debugger {
	return dom.Body{
		ClassName: "bg-red-500",
		Inner: []debugger.Debugger{
			&dom.Div{
				Inner: []debugger.Debugger{
					&dom.A{
						Href: "/contact",
						Inner: []debugger.Debugger{
							&dom.P{
								ClassName: "text-xl",
								Text:      "Contact PAGE",
							},
						},
					},
					privateComponent(),
				},
			},
		},
	}
}

func privateComponent() debugger.Debugger {
	return dom.A{
		Href: "/",
		Inner: []debugger.Debugger{
			&dom.P{
				ClassName: "text-sm",
				Text:      "Home",
			},
		},
	}
}
