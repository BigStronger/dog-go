package web

import (
	"github.com/kataras/iris/v12"
	"github.com/lucas-clemente/quic-go/http3"
	"log"
)

func New(callback Callback, addr string, mode Mode, keypair ...string) {
	server := iris.New()
	callback(server)
	switch mode {
	case ModeNormal:
		if err := server.Listen(addr); err != nil {
			log.Println(err)
		}
	case ModeWeb3:
		if err := http3.ListenAndServeQUIC(addr, keypair[0], keypair[1], server); err != nil {
			log.Println(err)
		}
	default:
		log.Fatalln("unknown mode")
	}

}
