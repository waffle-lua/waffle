package waffle

import (
	lua "github.com/yuin/gopher-lua"
)

type Waffle struct {
}

func New() *Waffle {
	return &Waffle{}
}

func (w *Waffle) init() {
	lua.NewState()
}
