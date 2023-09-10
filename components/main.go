package components

import (
	debugger "github.com/hassan-algo/golook/Debugger"
	dom "github.com/hassan-algo/golook/dom"
)

func Home() debugger.Debugger {
	return dom.Div{
		Id: "nais",
		Inner: dom.Div{
			ClassName: "nais",
			Style:     "background-color:red",
			Inner: dom.P{
				Text: "hello world",
			},
		},
	}
}
