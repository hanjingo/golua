package env

import (
	env "github.com/hanjingo/golib/env"
	glua "github.com/yuin/gopher-lua"
)

func GetCurrPath(state *glua.LState) int {
	path := env.GetCurrPath()
	state.Push(glua.LString(path))
	return 1
}

func GetOsType(state *glua.LState) int {
	typ := env.GetOsType()
	state.Push(glua.LString(typ))
	return 1
}
