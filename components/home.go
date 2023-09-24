package components

import (
	"github.com/hassan-algo/golook/attributes"
	"github.com/hassan-algo/golook/debugger"
	"github.com/hassan-algo/golook/dom"
)

func Home() debugger.Debugger {
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
						ClassName: "nais1 text-red-500",
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
								ClassName: "bg-red-500",
								Style: []attributes.Style{
									{
										"text-decoration", "none",
									},
								},
								Href: "./about",
								Inner: []debugger.Debugger{
									&dom.P{
										ClassName: "bg-red-500",
										Text:      "About",
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

func myName() debugger.Debugger {
	return &dom.P{
		ClassName: "bg-yellow-500",
		Text:      "Hassan Anwar",
	}
}
