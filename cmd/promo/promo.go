package main

import (
	"flag"

	"github.com/AlexSugak/skycoin-promo/src/promo"
	"github.com/AlexSugak/skycoin-promo/src/util/logger"
)

func main() {
	bindingFlag := flag.String("binding", "0.0.0.0:8081", "HTTP server binding")
	// TODO: specify default value for recaptchaSecret
	recaptchaSecret := flag.String("recaptchaSecret", "", "recaptcha secret")

	flag.Parse()

	log := logger.InitLogger()

	server := promo.NewHTTPServer(*bindingFlag, *recaptchaSecret, log)
	if err := server.Run(); err != nil {
		panic(err.Error())
	}
}
