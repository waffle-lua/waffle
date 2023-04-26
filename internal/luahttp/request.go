package luahttp

import (
	"bytes"
	"io"
	"net/http"

	"bits.chrsm.org/waffle"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

const typeUserdataRequest = "waffle_http_request_ud"

var _ = luar.New

type Request struct {
	*http.Request
}

func NewRequest(ls *lua.LState) int {
	const defUA = `waffle-http-client`
	_ = waffle.Version

	/**
	 * 1 = method
	 * 2 = url
	 * 3 = body|nil
	 */

	method := ls.CheckString(1)
	url := ls.CheckString(2)

	body := new(bytes.Buffer)
	if ls.GetTop() > 2 {
		body.WriteString(ls.CheckString(3))
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		ls.Push(lua.LNil)
		ls.Push(lua.LString(err.Error()))

		return 2
	}

	r := &Request{req}
	ud := ls.NewUserData()
	ud.Value = r
	ls.SetMetatable(ud, ls.GetTypeMetatable(typeUserdataRequest))
	ls.Push(ud)

	return 1
}

func DoRequest(ls *lua.LState) int {
	// look for userdata
	var req *Request

	ud := ls.CheckUserData(1)
	if v, ok := ud.Value.(*Request); ok {
		req = v
	}

	resp, err := http.DefaultClient.Do(req.Request)
	if err != nil {
		ls.Push(lua.LNil)
		ls.Push(lua.LString(err.Error()))

		return 2
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)

	headers := ls.NewTable()
	for k := range resp.Header {
		headers.RawSetString(k, lua.LString(resp.Header[k][0]))
	}

	res := ls.NewTable()
	ls.SetField(res, "status", lua.LNumber(resp.StatusCode))
	ls.SetField(res, "body", lua.LString(buf.String()))
	ls.SetField(res, "headers", headers)
	ls.Push(res)

	return 1
}
