package protocol

import (
	v1 "github.com/hanjingo/golib/protocol/v1"
	glua "github.com/yuin/gopher-lua"
)

func getCodec1(state *glua.LState, n int) *v1.Codec {
	ud := state.CheckUserData(n)
	codec, ok := ud.Value.(*v1.Codec)
	if ok {
		return codec
	}
	state.ArgError(n, "获得codec失败")
	return nil
}

func NewCodec1(state *glua.LState) int {
	codec := v1.NewCodec()
	data := state.NewUserData()
	data.Value = codec
	state.SetMetatable(data, state.GetTypeMetatable("codec1_ud"))
	state.Push(data)
	return 1
}

func Codec1ParseId(state *glua.LState) int {
	codec := getCodec1(state, 1)
	str := state.CheckString(2)
	id, err := codec.ParseId([]byte(str))
	if err != nil {
		state.Push(glua.LNil)
		state.Push(glua.LString(err.Error()))
		return 2
	}
	state.Push(glua.LNumber(id))
	return 1
}

func Codec1ParseLen(state *glua.LState) int {
	codec := getCodec1(state, 1)
	str := state.CheckString(2)
	length, err := codec.ParseLen([]byte(str))
	if err != nil {
		state.Push(glua.LNil)
		state.Push(glua.LString(err.Error()))
		return 2
	}
	state.Push(glua.LNumber(length))
	return 1
}

func Codec1ParseContent(state *glua.LState) int {
	codec := getCodec1(state, 1)
	str := state.CheckString(2)
	content, err := codec.ParseContent([]byte(str))
	if err != nil {
		state.Push(glua.LNil)
		state.Push(glua.LString(err.Error()))
		return 2
	}
	state.Push(glua.LString(content))
	return 1
}
