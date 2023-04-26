package luajson

import (
	lua "github.com/yuin/gopher-lua"
)

func Preload(ls *lua.LState) {
	ls.PreloadModule("waffle-json", loader)
}

func loader(ls *lua.LState) int {
	if err := ls.DoString(luajson); err != nil {
		ls.RaiseError("waffle-json error: %s", err.Error())
	}

	return 1
}
