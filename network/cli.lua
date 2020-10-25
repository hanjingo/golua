local network = require("network")
local sync = require("sync")
local time = require("time")

local addr = "127.0.0.1:10087"
function NewConn(s)
    print("cli new conn")
    local err = s:Run()
    if err then
        print("启动失败:", err)
    end
end

function Handle(s, n)
    print("cli handle")
end

function ConnClose(s)
    print("cli close")
end

local tcpCli = network.NewTcpCli(addr, "NewConn", "Handle", "ConnClose")
local conn, err = tcpCli:Dial("127.0.0.1:10086")
if err then
    print(err)
    return
end
time.Sleep(time.Millisecond * 500)
local str = "abc"
print("write:", str)
conn:WriteMsg(str)
time.Sleep(time.Second * 3)