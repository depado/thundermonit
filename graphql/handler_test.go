package graphql

import (
	"testing"

	"github.com/Depado/thundermonit/storer"
	"github.com/Depado/thundermonit/uc"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

func TestRequestHandler_RegisterAll(t *testing.T) {
	var err error
	var s *graphql.Schema
	var out []byte

	sb := schemabuilder.NewSchema()
	rh := RequestHandler{
		Interactor: uc.NewInteractor(storer.NewNoopStorer()),
	}
	rh.RegisterAll(sb)

	if s, err = sb.Build(); err != nil {
		t.Errorf("Registering CI generated an error: %s", err)
	}
	introspection.AddIntrospectionToSchema(s)
	if out, err = introspection.ComputeSchemaJSON(*sb); err != nil {
		t.Errorf("Computing schema JSON generated an error: %s", err)
	}
	if !TypeInSchema(out, repoName, repoDesc) {
		t.Errorf("Couldn't find 'service' type with all matching values in output schema")
	}
	if !TypeInSchema(out, ciName, ciDesc) {
		t.Errorf("Couldn't find 'service' type with all matching values in output schema")
	}
	if !TypeInSchema(out, serviceName, serviceDesc) {
		t.Errorf("Couldn't find 'service' type with all matching values in output schema")
	}
}
