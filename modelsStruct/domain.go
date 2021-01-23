package modelsStruct

type Domain struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"domain_name"`
	Flag    int    `json:"flag" db:"checked"`
	Created string `json:"created_at" db:"created_at"`
}

type Domains struct {
	Domains []Domain
}

type Flag struct {
	SetFlag int `json:"set_flag" binding:"required"`
}