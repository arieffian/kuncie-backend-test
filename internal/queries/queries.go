package queries

import (
	"github.com/arieffian/kuncie-backend-test/internal/connectors"
	"github.com/graphql-go/graphql"
)

var (
	TransactionRepo connectors.TransactionRepository
)

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"getOrder": GetOrderQuery(),
	}
}
