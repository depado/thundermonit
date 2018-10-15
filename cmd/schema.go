package cmd

import (
	"io/ioutil"

	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	graphqlimpl "github.com/Depado/thundermonit/graphql"
	"github.com/Depado/thundermonit/storer"
	"github.com/Depado/thundermonit/uc"
)

// SchemaCmd will write out the generated JSON schema
var SchemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Generate and write the schema.json file on disk",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var output []byte

		ofile := viper.GetString("output")
		builder := schemabuilder.NewSchema()
		rh := graphqlimpl.NewRequestHandler(uc.NewInteractor(storer.NewNoopStorer()))
		rh.RegisterAll(builder)
		introspection.AddIntrospectionToSchema(builder.MustBuild())
		if output, err = introspection.ComputeSchemaJSON(*builder); err != nil {
			logrus.WithError(err).Fatal("Couldn't compute the schema to JSON")
		}
		if err = ioutil.WriteFile(ofile, output, 0644); err != nil {
			logrus.WithError(err).Fatal("Couldn't write the schema on disk")
		}
		logrus.WithField("file", ofile).Info("Computed and built the schema")
	},
}

// AddSchemaFlags will add schema flags to the schema command
func AddSchemaFlags(c *cobra.Command) {
	c.Flags().StringP("output", "o", "output.json", "json file output")
	if err := viper.BindPFlags(c.Flags()); err != nil {
		logrus.WithError(err).WithField("step", "AddSchemaFlags").Fatal("Couldn't bind flags")
	}
}
