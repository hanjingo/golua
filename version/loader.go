package version

import (
	glua "github.com/yuin/gopher-lua"
)

func Preload(state *glua.LState) {
	state.PreloadModule("version", Loader)
}

func Loader(state *glua.LState) int {
	tb := state.NewTable()
	state.SetFuncs(tb, api)
	state.Push(tb)
	return 1
}

var api = map[string]glua.LGFunction{
	"CmpVersion": CmpVersion,
}
