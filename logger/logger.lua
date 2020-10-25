local log = require("logger")

local deflog = log.GetDefaultLogger()
deflog:Fatal("fatal!!!")
deflog:Error("error!!!")
deflog:Warning("warning!!!")
deflog:Notice("notice!!!")
deflog:Debug("debug!!!")
deflog:Info("info!!!")