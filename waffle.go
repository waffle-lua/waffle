package waffle

import (
	lua "github.com/yuin/gopher-lua"

	// lua modules
	lcrypto "github.com/tengattack/gluacrypto"
	lcmd "github.com/vadv/gopher-lua-libs/cmd"
	ldb "github.com/vadv/gopher-lua-libs/db"
	lfilepath "github.com/vadv/gopher-lua-libs/filepath"
	lhttp "github.com/vadv/gopher-lua-libs/http"
	linspect "github.com/vadv/gopher-lua-libs/inspect"
	lioutil "github.com/vadv/gopher-lua-libs/ioutil"
	ljson "github.com/vadv/gopher-lua-libs/json"
	lplugin "github.com/vadv/gopher-lua-libs/plugin"
	lregexp "github.com/vadv/gopher-lua-libs/regexp"
	lstrings "github.com/vadv/gopher-lua-libs/strings"
	ltcp "github.com/vadv/gopher-lua-libs/tcp"
	ltime "github.com/vadv/gopher-lua-libs/time"
)

var preloads = []func(*lua.LState){
	lhttp.Preload,
	ljson.Preload,
	lstrings.Preload,
	ldb.Preload,
	ltcp.Preload,
	linspect.Preload,
	ltime.Preload,
	lfilepath.Preload,
	lcmd.Preload,
	lioutil.Preload,
	lregexp.Preload,
	lplugin.Preload,
	lcrypto.Preload,
}

func PreloadAll(ls *lua.LState) {
	Preload(ls)

	for i := range preloads {
		preloads[i](ls)
	}
}

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
