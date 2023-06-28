package main

import (
	"fmt"
	"github.com/dimasbayuseno/farmacare-test/database"
	"github.com/dimasbayuseno/farmacare-test/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	db, err := database.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", handlers.Login(db))

	registerRoutes(e, db)

	addr := ":3000"
	fmt.Printf("Server listening on %s\n", addr)
	e.Start(addr)
}
