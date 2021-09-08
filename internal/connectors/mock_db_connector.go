package connectors

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockDb struct{}

//MockDBType ...
type MockDBType struct {
	mock.Mock
}

// GetBrandByID retrieves an BrandRecord from database where the brand id is specified.
func (m *MockDBType) GetBrandByID(ctx context.Context, brandID int) (*BrandRecord, error) {
	args := m.Called(ctx, brandID)
	return args.Get(0).(*BrandRecord), args.Error(1)
}

// CreateBrand insert an entity record of brand into database.
func (m *MockDBType) CreateBrand(ctx context.Context, rec *BrandRecord) (string, error) {
	args := m.Called(ctx, rec)
	return args.String(0), args.Error(1)
}

// CreateProduct insert an entity record of product into database.
func (m *MockDBType) CreateProduct(ctx context.Context, rec *ProductRecord) (string, error) {
	args := m.Called(ctx, rec)
	return args.String(0), args.Error(1)
}

// GetProductByID retrieves an ProductRecord from database where the product id is specified.
func (m *MockDBType) GetProductByID(ctx context.Context, productID int) (*ProductRecord, error) {
	args := m.Called(ctx, productID)
	return args.Get(0).(*ProductRecord), args.Error(1)
}

// GetProductByBrandID retrieves an array of ProductRecord from database where the brand id is specified.
func (m *MockDBType) GetProductByBrandID(ctx context.Context, brandID int) ([]*ProductRecord, error) {
	pList := make([]*ProductRecord, 0)
	args := m.Called(ctx, brandID)
	return pList, args.Error(1)
}

// GetTransactionByTransactionID retrieves the detail of a transaction from database where the transaction id is specified.
func (m *MockDBType) GetTransactionByTransactionID(ctx context.Context, transactionID int) (*TransactionRecord, error) {
	args := m.Called(ctx, transactionID)
	return args.Get(0).(*TransactionRecord), args.Error(1)
}

// GetTransactionByTransactionID retrieves the detail of a transaction from database where the transaction id is specified.
func (m *MockDBType) GetTransactionDetailByTransactionID(ctx context.Context, transactionID int) ([]*TransactionDetailRecord, error) {
	args := m.Called(ctx, transactionID)
	return args.Get(0).([]*TransactionDetailRecord), args.Error(1)
}

// CreateTransaction insert an entity record of transaction into database.
func (m *MockDBType) CreateTransaction(ctx context.Context, rec *TransactionRecord) (string, error) {
	args := m.Called(ctx, rec)
	return args.String(0), args.Error(1)
}

// GetUserByID retrieves an UserRecord from database where the user id is specified.
func (m *MockDBType) GetUserByID(ctx context.Context, userID int) (*UserRecord, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*UserRecord), args.Error(1)
}
