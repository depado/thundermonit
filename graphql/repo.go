package graphql

import (
	"github.com/samsarahq/thunder/graphql/schemabuilder"

	"github.com/Depado/thundermonit/domain"
)

// RegisterRepo registers the Repo struct to the Schema
func (rh RequestHandler) RegisterRepo(s *schemabuilder.Schema) {
	// Object registration
	obj := s.Object("repo", domain.Repo{})
	obj.Description = "A type that describes a repository associated to a service"

	// Queries registration
	q := s.Query()
	q.FieldFunc("repos", func() ([]*domain.Repo, error) {
		return rh.Interactor.GetAllRepos()
	})

	// Mutations registration
	m := s.Mutation()
	m.FieldFunc("addRepo", func(args struct {
		URL  string
		Path string
		Type string
		Host string
	}) domain.Repo {
		newRepo := domain.Repo{URL: args.URL, Path: args.Path, Type: args.Type, Host: args.Host}
		// Store
		return newRepo
	})
}
