package luahttp

import (
	"testing"

	"github.com/stretchr/testify/suite"
	lua "github.com/yuin/gopher-lua"
)

type RequestSuite struct {
	suite.Suite

	vm *lua.LState
}

func TestRequests(t *testing.T) {
	suite.Run(t, &RequestSuite{})
}

func (t *RequestSuite) SetupTest() {
	t.vm = lua.NewState()
	Preload(t.vm)
}

func (t *RequestSuite) TestGET() {
	const src = `
local http = require('waffle-http-client')
local r = http.new_request('GET', 'https://httpbin.org/get?xyz')

local x = http.do_request(r)
print(x.body)
`
	defer t.vm.Close()

	t.NoError(t.vm.DoString(src))
}

func (t *RequestSuite) TestPOST() {
	const src = `
local http = require('waffle-http-client')
local r = http.new_request('POST', 'https://httpbin.org/post?xyz', '{}')

local x = http.do_request(r)
print(x.body)
`
	defer t.vm.Close()

	t.NoError(t.vm.DoString(src))
}

func (t *RequestSuite) TestInvalidRequestType() {
	const src = `
local http = require('waffle-http-client')
local r = {}

local x = http.do_request(r)
print(x.body)
`
	defer t.vm.Close()

	err := t.vm.DoString(src)
	t.Error(err)
	t.T().Log(err.Error())
}
