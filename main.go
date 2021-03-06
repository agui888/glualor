package main

import (
	"github.com/yuin/gopher-lua"
	"github.com/zhu327/gluatemplate"
	"github.com/zhu327/gluaweb"
	luajson "layeh.com/gopher-json"
	"net/http"
	"pool"
)

var luaPool = pool.New(func() *lua.LState {
	L := lua.NewState()
	luajson.Preload(L)
	L.PreloadModule("template", gluatemplate.Loader)
	if err := L.DoFile("main.lua"); err != nil { // 加载 lua 代码
		panic(err)
	}
	return L
})

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	L := luaPool.Get()
	defer luaPool.Put(L)
	ctx := gluaweb.NewWebContext(w, r).WebContext(L)
	L.SetGlobal("gluaweb", ctx) // 传入 web 上下文
	app := L.GetGlobal("app")
	if err := L.CallByParam(lua.P{
		Fn:      L.GetField(app, "run"),
		NRet:    0,
		Protect: true,
	}, app); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	fsh := http.FileServer(http.Dir("./app/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsh))
	http.ListenAndServe(":8080", nil)
}
