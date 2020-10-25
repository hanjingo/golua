package uuid

import (
	"fmt"

	uuid "github.com/hanjingo/golib/uuid"
	glua "github.com/yuin/gopher-lua"
)

func GetUuidGenerator(state *glua.LState) int {
	gen := uuid.GetUuidGenerator()
	data := state.NewUserData()
	data.Value = gen
	state.SetMetatable(data, state.GetTypeMetatable(`uuid_ud`))
	state.Push(data)
	return 1
}

func GenUuid64(state *glua.LState) int {
	ud := state.CheckUserData(1)
	gen, ok := ud.Value.(*uuid.Uuid)
	if !ok {
		state.Push(glua.LNil)
		state.Push(glua.LString(fmt.Sprintln("获得uuid生成器")))
		return 2
	}
	out := []uint8{}
	for i := 2; i < state.GetTop()+1; i++ {
		v := state.CheckInt(i)
		out = append(out, uint8(v))
	}
	if len(out) > 0 {
		id, err := gen.GenUuid64(out...)
		if err != nil {
			state.Push(glua.LNil)
			state.Push(glua.LString(err.Error()))
			return 2
		}
		state.Push(glua.LNumber(id))
	} else {
		id, err := gen.GenUuid64()
		if err != nil {
			state.Push(glua.LNil)
			state.Push(glua.LString(err.Error()))
			return 2
		}
		state.Push(glua.LNumber(id))
	}
	return 1
}
