package time

import (
	"time"

	glua "github.com/yuin/gopher-lua"
)

func Now(state *glua.LState) int {
	now := time.Now()
	data := state.NewUserData()
	data.Value = now
	state.SetMetatable(data, state.GetTypeMetatable("time_ud"))
	state.Push(data)
	return 1
}

func Sleep(state *glua.LState) int {
	dur := getDur(state, 1)
	time.Sleep(dur)
	return 0
}

func Until(state *glua.LState) int {
	tm := getTime(state, 1)
	dur := time.Until(tm)
	state.Push(glua.LNumber(dur))
	return 1
}
