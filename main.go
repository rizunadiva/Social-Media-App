package main

import (
	"fmt"
	"socialmedia-app/config"
	"socialmedia-app/factory"
	"socialmedia-app/infrastruktur/database/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	e := echo.New()

	factory.Initfactory(e, db)

	fmt.Println("Running program ....")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
