package handlers

import (
	"github.com/dimasbayuseno/farmacare-test/utils"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func GetPokemonList(c echo.Context) error {
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > 100 {
		limit = 10
	}
	pokemonList, err := utils.FetchPokemonList(page, limit)
	if err != nil {
		log.Println("Failed to fetch Pokémon list:", err)
		return c.String(http.StatusInternalServerError, "Failed to fetch Pokémon list")
	}

	return c.JSON(http.StatusOK, pokemonList.Results)
}
