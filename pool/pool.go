package pool

import (
	lua "github.com/yuin/gopher-lua"
	"sync"
)

type Pool struct {
	lock  *sync.Mutex
	newVM func() *lua.LState
	vms   []*lua.LState
}

func (p *Pool) Get() *lua.LState {
	p.lock.Lock()
	defer p.lock.Unlock()

	var (
		n  = len(p.vms)
		vm *lua.LState
	)

	if n == 0 {
		return p.New()
	}

	vm = p.vms[n-1]
	p.vms = p.vms[0 : n-1]

	return vm
}

func (p *Pool) Put(L *lua.LState) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.vms = append(p.vms, L)
}

func (p *Pool) New() *lua.LState {
	return p.newVM()
}

func (p *Pool) Close() {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, vm := range p.vms {
		vm.Close()
	}
}

func New(newVM func() *lua.LState) *Pool {
	if newVM == nil {
		newVM = func() *lua.LState { return lua.NewState() }
	}

	return &Pool{
		lock:  &sync.Mutex{},
		newVM: newVM,
		vms:   make([]*lua.LState, 0, 8),
	}
}
