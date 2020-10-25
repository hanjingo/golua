package time

import (
	"time"

	glua "github.com/yuin/gopher-lua"
)

func Preload(state *glua.LState) {
	state.PreloadModule("time", Loader)
}

func Loader(state *glua.LState) int {

	timer_ud := state.NewTypeMetatable("timer_ud")
	state.SetGlobal("timer_ud", timer_ud)
	state.SetField(timer_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"Reset": TimerReset,
		"Stop":  TimerStop,
	}))

	ticker_ud := state.NewTypeMetatable("ticker_ud")
	state.SetGlobal("ticker_ud", ticker_ud)
	state.SetField(ticker_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"Stop": TickerStop,
	}))

	time_ud := state.NewTypeMetatable("time_ud")
	state.SetGlobal("time_ud", time_ud)
	state.SetField(time_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"After":    TmAfter,
		"Add":      TmAdd,
		"String":   TmString,
		"UTC":      TmUTC,
		"Unix":     TmUnix,
		"UnixNano": TmUnixNano,
		"Weekday":  TmWeekday,
		"Year":     TmYear,
		"YearDay":  TmYearDay,
		"Day":      TmDay,
		"Date":     TmDate,
	}))

	tb := state.NewTable()
	state.SetFuncs(tb, api)
	state.SetField(tb, "Millisecond", glua.LNumber(time.Millisecond))
	state.SetField(tb, "Nanosecond", glua.LNumber(time.Nanosecond))
	state.SetField(tb, "Second", glua.LNumber(time.Second))
	state.SetField(tb, "Minute", glua.LNumber(time.Minute))
	state.SetField(tb, "Hour", glua.LNumber(time.Hour))
	state.Push(tb)
	return 1
}

var api = map[string]glua.LGFunction{
	"Now":       Now,
	"Sleep":     Sleep,
	"Until":     Until,
	"NewTimer":  NewTimer,
	"NewTicker": NewTicker,
}
