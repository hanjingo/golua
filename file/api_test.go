package file

import (
	glua "github.com/yuin/gopher-lua"
)

func main() {
	state := glua.NewState()
	defer state.Close()
	state.PreloadModule("file", Loader)
	if err := state.DoFile("file.lua"); err != nil {
		panic(err)
	}
}
