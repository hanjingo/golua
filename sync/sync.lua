local sync = require("sync")

local wg = sync.NewWaitGroup()
wg:Add(2)
wg:Done()
wg:Done()
wg:Wait()
