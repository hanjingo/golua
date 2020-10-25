package component

import (
	"testing"

	glua "github.com/yuin/gopher-lua"
)

func Fm1OpenFileTest(t *testing.T) {
	state := glua.NewState()
	defer state.Close()
	state.PreloadModule("cp", Loader)
	if err := state.DoFile("handler2.lua"); err != nil {
		panic(err)
	}
}
