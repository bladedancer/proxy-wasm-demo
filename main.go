package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type helloHTTPContext struct {
	proxywasm.DefaultHttpContext
}

func main() {

	proxywasm.SetNewHttpContext(newHttpContext)
}

func newHttpContext(uint32, uint32) proxywasm.HttpContext {

	return &helloHTTPContext{}
}

func (ctx *helloHTTPContext) OnHttpRequestHeaders(numHeaders int, _ bool) types.Action {

	if numHeaders > 0 {
		headers, err := proxywasm.GetHttpRequestHeaders()
		if err != nil {
			proxywasm.LogErrorf("failed to get request headers with '%v'", err)
			return types.ActionContinue
		}
		proxywasm.LogInfof("request headers: '%+v'", headers)
	}

	return types.ActionContinue
}
