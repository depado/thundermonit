package graphql

import "github.com/Depado/thundermonit/uc"

// RequestHandler wraps an interactor
type RequestHandler struct {
	Interactor uc.Interactor
}
