package sync

import (
	glua "github.com/yuin/gopher-lua"
)

func Preload(state *glua.LState) {
	state.PreloadModule("sync", Loader)
}

func Loader(state *glua.LState) int {

	wg_ud := state.NewTypeMetatable(`wg_ud`)
	state.SetGlobal(`wg_ud`, wg_ud)
	state.SetField(wg_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"Add":  WgAdd,
		"Done": WgDone,
		"Wait": WgWait,
	}))

	tb := state.NewTable()
	state.SetFuncs(tb, api)
	state.Push(tb)
	return 1
}

var api = map[string]glua.LGFunction{
	"NewWaitGroup": NewWaitGroup,
}
