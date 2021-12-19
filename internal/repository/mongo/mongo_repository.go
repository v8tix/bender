package repository

import (
	. "github.com/v8tix/bender/internal/repository"
	. "github.com/v8tix/kit/app"
)

type RepositoryI struct {
	app            App
	beerRepository BeerRepository
}

func NewRepositoryI(
	app App,
) *RepositoryI {
	repository := &RepositoryI{
		app:            app,
		beerRepository: NewBeerRepository(app),
	}

	return repository
}

func (m *RepositoryI) GetBeerRepository() BeerRepository {
	return m.beerRepository
}
