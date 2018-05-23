package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/namsral/flag"

	"github.com/AlexSugak/skycoin-promo/db"
	"github.com/AlexSugak/skycoin-promo/src/promo"
	"github.com/AlexSugak/skycoin-promo/src/util/logger"
	"github.com/jmoiron/sqlx"
)

func main() {
	flag.String(flag.DefaultConfigFlagname, "", "path to config file")
	bindingFlag := flag.String("binding", "0.0.0.0:8081", "HTTP server binding")
	// TODO: specify default value for recaptchaSecret
	recaptchaSecret := flag.String("recaptchaSecret", "", "Recaptcha secret")
	mysqlFlag := flag.String("mysql", "root:root@(0.0.0.0:3306)", "MySQL connect string")

	flag.Parse()

	sqlDb, err := initDb(*mysqlFlag)
	if err != nil {
		panic(err.Error())
	}

	storage := db.NewStorage(sqlDb)
	log := logger.InitLogger()
	activator := storage
	server := promo.NewHTTPServer(*bindingFlag, *recaptchaSecret, log, activator)

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
