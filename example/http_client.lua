local http = require('http')
local json = require('json')

local r = http.request('POST', 'https://httpbin.org/post?xyz', json.encode(
  {
    a = "b",
    b = "c",
    d = {
    },
  }
))

local client = http.client({ timeout = 30 })
local resp, err = client:do_request(r)
if err then
  error(err)
end

print(resp.body)
