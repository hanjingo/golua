package time

import (
	"time"

	glua "github.com/yuin/gopher-lua"
)

func getTime(state *glua.LState, n int) time.Time {
	ud := state.CheckUserData(n)
	return ud.Value.(time.Time)
}

func getDur(state *glua.LState, n int) time.Duration {
	dur := state.CheckInt64(n)
	return time.Duration(dur)
}

func getTimer(state *glua.LState, n int) *time.Timer {
	ud := state.CheckUserData(n)
	return ud.Value.(*time.Timer)
}

func getTicker(state *glua.LState, n int) *time.Ticker {
	ud := state.CheckUserData(n)
	return ud.Value.(*time.Ticker)
}
