package mutations

import (
	"time"

	"github.com/arieffian/kuncie-backend-test/internal/connectors"
	"github.com/arieffian/kuncie-backend-test/internal/types"
	"github.com/graphql-go/graphql"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/mitchellh/mapstructure"
)

type Fact struct {
	Product  []*FactProduct
	Discount int64
	Reason   string
}

type FactProduct struct {
	ProductID int64
	Price     int64
	Qty       int64
}

func (f *Fact) IsExist(productID int64) bool {
	exist := false
	for _, product := range f.Product {
		if product.ProductID == productID {
			exist = true
			break
		}
	}
	return exist
}

func (f *Fact) GetProductIndex(productID int64) (key int64) {
	key = -1
	for k, product := range f.Product {
		if product.ProductID == productID {
			key = int64(k)
			break
		}
	}
	return key
}

func (f *Fact) GetQty(productID int64) (qty int64) {
	qty = -1
	for _, product := range f.Product {
		if product.ProductID == productID {
			qty = product.Qty
			break
		}
	}
	return qty
}

func CreateOrder() *graphql.Field {
	return &graphql.Field{
		Type: types.OrderType,
		Args: graphql.FieldConfigArgument{
			"order": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(types.OrderInputType),
			},
		},
		Resolve: createTransaction,
	}
}

func calculateDiscount(fact *Fact) {

	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("Ft", fact)
	if err != nil {
		panic(err)
	}

	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	bs := pkg.NewFileResource("rules/discounts.grl")
	err = ruleBuilder.BuildRuleFromResource("DiscountRules", "0.0.1", bs)
	if err != nil {
		panic(err)
	}

	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("DiscountRules", "0.0.1")

	engine := engine.NewGruleEngine()
	err = engine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		panic(err)
	}
}

func createTransaction(p graphql.ResolveParams) (interface{}, error) {

	var order types.OrderInput

	mapstructure.Decode(p.Args["order"], &order)

	// validate user id exists
	_, err := UserRepo.GetUserByID(p.Context, order.UserID)
	if err != nil {
		return graphql.Interface{}, err
	}

	trans := &connectors.TransactionRecord{
		UserID: order.UserID,
		Date:   time.Now(),
	}

	// calculate discount
	var f Fact
	var d []*FactProduct
	var td []*connectors.TransactionDetailRecord
	for i := 0; i < len(order.Items); i++ {
		// validate product id exists
		product, err := ProductRepo.GetProductByID(p.Context, order.Items[i].ProductID)
		if err != nil {
			return graphql.Interface{}, err
		}

		td = append(td, &connectors.TransactionDetailRecord{
			ProductID: product.ID,
			Qty:       order.Items[i].Qty,
		})

		d = append(d, &FactProduct{
			ProductID: int64(product.ID),
			Price:     int64(product.Price),
			Qty:       int64(order.Items[i].Qty),
		})

	}
	f.Product = d
	calculateDiscount(&f)

	trans.Discount = int(f.Discount)
	trans.Reason = f.Reason
	trans.TransactionDetail = td

	tID, err := TransactionRepo.CreateTransaction(p.Context, trans)
	if err != nil {
		return graphql.Interface{}, err
	}

	trans, err = TransactionRepo.GetTransactionByTransactionID(p.Context, tID)
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
		Discount:    trans.Discount,
		Reason:      trans.Reason,
	}

	return o, nil
}
