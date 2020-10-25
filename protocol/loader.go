package protocol

import (
	glua "github.com/yuin/gopher-lua"
)

func Preload(state *glua.LState) {
	state.PreloadModule("protocol", Loader)
}

func Loader(state *glua.LState) int {

	//协议1
	msg1 := state.NewTypeMetatable("Message1_ud")
	state.SetGlobal("Message1_ud", msg1)
	//编解码器1
	codec1_ud := state.NewTypeMetatable(`codec1_ud`)
	state.SetGlobal(`codec1_ud`, codec1_ud)
	state.SetField(codec1_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"ParseId":      Codec1ParseId,
		"ParseLen":     Codec1ParseLen,
		"ParseContent": Codec1ParseContent,
	}))

	tb := state.NewTable()
	state.SetFuncs(tb, api)
	state.Push(tb)
	return 1
}

var api = map[string]glua.LGFunction{
	"NewCodec1": NewCodec1,
}
