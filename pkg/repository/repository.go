package repository

type Domains struct {

}

type Prices struct {

}

type Repository struct {
	Domains
	Prices
}

func NewRepository() *Repository {
	return &Repository{}
}
