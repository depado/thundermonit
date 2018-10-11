package uc

import "github.com/Depado/thundermonit/domain"

func (i Interactor) GetAllServices() ([]*domain.Service, error) {
	return i.storer.GetAllServices()
}

func (i Interactor) GetService(id int64) (*domain.Service, error) {
	return i.storer.GetService(id)
}
