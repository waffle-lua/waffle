package waffle

import (
	lua "github.com/yuin/gopher-lua"
)

func Preload(ls *lua.LState) {
	ls.PreloadModule("waffle", func(l *lua.LState) int {
		ls.Push(ls.SetFuncs(ls.NewTable(), exports))

		return 1
	})
}

var exports = map[string]lua.LGFunction{
	"version": func(ls *lua.LState) int {
		ls.Push(lua.LString(Version))

		return 1
	},
	"date": func(ls *lua.LState) int {
		ls.Push(lua.LString(Date))

		return 1
	},
}
