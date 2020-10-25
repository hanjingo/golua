package uuid

import (
	glua "github.com/yuin/gopher-lua"
)

func Preload(state *glua.LState) {
	state.PreloadModule("uuid", Loader)
}

func Loader(state *glua.LState) int {

	uuid_ud := state.NewTypeMetatable(`uuid_ud`)
	state.SetGlobal(`uuid_ud`, uuid_ud)
	state.SetField(uuid_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"GenUuid64": GenUuid64,
	}))

	tb := state.NewTable()
	state.SetFuncs(tb, api)
	state.Push(tb)
	return 1
}

var api = map[string]glua.LGFunction{
	"GetUuidGenerator": GetUuidGenerator,
}
