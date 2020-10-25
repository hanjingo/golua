package version

import (
	version "github.com/hanjingo/golib/version"
	glua "github.com/yuin/gopher-lua"
)

func CmpVersion(state *glua.LState) int {
	v1 := state.CheckString(1)
	v2 := state.CheckString(2)
	result := version.CmpVersion(v1, v2)
	state.Push(glua.LString(result))
	return 1
}
