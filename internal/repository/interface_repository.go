package repository

import (
	. "github.com/v8tix/bender/internal/models"
)

type Repository interface {
	GetBeerRepository() BeerRepository
}

type BeerRepository interface {
	FindBeerById(id string) (*Beer, error)
	InsertBeer(beer *Beer) (string, error)
	UpdateBeer(beer *Beer) (string, error)
}
