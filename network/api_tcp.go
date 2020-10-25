package network

import (
	network "github.com/hanjingo/golib/network"
	glua "github.com/yuin/gopher-lua"
)

//tcp conn
func TcpRun(state *glua.LState) int {
	conn := getTcpConn(state, 1)
	if err := conn.Run(); err != nil {
		state.Push(glua.LString(err.Error()))
		return 1
	}
	return 0
}

func TcpDestroy(state *glua.LState) int {
	conn := getTcpConn(state, 1)
	conn.Destroy()
	return 0
}

func TcpReadMsg(state *glua.LState) int {
	conn := getTcpConn(state, 1)
	length := state.CheckInt(2)
	bytes := make([]byte, length)
	n, err := conn.ReadMsg(bytes)
	if err != nil {
		state.Push(glua.LNil)
		state.Push(glua.LString(err.Error()))
		return 2
	}
	state.Push(glua.LString(bytes[:n]))
	return 1
}

func TcpWriteMsg(state *glua.LState) int {
	conn := getTcpConn(state, 1)
	datas := [][]byte{}
	for i := 2; i < state.GetTop()+1; i++ {
		datas = append(datas, []byte(state.CheckString(i)))
	}
	if _, err := conn.WriteMsg(datas...); err != nil {
		state.Push(glua.LString(err.Error()))
		return 1
	}
	return 0
}

func TcpClose(state *glua.LState) int {
	conn := getTcpConn(state, 1)
	conn.Close()
	return 0
}

func TcpIsValid(state *glua.LState) int {
	conn := getTcpConn(state, 1)
	state.Push(glua.LBool(conn.IsValid()))
	return 1
}

func TcpSetParam(state *glua.LState) int {
	conn := getTcpConn(state, 1)
	key := state.CheckString(2)
	value := state.CheckAny(3)
	if err := conn.SetParam(key, value); err != nil {
		state.Push(glua.LString(err.Error()))
		return 1
	}
	return 0
}

//tcp服务器
//新建tcp服务器
func NewTcpServ(state *glua.LState) int {
	addr := state.CheckString(1)
	fname1 := state.CheckString(2)
	fname2 := state.CheckString(3)
	fname3 := state.CheckString(4)
	newConn := func(s network.SessionI) {
		data := state.NewUserData()
		data.Value = s
		state.CallByParam(glua.P{
			Fn:      state.GetGlobal(fname1),
			NRet:    0,
			Protect: true,
		}, data)
	}
	handle := func(s network.SessionI, n int) {
		data := state.NewUserData()
		data.Value = s
		state.CallByParam(glua.P{
			Fn:      state.GetGlobal(fname2),
			NRet:    0,
			Protect: true,
		}, data, glua.LNumber(n))
	}
	connClose := func(s network.SessionI) {
		data := state.NewUserData()
		data.Value = s
		state.CallByParam(glua.P{
			Fn:      state.GetGlobal(fname3),
			NRet:    0,
			Protect: true,
		}, data)
	}
	serv, err := network.NewTcpServer(addr, newConn, handle, connClose)
	if err != nil {
		state.Push(glua.LNil)
		state.Push(glua.LString(err.Error()))
		return 2
	}
	data := state.NewUserData()
	data.Value = serv
	state.SetMetatable(data, state.GetTypeMetatable(`tcp_serv_ud`))
	state.Push(data)
	return 1
}

func TcpServRun(state *glua.LState) int {
	serv := getTcpServ(state, 1)
	wg := getWaitGroup(state, 2)
	serv.Run(wg)
	return 0
}

func TcpServClose(state *glua.LState) int {
	serv := getTcpServ(state, 1)
	serv.Close()
	return 0
}

func TcpServSetParam(state *glua.LState) int {
	serv := getTcpServ(state, 1)
	key := state.CheckString(2)
	value := state.CheckAny(3)
	if err := serv.SetParam(key, value); err != nil {
		state.Push(glua.LString(err.Error()))
		return 1
	}
	return 0
}

//tcp客户端
//新建tcp客户端
func NewTcpCli(state *glua.LState) int {
	addr := state.CheckString(1)
	fname1 := state.CheckString(2)
	fname2 := state.CheckString(3)
	fname3 := state.CheckString(4)
	newConn := func(s network.SessionI) {
		data := state.NewUserData()
		data.Value = s
		state.CallByParam(glua.P{
			Fn:      state.GetGlobal(fname1),
			NRet:    0,
			Protect: true,
		}, data)
	}
	handle := func(s network.SessionI, n int) {
		data := state.NewUserData()
		data.Value = s
		state.CallByParam(glua.P{
			Fn:      state.GetGlobal(fname2),
			NRet:    0,
			Protect: true,
		}, data, glua.LNumber(n))
	}
	connClose := func(s network.SessionI) {
		data := state.NewUserData()
		data.Value = s
		state.CallByParam(glua.P{
			Fn:      state.GetGlobal(fname3),
			NRet:    0,
			Protect: true,
		}, data)
	}
	cli, err := network.NewTcpCli(addr, newConn, handle, connClose)
	if err != nil {
		state.Push(glua.LNil)
		state.Push(glua.LString(err.Error()))
		return 2
	}
	data := state.NewUserData()
	data.Value = cli
	state.SetMetatable(data, state.GetTypeMetatable(`tcp_cli_ud`))
	state.Push(data)
	return 1
}

//tcp客户端拨号
func TcpCliDial(state *glua.LState) int {
	cli := getTcpCli(state, 1)
	addr := state.CheckString(2)
	conn, err := cli.Dial(addr)
	if err != nil {
		state.Push(glua.LNil)
		state.Push(glua.LString(err.Error()))
		return 2
	}
	data := state.NewUserData()
	data.Value = conn
	state.SetMetatable(data, state.GetTypeMetatable(`tcp_conn_ud`))
	state.Push(data)
	return 1
}

//tcp客户端设置参数
func TcpCliSetParam(state *glua.LState) int {
	cli := getTcpCli(state, 1)
	key := state.CheckString(2)
	value := state.CheckAny(3)
	if err := cli.SetParam(key, value); err != nil {
		state.Push(glua.LString(err.Error()))
		return 1
	}
	return 0
}
