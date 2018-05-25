package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/namsral/flag"

	"github.com/AlexSugak/skycoin-promo/src/db"
	"github.com/AlexSugak/skycoin-promo/src/promo"
	"github.com/AlexSugak/skycoin-promo/src/skynode"
	"github.com/AlexSugak/skycoin-promo/src/util/logger"
	"github.com/jmoiron/sqlx"
)

func main() {
	flag.String(flag.DefaultConfigFlagname, "", "path to config file")
	bindingFlag := flag.String("binding", "0.0.0.0:8081", "HTTP server binding")
	recaptchaSecret := flag.String("recaptchaSecret", "6LcIDlkUAAAAAB7-YebjJSUBb2aINasOnNk0zF8W", "Recaptcha secret")
	skyNodeURL := flag.String("skyNodeURL", "http://127.0.0.1:6420", "A base URL of skynode")
	mysqlFlag := flag.String("mysql", "root:root@(0.0.0.0:3306)", "MySQL connect string")

	flag.Parse()

	sqlDb, err := initDb(*mysqlFlag)
	if err != nil {
		panic(err.Error())
	}

	storage := db.NewStorage(sqlDb)
	log := logger.InitLogger()
	activator := storage
	skyNode := skynode.NewSkyNode(*skyNodeURL)
	server := promo.NewHTTPServer(*bindingFlag, *recaptchaSecret, log, activator, skyNode)

	if err := server.Run(); err != nil {
		panic(err.Error())
	}
}

func initDb(addr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s/skycoinpromo?parseTime=true", addr))
	if err != nil {
		return nil, err
	}

	return db, nil
}
