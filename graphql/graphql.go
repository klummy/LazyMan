package lmgraphql

import (
	"github.com/graphql-go/graphql"
	"github.com/klummy/Grest"

	lmresolvers "github.com/klummy/LazyMan-Go/graphql/resolvers"
)

// 1. Create a resolver client that contains the resolver function
// Client is called with
// 		- Resolvers to use for graphql actions (queries/mutations)

// ResolveType - Type for Resolve
type ResolveType func(resolverPath string, args ...interface{}) graphql.FieldResolveFn

// Resolver - Type definition for a resolvers
type Resolver map[string]func(args ...interface{}) graphql.FieldResolveFn

// ResolverClient - The type for the resolver client
type ResolverClient struct {
	Resolve ResolveType
}

// DefaultResolvers are the default resolvers
func DefaultResolvers() Resolver {
	return Resolver{
		"read": lmresolvers.DefaultReadResolver,
	}
}

var resolvers = DefaultResolvers()

// NewResolverClient creates a new Graphql Resolver from the provided configuration.
func NewResolverClient(customerResolvers Resolver) ResolverClient {
	if customerResolvers != nil {
		resolvers = customerResolvers
	}

	return ResolverClient{
		Resolve: Resolve,
	}

	// for val, key := range resolvers {
	// 	log.Println(val)
	// 	log.Println(key)
	// }
}

// Resolve using a selected config
func Resolve(resolverPath string, args ...interface{}) graphql.FieldResolveFn {
	resolve := resolvers[resolverPath]

	if resolve == nil {
		return func(params graphql.ResolveParams) (interface{}, error) {
			return nil, grest.CustomError("Resolve function not found in list of resolvers for - " + resolverPath)
		}
	}

	return resolve(args...)
}
