package logger

import (
	glua "github.com/yuin/gopher-lua"
)

func GetDefaultLoggerTest() {
	state := glua.NewState()
	defer state.Close()
	state.PreloadModule("logger", Loader)
	if err := state.DoFile("logger.lua"); err != nil {
		panic(err)
	}
}
