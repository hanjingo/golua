package main

import (
	uuid "github.com/hanjingo/go_for_lua/uuid"
	glua "github.com/yuin/gopher-lua"
)

func main() {
	state := glua.NewState()
	defer state.Close()
	state.PreloadModule("uuid", uuid.Loader)
	if err := state.DoFile("uuid.lua"); err != nil {
		panic(err)
	}
}
