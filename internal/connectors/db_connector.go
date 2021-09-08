package connectors

import (
	"context"
	"time"

	//Anonymous import for mysql initialization
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithField("module", "db_connector")
)

// BrandRecord an entity representative of brands table
type BrandRecord struct {
	ID   int
	Name string
}

// UserRecord an entity representative of users table
type UserRecord struct {
	ID      int
	Name    string
	Email   string
	Address string
}

// ProductRecord an entity representative of products table
type ProductRecord struct {
	ID      int
	BrandID int
	SKU     string
	Name    string
	Qty     int
	Price   int
}

// TransactionRecord an entity representative of transactions table
type TransactionRecord struct {
	ID         int
	UserID     int
	Date       time.Time
	GrandTotal int
	Discount   int
	Reason     string

	TransactionDetail []*TransactionDetailRecord
}

// TransactionDetailRecord an entity representative of transaction_detail table
type TransactionDetailRecord struct {
	TransactionID int
	ProductID     int
	SKU           string
	Qty           int
	SubTotal      int
}

type UserRepository interface {
	// GetUserByID retrieves an UserRecord from database where the user id is specified.
	GetUserByID(ctx context.Context, userID int) (*UserRecord, error)
}

type BrandRepository interface {
	// GetBrandByID retrieves an BrandRecord from database where the brand id is specified.
	GetBrandByID(ctx context.Context, brandID int) (*BrandRecord, error)

	// CreateBrand insert an entity record of brand into database.
	CreateBrand(ctx context.Context, rec *BrandRecord) (string, error)
}

type ProductRepository interface {
	// CreateProduct insert an entity record of product into database.
	CreateProduct(ctx context.Context, rec *ProductRecord) (string, error)

	// GetProductByID retrieves an ProductRecord from database where the product id is specified.
	GetProductByID(ctx context.Context, productID int) (*ProductRecord, error)

	// GetProductByBrandID retrieves an array of ProductRecord from database where the brand id is specified.
	GetProductByBrandID(ctx context.Context, brandID int) ([]*ProductRecord, error)
}

type TransactionRepository interface {
	// CreateTransaction insert an entity record of transaction into database.
	CreateTransaction(ctx context.Context, rec *TransactionRecord) (string, error)

	// GetTransactionByTransactionID retrieves the detail of a transaction from database where the transaction id is specified.
	GetTransactionByTransactionID(ctx context.Context, transactionID int) (*TransactionRecord, error)

	// GetTransactionDetailByTransactionID retrieves the detail of a transaction from database where the transaction id is specified.
	GetTransactionDetailByTransactionID(ctx context.Context, transactionID int) ([]*TransactionDetailRecord, error)
}
