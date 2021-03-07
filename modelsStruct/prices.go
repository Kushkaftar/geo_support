package modelsStruct

// Prices ...
type Prices struct {
	Prices []Price `json:"prices"`
}

// Price ...
type Price struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	GEO      string `json:"geo"`
	Country  string `json:"country"`
	Country1 string `json:"country1"`
	Country2 string `json:"country2"`
	Old      string `json:"old"`
	New      string `json:"new"`
	Money    string `json:"money"`
	Tel      string `json:"tel"`
}
