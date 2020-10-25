package logger

import (
	glua "github.com/yuin/gopher-lua"
)

func Preload(state *glua.LState) {
	state.PreloadModule("logger", Loader)
}

func Loader(state *glua.LState) int {

	logger_ud := state.NewTypeMetatable(`logger_ud`)
	state.SetGlobal(`logger_ud`, logger_ud)
	state.SetField(logger_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"Fatal":   Fatal,
		"Error":   Error,
		"Warning": Warning,
		"Notice":  Notice,
		"Debug":   Debug,
		"Info":    Info,
		"Write":   Write,
	}))

	tb := state.NewTable()
	state.SetFuncs(tb, api)
	state.Push(tb)
	return 1
}

var api = map[string]glua.LGFunction{
	"GetDefaultLogger": GetDefaultLogger,
}
