package uc

import "github.com/Depado/thundermonit/domain"

func (i Interactor) GetAllRepos() ([]*domain.Repo, error) {
	return i.storer.GetAllRepos()
}

func (i Interactor) GetRepo(s *domain.Service) (*domain.Repo, error) {
	return i.storer.GetRepo(s)
}
