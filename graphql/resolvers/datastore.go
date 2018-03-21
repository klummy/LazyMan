package lmresolvers

import (
	"github.com/graphql-go/graphql"

	lmdatastore "github.com/klummy/LazyMan-Go/databases/datastore"
)

// DatastoreReadResolver - Resolver for Google Datastore
func DatastoreReadResolver(args ...interface{}) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		entityName := args[0].(string)
		entityID := params.Args["ID"].(string)
		entityStruct := args[1]

		entity, entityReadErr := lmdatastore.DatastoreRead(params.Context, entityName, entityID, entityStruct)
		if entityReadErr != nil {
			return nil, entityReadErr
		}

		return entity, nil
	}
}
