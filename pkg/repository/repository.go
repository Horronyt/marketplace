package repository

type Authorization interface {
}

type Listing interface {
}

type Repository struct {
	Authorization
	Listing
}

func NewRepository() *Repository {
	return &Repository{}
}
