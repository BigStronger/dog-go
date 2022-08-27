package web

import (
	"github.com/kataras/iris/v12"
)

type Callback func(server *iris.Application)

type Mode int

const (
	ModeNormal Mode = iota + 1
	ModeWeb3
)
