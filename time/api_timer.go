package time

import (
	"time"

	glua "github.com/yuin/gopher-lua"
)

func NewTimer(state *glua.LState) int {
	dur := getDur(state, 1)
	data := state.NewUserData()
	data.Value = time.NewTimer(dur)
	state.SetMetatable(data, state.GetTypeMetatable("timer_ud"))
	state.Push(data)
	return 1
}

func TimerReset(state *glua.LState) int {
	tm := getTimer(state, 1)
	dur := getDur(state, 2)
	state.Push(glua.LBool(tm.Reset(dur)))
	return 1
}

func TimerStop(state *glua.LState) int {
	tm := getTimer(state, 1)
	state.Push(glua.LBool(tm.Stop()))
	return 1
}
