package buffer

import (
	glua "github.com/yuin/gopher-lua"
)

func NewBufferTest() {
	state := glua.NewState()
	defer state.Close()
	state.PreloadModule("buffer", Loader)
	if err := state.DoFile("lua"); err != nil {
		panic(err)
	}
}
