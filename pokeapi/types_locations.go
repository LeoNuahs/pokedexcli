package pokeapi

type RespShallowLocation struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"results"`
	} `json:"results"`
	// Results []LocationArea `json:"results"`
}

// type LocationArea struct {
// 	Name string `json:"name"`
// 	Url  string `json:"url"`
// }
