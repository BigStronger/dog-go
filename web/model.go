package web

import "github.com/gin-gonic/gin"

type Callback func(server *gin.Engine)

type Mode int

const (
	ModeNormal Mode = iota + 1
	ModeWeb3
)
