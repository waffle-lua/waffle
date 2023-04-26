package luajson

import (
	"testing"

	"github.com/stretchr/testify/suite"
	lua "github.com/yuin/gopher-lua"
)

// we don't really need to test everything here, as rxi already has tests
type JsonSuite struct {
	suite.Suite

	vm *lua.LState
}

func TestJson(t *testing.T) {
	suite.Run(t, &JsonSuite{})
}

func (t *JsonSuite) SetupTest() {
	t.vm = lua.NewState()
	Preload(t.vm)
}

func (t *JsonSuite) TearDownTest() {
	t.vm.Close()
}

func (t *JsonSuite) TestDecode() {
	const src = `
local json = require('waffle-json')

local x = "[1,2,3,{\"x\":10}]"

local dec = json.decode(x)

if #dec ~= 4 then
	error("dec length expected to be 4")
end


if dec[1] ~= 1 or dec[2] ~= 2 or dec[3] ~= 3 then
	error("dec[1,2,3] expected to be 1, 2 and 3 respectively")
end

if type(dec[4]) ~= "table" or dec[4].x ~= 10 then
	error("dec[4] expected to be table{x=10}")
end
`

	t.NoError(t.vm.DoString(src))
}
