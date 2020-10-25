package time

import (
	"time"

	glua "github.com/yuin/gopher-lua"
)

func NewTicker(state *glua.LState) int {
	dur := getDur(state, 1)
	data := state.NewUserData()
	data.Value = time.NewTicker(dur)
	state.SetMetatable(data, state.GetTypeMetatable("ticker_ud"))
	state.Push(data)
	return 1
}

func TickerStop(state *glua.LState) int {
	tk := getTicker(state, 1)
	tk.Stop()
	return 0
}
