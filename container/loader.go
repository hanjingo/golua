package container

import (
	glua "github.com/yuin/gopher-lua"
)

const ModName = "Container"
const UdSafeMap = "UdSafeMap"

func Preload(state *glua.LState) {
	state.PreloadModule(ModName, Loader)
}

func Loader(state *glua.LState) int {

	safe_map := state.NewTypeMetatable(UdSafeMap)
	state.SetGlobal(UdSafeMap, safe_map)
	state.SetField(safe_map, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"Get":        Get,
		"Set":        nil,
		"TrySet":     nil,
		"Del":        nil,
		"Len":        nil,
		"RLockRange": nil,
		"LockRange":  nil,
		"Clean":      nil,
	}))

	tb := state.NewTable()
	state.SetFuncs(tb, api)
	state.Push(tb)
	return 1
}

var api = map[string]glua.LGFunction{
	"NewSafeMap": NewSafeMap,
}
