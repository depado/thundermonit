package uc

import (
	"github.com/Depado/thundermonit/domain"
)

// Interactor is the interactor
type Interactor struct {
	storer Storer
}

// NewInteractor returns a new Interactor
func NewInteractor(s Storer) Interactor {
	return Interactor{
		storer: s,
	}
}

// Storer is an interface that defines how to access stored data
type Storer interface {
	GetAllServices() ([]*domain.Service, error)
	GetService(id int64) (*domain.Service, error)
	GetRepo(*domain.Service) (*domain.Repo, error)
	GetAllRepos() ([]*domain.Repo, error)
}
