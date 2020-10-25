package main

import (
	sync "github.com/hanjingo/go_for_lua/sync"
	glua "github.com/yuin/gopher-lua"
)

func main() {
	state := glua.NewState()
	defer state.Close()
	state.PreloadModule("sync", sync.Loader)
	if err := state.DoFile("sync.lua"); err != nil {
		panic(err)
	}
}
