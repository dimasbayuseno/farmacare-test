package main

import (
	"github.com/dimasbayuseno/farmacare-test/handlers"
	"github.com/dimasbayuseno/farmacare-test/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func registerRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/battle-information", handlers.GetAllBattleInfoHandler(db), middleware.AuthMiddleware(db, "boss"))
	e.GET("/pokemon", handlers.GetPokemonList, middleware.AuthMiddleware(db, "operational"))
	e.GET("/score-list", handlers.GetScoreListHandler(db), middleware.AuthMiddleware(db, "boss"))
	e.POST("/battles", handlers.CreateBattle(db), middleware.AuthMiddleware(db, "merchant"))
	e.PUT("/battles", handlers.UpdateScoresHandler(db), middleware.AuthMiddleware(db, "boss"))
}
