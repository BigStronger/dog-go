package web

import (
	"net/http"
	_ "net/http/pprof"
)

func PProf(addr string) error {
	return http.ListenAndServe(addr, nil)
}
