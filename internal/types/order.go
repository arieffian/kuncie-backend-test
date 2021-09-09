package types

import (
	"github.com/graphql-go/graphql"
)

type Order struct {
	ID          int           `json:"id"`
	UserID      int           `json:"user_id"`
	GrandTotal  int           `json:"grand_total"`
	Discount    int           `json:"discount"`
	Reason      string        `json:"reason"`
	OrderDetail []OrderDetail `json:"detail_order"`
}

type OrderDetail struct {
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	SKU           string `json:"sku"`
	Qty           int    `json:"qty"`
	SubTotal      int    `json:"sub_total"`
}

var DetailType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Detail",
	Fields: graphql.Fields{
		"transaction_id": &graphql.Field{Type: graphql.Int},
		"product_id":     &graphql.Field{Type: graphql.Int},
		"sku":            &graphql.Field{Type: graphql.String},
		"qty":            &graphql.Field{Type: graphql.Int},
		"sub_total":      &graphql.Field{Type: graphql.Int},
	},
})

var OrderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Order",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.Int},
		"user_id":     &graphql.Field{Type: graphql.Int},
		"grand_total": &graphql.Field{Type: graphql.Int},
		"discount":    &graphql.Field{Type: graphql.Int},
		"reason":      &graphql.Field{Type: graphql.String},
		"detail_order": &graphql.Field{
			Type: graphql.NewList(DetailType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				detail := p.Source.(Order).OrderDetail
				return detail, nil
			},
		},
	},
})
