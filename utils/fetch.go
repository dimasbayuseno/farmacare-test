package utils

import (
	"encoding/json"
	"fmt"
	"github.com/dimasbayuseno/farmacare-test/models"
	"net/http"
)

func FetchPokemonList(page, limit int) (*models.PokemonList, error) {
	offset := (page - 1) * limit
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?offset=%d&limit=%d", offset, limit)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pokemonList models.PokemonList
	err = json.NewDecoder(resp.Body).Decode(&pokemonList)
	if err != nil {
		return nil, err
	}

	return &pokemonList, nil
}
