local time = require("time")

local now = time.Now()
print("now:", now:String())
time.Sleep(time.Second * 1)
print("after:", time.Now():String())
print("millsec:", time.Millisecond)
print("nanosec:", time.Nanosecond)
print("sec:", time.Second)
print("minute:", time.Minute)
print("hour:", time.Hour)