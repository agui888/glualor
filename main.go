package main

import (
	"github.com/yuin/gopher-lua"
	"github.com/zhu327/gluatemplate"
	"github.com/zhu327/gluaweb"
	luajson "layeh.com/gopher-json"
	"net/http"
	"sync"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	L := luaPool.Get()
	defer func() {
		luaPool.Put(L)
	}()
	ctx := gluaweb.NewWebContext(w, r).WebContext(L)
	L.SetGlobal("gluaweb", ctx) // load http
	app := L.GetGlobal("app")
	if err := L.CallByParam(lua.P{
		Fn:      L.GetField(app, "run"),
		NRet:    1,
		Protect: true,
	}, app); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":8080", nil)
}

type lStatePool struct {
	m     sync.Mutex
	saved []*lua.LState
}

func (pl *lStatePool) Get() *lua.LState {
	pl.m.Lock()
	defer pl.m.Unlock()
	n := len(pl.saved)
	if n == 0 {
		return pl.New()
	}
	x := pl.saved[n-1]
	pl.saved = pl.saved[0 : n-1]
	return x
}

func (pl *lStatePool) New() *lua.LState {
	L := lua.NewState()
	luajson.Preload(L)
	L.PreloadModule("template", gluatemplate.Loader)
	if err := L.DoFile("main.lua"); err != nil {
		panic(err)
	}
	return L
}

func (pl *lStatePool) Put(L *lua.LState) {
	pl.m.Lock()
	defer pl.m.Unlock()
	pl.saved = append(pl.saved, L)
}

func (pl *lStatePool) Shutdown() {
	for _, L := range pl.saved {
		L.Close()
	}
}

// Global LState pool
var luaPool = &lStatePool{
	saved: make([]*lua.LState, 0, 4),
}
