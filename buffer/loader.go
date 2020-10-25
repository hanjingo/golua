package buffer

import (
	glua "github.com/yuin/gopher-lua"
)

func Preload(state *glua.LState) {
	state.PreloadModule("buffer", Loader)
}

func Loader(state *glua.LState) int {

	buffer_ud := state.NewTypeMetatable(`buffer_ud`)
	state.SetGlobal(`buffer_ud`, buffer_ud)
	state.SetField(buffer_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"Len":   Len,
		"Read":  Read,
		"Write": Write,
	}))

	tb := state.NewTable()
	state.SetFuncs(tb, api)
	state.Push(tb)
	return 1
}

var api = map[string]glua.LGFunction{
	"NewBuffer": NewBuffer,
}
