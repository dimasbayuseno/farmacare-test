package models

type PokemonList struct {
	Results []Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
