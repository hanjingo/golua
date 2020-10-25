package sync

import (
	"sync"

	glua "github.com/yuin/gopher-lua"
)

func getWg(state *glua.LState, n int) *sync.WaitGroup {
	ud := state.CheckUserData(n)
	wg, ok := ud.Value.(*sync.WaitGroup)
	if ok {
		return wg
	}
	state.ArgError(n, "无法读取WaitGroup")
	return nil
}

func NewWaitGroup(state *glua.LState) int {
	data := state.NewUserData()
	data.Value = new(sync.WaitGroup)
	state.SetMetatable(data, state.GetTypeMetatable(`wg_ud`))
	state.Push(data)
	return 1
}

func WgAdd(state *glua.LState) int {
	if wg := getWg(state, 1); wg != nil {
		n := state.CheckInt(2)
		wg.Add(n)
	}
	return 0
}

func WgDone(state *glua.LState) int {
	if wg := getWg(state, 1); wg != nil {
		wg.Done()
	}
	return 0
}

func WgWait(state *glua.LState) int {
	if wg := getWg(state, 1); wg != nil {
		wg.Wait()
	}
	return 0
}
