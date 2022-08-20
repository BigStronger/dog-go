package app

import (
	"os"
	"os/signal"
	"syscall"
)

func Blocking() os.Signal {
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	select {
	case sign := <-signChan:
		return sign
	}
}
