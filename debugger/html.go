package debugger

type Debugger interface {
	Convert() string
}

func ConvertDom(debugger Debugger) string {
	return debugger.Convert()
}
