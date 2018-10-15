package graphql

import (
	"github.com/samsarahq/thunder/graphql/schemabuilder"

	"github.com/Depado/thundermonit/domain"
)

const (
	repoName = "repo"
	repoDesc = "A type that describes a repository associated to a service"
)

// RegisterRepo registers the Repo struct to the Schema
func (rh RequestHandler) RegisterRepo(s *schemabuilder.Schema) {
	// Object registration
	obj := s.Object(repoName, domain.Repo{})
	obj.Description = repoDesc

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
