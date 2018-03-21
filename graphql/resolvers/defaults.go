package lmresolvers

import (
	"log"

	"github.com/graphql-go/graphql"
)

// DefaultReadResolver - Default read resolver
func DefaultReadResolver(args ...interface{}) graphql.FieldResolveFn {
	log.Println("Default resolver working")

	return func(params graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	}
}

// DefaultListResolver - Default list resolver
// func DefaultListResolver(action map[string]string, params graphql.ResolveParams) (interface{}, error) {
// 	return nil, nil
// }
