local uuid = require("uuid")

local gen = uuid.GetUuidGenerator()
print("uuid:", gen:GenUuid64())