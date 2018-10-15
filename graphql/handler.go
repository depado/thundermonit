package graphql

import (
	"github.com/samsarahq/thunder/graphql/schemabuilder"

	"github.com/Depado/thundermonit/uc"
)

// RequestHandler wraps an interactor
type RequestHandler struct {
	Interactor uc.Interactor
}

// NewRequestHandler will return a new RequestHandler with the given interactor
// associated
func NewRequestHandler(i uc.Interactor) RequestHandler {
	return RequestHandler{i}
}

// RegisterAll will register all the required structs to the schemabuilder
func (rh RequestHandler) RegisterAll(s *schemabuilder.Schema) {
	rh.RegisterService(s)
	rh.RegisterRepo(s)
	rh.RegisterCI(s)
}
