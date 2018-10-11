package graphql

import (
	"github.com/samsarahq/thunder/graphql/schemabuilder"

	"github.com/Depado/thundermonit/domain"
)

// RegisterCI registers the CI struct to the Schema
func (rh RequestHandler) RegisterCI(s *schemabuilder.Schema) {
	s.Object("ci", domain.CI{})
}
