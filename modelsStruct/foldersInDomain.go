package modelsstruct

// Folder ...
type Folder struct {
	ID           int    `json:"id" db:"id"`
	Folder       string `json:"folder" db:"folder"`
	Flag         int    `json:"flag" db:"flag"`
	DomainID     int    `json:"domain_id" db:"domain_id"`
	OfferNamesID int    `json:"offer_names_id" db:"offer_names_id"`
}

// Folders ...
type Folders struct {
	Folders []Folder `json:"folders"`
}
