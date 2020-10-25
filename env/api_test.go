package env

import (
	glua "github.com/yuin/gopher-lua"
)

func main() {
	state := glua.NewState()
	defer state.Close()
	state.PreloadModule("env", Loader)
	if err := state.DoFile("env.lua"); err != nil {
		panic(err)
	}
}
