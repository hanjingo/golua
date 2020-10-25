package network

import (
	glua "github.com/yuin/gopher-lua"
)

func Preload(state *glua.LState) {
	state.PreloadModule("network", Loader)
}

func Loader(state *glua.LState) int {
	//tcp conn
	tcp_conn_ud := state.NewTypeMetatable(`tcp_conn_ud`)
	state.SetGlobal(`tcp_conn_ud`, tcp_conn_ud)
	state.SetField(tcp_conn_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"Run":      TcpRun,
		"Destroy":  TcpDestroy,
		"ReadMsg":  TcpReadMsg,
		"WriteMsg": TcpWriteMsg,
		"Close":    TcpClose,
		"IsValid":  TcpIsValid,
		"SetParam": TcpSetParam,
	}))
	//tcp serv
	tcp_serv_ud := state.NewTypeMetatable(`tcp_serv_ud`)
	state.SetGlobal(`tcp_serv_ud`, tcp_serv_ud)
	state.SetField(tcp_serv_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"Run":      TcpServRun,
		"Close":    TcpServClose,
		"SetParam": TcpServSetParam,
	}))
	//tcp cli
	tcp_cli_ud := state.NewTypeMetatable(`tcp_cli_ud`)
	state.SetGlobal(`tcp_cli_ud`, tcp_cli_ud)
	state.SetField(tcp_cli_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"Dial":    TcpCliDial,
		"SetParm": TcpSetParam,
	}))

	//ws conn
	ws_conn_ud := state.NewTypeMetatable(`ws_conn_ud`)
	state.SetGlobal(`ws_conn_ud`, ws_conn_ud)
	state.SetField(ws_conn_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"Run":      WsRun,
		"Destroy":  WsDestroy,
		"ReadMsg":  WsReadMsg,
		"WriteMsg": WsWriteMsg,
		"Close":    WsClose,
		"IsValid":  WsIsValid,
		"SetParam": WsSetParam,
	}))
	//ws serv
	ws_serv_ud := state.NewTypeMetatable(`ws_serv_ud`)
	state.SetGlobal(`ws_serv_ud`, ws_serv_ud)
	state.SetField(ws_serv_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"Run":      WsServRun,
		"Close":    WsServClose,
		"SetParam": WsServSetParam,
	}))
	//ws cli
	ws_cli_ud := state.NewTypeMetatable(`ws_cli_ud`)
	state.SetGlobal(`ws_cli_ud`, ws_cli_ud)
	state.SetField(ws_cli_ud, "__index", state.SetFuncs(state.NewTable(), map[string]glua.LGFunction{
		"Dial":    WsCliDial,
		"SetParm": WsSetParam,
	}))

	tb := state.NewTable()
	state.SetFuncs(tb, api)
	state.Push(tb)
	return 1
}

var api = map[string]glua.LGFunction{
	"NewTcpServ": NewTcpServ,
	"NewTcpCli":  NewTcpCli,
	"NewWsServ":  NewWsServ,
	"NewWsCli":   NewWsCli,
}
