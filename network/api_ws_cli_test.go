package main

import (
	network "github.com/hanjingo/go_for_lua/network"
	sync "github.com/hanjingo/go_for_lua/sync"
	time "github.com/hanjingo/go_for_lua/time"
	glua "github.com/yuin/gopher-lua"
)

func main() {
	state := glua.NewState()
	defer state.Close()
	state.PreloadModule("sync", sync.Loader)
	state.PreloadModule("time", time.Loader)
	state.PreloadModule("network", network.Loader)
	if err := state.DoFile("network.lua"); err != nil {
		panic(err)
	}
}
