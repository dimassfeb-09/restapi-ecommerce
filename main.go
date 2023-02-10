package main

import (
	"github.com/dimassfeb-09/restapi-ecommerce.git/app"
	"github.com/dimassfeb-09/restapi-ecommerce.git/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	db, err := app.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	r.HandleMethodNotAllowed = true
	r.NoRoute(middleware.NoRoute)
	r.Use(middleware.AllowAccessMiddleware)

	initializer := app.NewInitialized(db)
	app.RoutesApi(initializer, r)

	r.Run(":3000")
}
