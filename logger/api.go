package logger

import (
	log "github.com/hanjingo/golib/logger"
	glua "github.com/yuin/gopher-lua"
)

func GetDefaultLogger(state *glua.LState) int {
	lg := log.GetDefaultLogger()
	data := state.NewUserData()
	data.Value = lg
	state.SetMetatable(data, state.GetTypeMetatable(`logger_ud`))
	state.Push(data)
	return 1
}

func Write(state *glua.LState) int {
	ud := state.CheckUserData(1)
	log, ok := ud.Value.(*log.Logger)
	if !ok {
		return 0
	}
	lvl := state.CheckInt(2)
	fmt_str := state.CheckString(3)
	var v []interface{}
	for i := 4; i < state.GetTop()+1; i++ {
		v = append(v, state.CheckAny(i))
	}
	log.Write(lvl, fmt_str, v...)
	return 0
}

func Fatal(state *glua.LState) int {
	doLog(state, "fatal")
	return 0
}

func Error(state *glua.LState) int {
	doLog(state, "error")
	return 0
}

func Warning(state *glua.LState) int {
	doLog(state, "warning")
	return 0
}

func Notice(state *glua.LState) int {
	doLog(state, "notice")
	return 0
}

func Debug(state *glua.LState) int {
	doLog(state, "debug")
	return 0
}

func Info(state *glua.LState) int {
	doLog(state, "info")
	return 0
}

func doLog(state *glua.LState, action string) {
	ud := state.CheckUserData(1)
	log, ok := ud.Value.(*log.Logger)
	if !ok {
		return
	}
	fmt_str := state.CheckString(2)
	var v []interface{}
	for i := 3; i < state.GetTop()+1; i++ {
		v = append(v, state.CheckAny(i))
	}
	switch action {
	case "fatal":
		log.Debug(fmt_str, v...)
	case "error":
		log.Error(fmt_str, v...)
	case "warning":
		log.Warning(fmt_str, v...)
	case "notice":
		log.Notice(fmt_str, v...)
	case "debug":
		log.Debug(fmt_str, v...)
	case "info":
		log.Info(fmt_str, v...)
	}
}
