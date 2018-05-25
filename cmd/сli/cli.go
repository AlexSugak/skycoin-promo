package main

import (
	"fmt"
	"os"

	"github.com/AlexSugak/skycoin-promo/src/cli"
)

func main() {
	app := cli.NewApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
