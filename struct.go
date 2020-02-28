package assignment1

//Country struct
type Country struct {
	Alpha2Code string   `json:"alpha2Code"`
	Name       string   `json:"name"`
	Flag       string   `json:"flag"`
	SpeciesKey []int    `json:"speciesKey"`
	Species    []string `json:"species"`
}

//KeyAndName struct
type KeyAndName struct {
	Key    int    `json:"speciesKey"`
	Specie string `json:"species"`
}

//Results struct
type Results struct {
	Result []KeyAndName `json:"results"`
}

//Species struct
type Species struct {
	Key            int    `json:"key"`
	Kingdom        string `json:"kingdom"`
	Phylum         string `json:"phulum"`
	Order          string `json:"order"`
	Family         string `json:"family"`
	Genus          string `json:"genus"`
	ScientificName string `json:"scientificName"`
	CanonicalName  string `json:"canonicalName"`
	Year           string `json:"year"`
}

//SpeciesYear struct
type SpeciesYear struct {
	Year string `json:"bracketYear"`
}

//Diag struct
type Diag struct {
	GBIF        int
	RestCountry int
	Version     string
	Uptime      float64
}
