local network = require("network")
local sync = require("sync")

local addr = "127.0.0.1:10086"
function NewConn(s)
    print("serv conn new")
    local err = s:Run()
    if err then
        print("启动失败:", err)
    end
end

function Handle(s, n)
    print("serv conn handle")
    local data, err = s:ReadMsg(10)
    if err then
        print(err)
    end
    print("收到:", data)
end

function ConnClose(s)
    print("serv conn close")
end

local tcpServer = network.NewTcpServ(addr, "NewConn", "Handle", "ConnClose")
local wg = sync.NewWaitGroup()
tcpServer:Run(wg)
wg:Wait()