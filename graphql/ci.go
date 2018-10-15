package graphql

import (
	"github.com/samsarahq/thunder/graphql/schemabuilder"

	"github.com/Depado/thundermonit/domain"
)

const (
	ciDesc = "A type representing a CI object associated to a service"
	ciName = "ci"
)

// RegisterCI registers the CI struct to the Schema
func (rh RequestHandler) RegisterCI(s *schemabuilder.Schema) {
	obj := s.Object(ciName, domain.CI{})
	obj.Description = ciDesc
}
