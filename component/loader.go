package component

import (
	glua "github.com/yuin/gopher-lua"
)

func Preload(state *glua.LState) {
	state.PreloadModule("component", Loader)
}

func Loader(state *glua.LState) int {

	//file manager1
	file_manager1_ud := state.NewTypeMetatable(`file_manager1_ud`)
	state.SetGlobal(`file_manager1_ud`, file_manager1_ud)
	state.SetField(file_manager1_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"OpenFile":  Fm1OpenFile,
		"Write":     Fm1Write,
		"CloseFile": Fm1CloseFile,
		"CleanFile": Fm1CleanFile,
	}))

	tb := state.NewTable()
	state.SetFuncs(tb, api)
	state.Push(tb)
	return 1
}

var api = map[string]glua.LGFunction{
	"NewFileManager1": nil,
}
