package main

import (
	network "github.com/hanjingo/go_for_lua/network"
	sync "github.com/hanjingo/go_for_lua/sync"
	glua "github.com/yuin/gopher-lua"
)

func main() {
	state := glua.NewState()
	defer state.Close()
	state.PreloadModule("sync", sync.Loader)
	state.PreloadModule("network", network.Loader)
	if err := state.DoFile("network.lua"); err != nil {
		panic(err)
	}
}
