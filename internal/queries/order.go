package queries

import (
	"github.com/arieffian/kuncie-backend-test/internal/types"
	"github.com/graphql-go/graphql"
)

func GetOrderQuery() *graphql.Field {
	return &graphql.Field{
		Type: types.OrderType,
		Args: graphql.FieldConfigArgument{
			"orderID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: getTransaction,
	}
}

func getTransaction(p graphql.ResolveParams) (interface{}, error) {
	orderID := p.Args["orderID"].(int)

	trans, err := TransactionRepo.GetTransactionByTransactionID(p.Context, orderID)
	if err != nil {
		return graphql.Interface{}, err
	}

	var oDetail []types.OrderDetail
	for i := 0; i < len(trans.TransactionDetail); i++ {
		detail := trans.TransactionDetail[i]

		d := types.OrderDetail{
			TransactionID: detail.TransactionID,
			ProductID:     detail.ProductID,
			SKU:           detail.SKU,
			Qty:           detail.Qty,
			SubTotal:      detail.SubTotal,
		}

		oDetail = append(oDetail, d)
	}

	o := types.Order{
		ID:          trans.ID,
		UserID:      trans.UserID,
		GrandTotal:  trans.GrandTotal,
		OrderDetail: oDetail,
	}

	return o, nil
}
