package main

import (
	"fmt"
	"os"

	"bits.chrsm.org/waffle"

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
	waffle.Preload,
}

func main() {
	vm := lua.NewState(lua.Options{
		IncludeGoStackTrace: true,
	})

	for i := range preloads {
		preloads[i](vm)
	}

	if err := vm.DoFile(os.Args[1]); err != nil {
		fmt.Println(err)
	}
}
