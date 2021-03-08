package modelsstruct

// Prices ...
type Prices struct {
	Prices []Price `json:"prices"`
}

// Price ...
type Price struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	GEO      string `json:"geo" db:"geo"`
	Country  string `json:"country" db:"country"`
	Country1 string `json:"country1" db:"country_1"`
	Country2 string `json:"country2" db:"country_2"`
	Old      string `json:"old" db:"old"`
	New      string `json:"new" db:"new"`
	Money    string `json:"money" db:"money"`
	Tel      string `json:"tel" db:"tel"`
}
