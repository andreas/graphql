package graphql_test

import (
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/testutil"
)

func TestValidate_KnownTypeNames_KnownTypeNamesAreValid(t *testing.T) {
	testutil.ExpectPassesRule(t, graphql.KnownTypeNamesRule, `
      query Foo($var: String, $required: [String!]!) {
        user(id: 4) {
          pets { ... on Pet { name }, ...PetFields }
        }
      }
      fragment PetFields on Pet {
        name
      }
    `)
}
func TestValidate_KnownTypeNames_UnknownTypeNamesAreInValid(t *testing.T) {
	testutil.ExpectFailsRule(t, graphql.KnownTypeNamesRule, `
      query Foo($var: JumbledUpLetters) {
        user(id: 4) {
          name
          pets { ... on Badger { name }, ...PetFields }
        }
      }
      fragment PetFields on Peettt {
        name
      }
    `, []gqlerrors.FormattedError{
		testutil.RuleError(`Unknown type "JumbledUpLetters".`, 2, 23),
		testutil.RuleError(`Unknown type "Badger".`, 5, 25),
		testutil.RuleError(`Unknown type "Peettt".`, 8, 29),
	})
}
