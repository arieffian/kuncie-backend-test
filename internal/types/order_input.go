package types

import "github.com/graphql-go/graphql"

type OrderInput struct {
	UserID int         `json:"user_id" mapstructure:"user_id"`
	Items  []ItemInput `json:"items" mapstructure:"items"`
}

type ItemInput struct {
	ProductID int    `json:"product_id" mapstructure:"product_id"`
	SKU       string `json:"sku" mapstructure:"sku"`
	Qty       int    `json:"qty" mapstructure:"qty"`
}

var OrderInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "orderInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"user_id": &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"items": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(ItemInputType),
		},
	},
})

var ItemInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "itemInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"product_id": &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"sku":        &graphql.InputObjectFieldConfig{Type: graphql.String},
		"qty":        &graphql.InputObjectFieldConfig{Type: graphql.Int},
	},
})
