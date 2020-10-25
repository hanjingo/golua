package main

import (
	time "github.com/hanjingo/go_for_lua/time"
	glua "github.com/yuin/gopher-lua"
)

func main() {
	state := glua.NewState()
	defer state.Close()
	state.PreloadModule("time", time.Loader)
	if err := state.DoFile("time.lua"); err != nil {
		panic(err)
	}
}
