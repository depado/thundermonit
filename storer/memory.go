package storer

import (
	"github.com/Depado/thundermonit/domain"
	"github.com/Depado/thundermonit/uc"
)

type memoryStorer struct {
	s []*domain.Service
	m map[int64]*domain.Service
}

func NewMemoryStorer(ss ...*domain.Service) uc.Storer {
	m := memoryStorer{
		m: make(map[int64]*domain.Service),
	}
	for _, s := range ss {
		m.s = append(m.s, s)
		m.m[s.ID] = s
	}
	return m
}

func (m memoryStorer) GetAllServices() ([]*domain.Service, error) {
	return m.s, nil
}

func (m memoryStorer) GetService(id int64) (*domain.Service, error) {
	if s, ok := m.m[id]; ok {
		return s, nil
	}
	return nil, nil
}

func (m memoryStorer) GetRepo(s *domain.Service) (*domain.Repo, error) {
	return s.Repo, nil
}

func (m memoryStorer) GetAllRepos() ([]*domain.Repo, error) {
	var out []*domain.Repo
	for _, s := range m.s {
		if s.Repo != nil {
			out = append(out, s.Repo)
		}
	}
	return out, nil
}
