package container

import (
	sm "github.com/hanjingo/golib/container/safe_map"
	glua "github.com/yuin/gopher-lua"
)

func NewSafeMap(state *glua.LState) int {
	data := state.NewUserData()
	data.Value = sm.NewSafeMap()
	state.SetMetatable(data, state.GetTypeMetatable(`safe_map_ud`))
	state.Push(data)
	return 1
}

func Get(state *glua.LState) int {
	return 1
}
