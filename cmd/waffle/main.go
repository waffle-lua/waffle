package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/waffle-lua/waffle"

	lua "github.com/yuin/gopher-lua"
)

var (
	src   string
	stdin bool
)

func main() {
	flag.StringVar(&src, "src", "", "lua script to execute; if ending in .yue, will run yue first")
	flag.BoolVar(&stdin, "stdin", false, "whether to look at stdin for source code")
	flag.Parse()

	vm := lua.NewState(lua.Options{
		IncludeGoStackTrace: true,
	})

	waffle.PreloadAll(vm)

	switch {
	case src != "":
		if strings.Contains(src, ".yue") {
			// first, we need to generate lua code from this
			cmd := exec.Command("yue", "--")
			buf := new(bytes.Buffer)

			scriptContent, err := ioutil.ReadFile(src)
			if err != nil {
				fmt.Println("error converting yue:", err)
				os.Exit(1)
			}

			cmd.Stdin = strings.NewReader(string(scriptContent))
			cmd.Stdout = buf

			cmd.Start()
			cmd.Wait()

			if err := vm.DoString(string(buf.Bytes())); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			if err := vm.DoFile(src); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case stdin:
		in, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
		}

		if err := vm.DoString(string(in)); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("error: no source provided. use `-src` or `-stdin`")
		os.Exit(1)
	}
}
