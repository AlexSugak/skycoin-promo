package main

import (
	"fmt"

	"github.com/AlexSugak/skycoin-promo/src/promo_codes_generator"

	_ "github.com/go-sql-driver/mysql"
	"github.com/namsral/flag"

	"github.com/AlexSugak/skycoin-promo/src/promo"
	activator "github.com/AlexSugak/skycoin-promo/src/promo_activator"
	"github.com/AlexSugak/skycoin-promo/src/skynode"
	"github.com/AlexSugak/skycoin-promo/src/util/logger"
	"github.com/jmoiron/sqlx"
)

func main() {
	flag.String(flag.DefaultConfigFlagname, "", "path to config file")
	bindingFlag := flag.String("binding", "0.0.0.0:8081", "HTTP server binding")
	recaptchaSecret := flag.String("recaptchaSecret", "6LdTKksUAAAAAAMgKNhOcxgWYYCDRrgx8YoEH5qX", "Recaptcha secret")
	skyNodeURL := flag.String("skyNodeURL", "http://127.0.0.1:6420", "A base URL of skynode")
	mysqlFlag := flag.String("mysql", "root:root@(0.0.0.0:3306)", "MySQL connect string")

	flag.Parse()

	sqlDb, err := initDb(*mysqlFlag)
	if err != nil {
		panic(err.Error())
	}

	log := logger.InitLogger()
	generator := generator.NewGenerator(sqlDb)
	activator := activator.NewActivator(sqlDb)
	skyNode := skynode.NewSkyNode(*skyNodeURL)
	server := promo.NewHTTPServer(*bindingFlag, *recaptchaSecret, log, activator, skyNode, *generator)

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
