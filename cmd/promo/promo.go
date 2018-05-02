package main

import (
	"flag"

	"github.com/AlexSugak/skycoin-promo/src/promo"
	"github.com/AlexSugak/skycoin-promo/src/util/logger"
)

func main() {
	bindingFlag := flag.String("binding", "0.0.0.0:8081", "HTTP server binding")

	flag.Parse()

	log := logger.InitLogger()

	server := promo.NewHTTPServer(*bindingFlag, log)
	if err := server.Run(); err != nil {
		panic(err.Error())
	}
}
