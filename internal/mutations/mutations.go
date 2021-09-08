package mutations

import (
	"github.com/arieffian/kuncie-backend-test/internal/connectors"
	"github.com/graphql-go/graphql"
)

var (
	TransactionRepo connectors.TransactionRepository
	UserRepo        connectors.UserRepository
	ProductRepo     connectors.ProductRepository
)

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"createOrder": CreateOrder(),
	}
}
