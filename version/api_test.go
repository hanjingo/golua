package main

import (
	version "github.com/hanjingo/go_for_lua/version"
	glua "github.com/yuin/gopher-lua"
)

func main() {
	state := glua.NewState()
	defer state.Close()
	state.PreloadModule("version", version.Loader)
	if err := state.DoFile("version.lua"); err != nil {
		panic(err)
	}
}
