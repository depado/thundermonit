package storer

import (
	"github.com/Depado/thundermonit/domain"
	"github.com/Depado/thundermonit/uc"
)

type noopStorer struct{}

// NewNoopStorer returns a new Noop storer which basically does nothing
// Useful for creating an empty requestHandler
func NewNoopStorer() uc.Storer {
	return noopStorer{}
}

func (noopStorer) GetAllServices() ([]*domain.Service, error) {
	return nil, nil
}

func (noopStorer) GetService(id int64) (*domain.Service, error) {
	return nil, nil
}

func (noopStorer) GetRepo(*domain.Service) (*domain.Repo, error) {
	return nil, nil
}

func (noopStorer) GetAllRepos() ([]*domain.Repo, error) {
	return nil, nil
}
