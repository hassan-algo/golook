package components

import (
	"github.com/hassan-algo/golook/debugger"
	"github.com/hassan-algo/golook/dom"
)

func About() debugger.Debugger {
	return dom.Body{
		ClassName: "bg-slate-200",
		Inner: []debugger.Debugger{
			dom.Div{
				Id: "nais",
				Inner: []debugger.Debugger{
					&dom.Div{
						ClassName: "nais1 text-blue-400",
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
								Inner: []debugger.Debugger{
									&dom.P{
										Text: "HOME",
									},
								},
							},
						},
					},
					&dom.Div{
						ClassName: "links",
						Inner: []debugger.Debugger{
							&dom.A{
								Href: "/contact",
								Inner: []debugger.Debugger{
									&dom.P{
										Text: "Contact",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
