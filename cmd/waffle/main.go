package main

import (
	"fmt"
	"os"

	"bits.chrsm.org/waffle"
	"bits.chrsm.org/waffle/internal/luahttp"
	"bits.chrsm.org/waffle/internal/luajson"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	//log.Println(waffle.Version, waffle.Date)
	_ = waffle.Version

	vm := lua.NewState(lua.Options{
		IncludeGoStackTrace: true,
	})

	luahttp.Preload(vm)
	luajson.Preload(vm)

	if err := vm.DoFile(os.Args[1]); err != nil {
		fmt.Println(err)
	}
}
