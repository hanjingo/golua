package buffer

import (
	"bytes"

	glua "github.com/yuin/gopher-lua"
)

func NewBuffer(state *glua.LState) int {
	buf := new(bytes.Buffer)
	data := state.NewUserData()
	data.Value = buf
	state.SetMetatable(data, state.GetTypeMetatable(`buffer_ud`))
	state.Push(data)
	return 1
}

func Len(state *glua.LState) int {
	buf := getBuf(state)
	if buf == nil {
		return 0
	}
	state.Push(glua.LNumber(buf.Len()))
	return 1
}

func Read(state *glua.LState) int {
	buf := getBuf(state)
	if buf == nil {
		return 0
	}
	n := state.CheckInt(2)
	data := make([]byte, n)
	length, err := buf.Read(data)
	if err != nil {
		return 0
	}
	state.Push(glua.LString(string(data[:length])))
	return 1
}

func Write(state *glua.LState) int {
	buf := getBuf(state)
	if buf == nil {
		return 0
	}
	str := state.CheckString(2)
	n, err := buf.Write([]byte(str))
	if err != nil {
		return 0
	}
	state.Push(glua.LNumber(n))
	return 1
}

func getBuf(state *glua.LState) *bytes.Buffer {
	ud := state.CheckUserData(1)
	if buf, ok := ud.Value.(*bytes.Buffer); ok {
		return buf
	}
	return nil
}
