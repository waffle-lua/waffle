local http = require('http')
local json = require('json')
local waffle = require('waffle')

print(waffle.version(), waffle.date())

local srv, err = http.server('localhost:8080')
if err then
  error(err)
end

while true do
  local _, resp = srv:accept()

  resp:code(200)
  resp:write(json.encode("a"))
  resp:done()
end
