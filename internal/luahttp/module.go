package luahttp

import (
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func Preload(ls *lua.LState) {
	ls.PreloadModule("waffle-http-client", loader)

	_ = luar.New
}

func loader(ls *lua.LState) int {
	ls.Push(ls.SetFuncs(ls.NewTable(), exports))

	waffleUD := ls.NewTypeMetatable(typeUserdataRequest)
	ls.SetGlobal(typeUserdataRequest, waffleUD)
	ls.SetField(waffleUD, "__index", ls.SetFuncs(
		ls.NewTable(),
		map[string]lua.LGFunction{},
	))

	return 1
}

var exports = map[string]lua.LGFunction{
	"new_request": NewRequest,
	"do_request":  DoRequest,
}
