package network

import (
	"sync"

	network "github.com/hanjingo/golib/network"
	glua "github.com/yuin/gopher-lua"
)

func getTcpConn(state *glua.LState, n int) *network.TcpConn {
	ud := state.CheckUserData(n)
	conn, ok := ud.Value.(*network.TcpConn)
	if ok {
		return conn
	}
	state.ArgError(n, "无法读取tcp连接")
	return nil
}

func getTcpServ(state *glua.LState, n int) *network.TcpServer {
	ud := state.CheckUserData(n)
	conn, ok := ud.Value.(*network.TcpServer)
	if ok {
		return conn
	}
	state.ArgError(n, "无法读取tcp服务器")
	return nil
}

func getTcpCli(state *glua.LState, n int) *network.TcpCli {
	ud := state.CheckUserData(n)
	conn, ok := ud.Value.(*network.TcpCli)
	if ok {
		return conn
	}
	state.ArgError(n, "无法读取tcp客户端")
	return nil
}

func getWsConn(state *glua.LState, n int) *network.WsConn {
	ud := state.CheckUserData(n)
	conn, ok := ud.Value.(*network.WsConn)
	if ok {
		return conn
	}
	state.ArgError(n, "无法读取ws连接")
	return nil
}

func getWsServ(state *glua.LState, n int) *network.WsServer {
	ud := state.CheckUserData(n)
	conn, ok := ud.Value.(*network.WsServer)
	if ok {
		return conn
	}
	state.ArgError(n, "无法读取ws服务器")
	return nil
}

func getWsCli(state *glua.LState, n int) *network.WsCli {
	ud := state.CheckUserData(n)
	conn, ok := ud.Value.(*network.WsCli)
	if ok {
		return conn
	}
	state.ArgError(n, "无法读取ws客户端")
	return nil
}

func getWaitGroup(state *glua.LState, n int) *sync.WaitGroup {
	ud := state.CheckUserData(n)
	wg, ok := ud.Value.(*sync.WaitGroup)
	if ok {
		return wg
	}
	state.ArgError(n, "无法读取wait group")
	return nil
}
