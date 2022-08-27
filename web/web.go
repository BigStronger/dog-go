package web

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-clemente/quic-go/http3"
	"log"
)

func New(callback Callback, addr string, mode Mode, keypair ...string) {
	server := gin.Default()
	callback(server)
	switch mode {
	case ModeNormal:
		if err := server.Run(addr); err != nil {
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

func ReleaseMode() {
	gin.SetMode(gin.ReleaseMode)
}
