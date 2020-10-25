package time

import glua "github.com/yuin/gopher-lua"

func TmAfter(state *glua.LState) int {
	tm1 := getTime(state, 1)
	tm2 := getTime(state, 2)
	ok := tm1.After(tm2)
	state.Push(glua.LBool(ok))
	return 1
}

func TmAdd(state *glua.LState) int {
	tm1 := getTime(state, 1)
	dur := getDur(state, 2)
	tm2 := tm1.Add(dur)
	data := state.NewUserData()
	data.Value = tm2
	state.SetMetatable(data, state.GetTypeMetatable("time_ud"))
	state.Push(data)
	return 1
}

func TmSub(state *glua.LState) int {
	tm1 := getTime(state, 1)
	tm2 := getTime(state, 2)
	dur := tm1.Sub(tm2)
	state.Push(glua.LNumber(dur))
	return 1
}

func TmString(state *glua.LState) int {
	tm := getTime(state, 1)
	state.Push(glua.LString(tm.String()))
	return 1
}

func TmUTC(state *glua.LState) int {
	tm := getTime(state, 1)
	data := state.NewUserData()
	data.Value = tm.UTC()
	state.SetMetatable(data, state.GetTypeMetatable("time_ud"))
	state.Push(data)
	return 1
}

func TmUnix(state *glua.LState) int {
	tm := getTime(state, 1)
	state.Push(glua.LNumber(tm.Unix()))
	return 1
}

func TmUnixNano(state *glua.LState) int {
	tm := getTime(state, 1)
	state.Push(glua.LNumber(tm.UnixNano()))
	return 1
}

func TmWeekday(state *glua.LState) int {
	tm := getTime(state, 1)
	state.Push(glua.LNumber(tm.Weekday()))
	return 1
}

func TmYear(state *glua.LState) int {
	tm := getTime(state, 1)
	state.Push(glua.LNumber(tm.Year()))
	return 1
}

func TmYearDay(state *glua.LState) int {
	tm := getTime(state, 1)
	state.Push(glua.LNumber(tm.YearDay()))
	return 1
}

func TmDay(state *glua.LState) int {
	tm := getTime(state, 1)
	state.Push(glua.LNumber(tm.Day()))
	return 1
}

func TmDate(state *glua.LState) int {
	tm := getTime(state, 1)
	year, month, day := tm.Date()
	state.Push(glua.LNumber(year))
	state.Push(glua.LString(month.String()))
	state.Push(glua.LNumber(day))
	return 3
}
