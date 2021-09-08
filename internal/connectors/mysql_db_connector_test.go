package connectors

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
)

func TestCreateBrand(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(ioutil.Discard)

	t.Run("error-create-brand", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectExec("INSERT INTO brands").WillReturnError(fmt.Errorf("Error DB"))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		brand := &BrandRecord{
			Name: "test brand",
		}

		_, err = mySQL.CreateBrand(context.Background(), brand)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectExec("INSERT INTO brands").WillReturnResult(sqlmock.NewResult(12, 1))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		brand := &BrandRecord{
			Name: "test brand",
		}

		_, err = mySQL.CreateBrand(context.Background(), brand)
		if err != nil {
			t.Error("error shouldnt be occurs")
			t.FailNow()
		}
	})
}

func TestGetBrandByID(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(ioutil.Discard)

	t.Run("error-exec-query-row-context", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM brands").WillReturnError(fmt.Errorf("Error DB"))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetBrandByID(context.Background(), 1)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("error-not-found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM brands").WillReturnError(sql.ErrNoRows)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetBrandByID(context.Background(), 1)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test")

		mock.ExpectQuery("SELECT (.+) FROM brands").WillReturnRows(rows)

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetBrandByID(context.Background(), 1)
		if err != nil {
			t.Error("error shouldnt be occurs")
			t.FailNow()
		}
	})
}

func TestCreateProduct(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(ioutil.Discard)

	t.Run("error-create-product", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectExec("INSERT INTO products").WillReturnError(fmt.Errorf("Error DB"))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		product := &ProductRecord{
			BrandID: 1,
			Name:    "test product",
			Qty:     1,
			Price:   1000,
		}

		_, err = mySQL.CreateProduct(context.Background(), product)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(12, 1))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		product := &ProductRecord{
			BrandID: 1,
			Name:    "test product",
			Qty:     1,
			Price:   1000,
		}

		_, err = mySQL.CreateProduct(context.Background(), product)
		if err != nil {
			t.Error("error shouldnt be occurs")
			t.FailNow()
		}
	})
}

func TestGetProductByID(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(ioutil.Discard)

	t.Run("error-exec-query-row-context", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM products").WillReturnError(fmt.Errorf("Error DB"))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetProductByID(context.Background(), 1)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("error-not-found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM products").WillReturnError(sql.ErrNoRows)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetProductByID(context.Background(), 1)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		rows := sqlmock.NewRows([]string{"id", "product_id", "name", "qty", "price"}).AddRow(1, 1, "name", 1, 1000)

		mock.ExpectQuery("SELECT (.+) FROM products").WillReturnRows(rows)

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetProductByID(context.Background(), 1)
		if err != nil {
			t.Error("error shouldnt be occurs")
			t.FailNow()
		}
	})
}

func TestGetProductByBrandID(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(ioutil.Discard)

	t.Run("error-exec-query-row-context", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM products").WillReturnError(fmt.Errorf("Error DB"))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetProductByBrandID(context.Background(), 1)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("error-not-found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM products").WillReturnError(sql.ErrNoRows)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetProductByBrandID(context.Background(), 1)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		rows := sqlmock.NewRows([]string{"id", "product_id", "name", "qty", "price"}).
			AddRow(1, 1, "name 1", 1, 1000).
			AddRow(2, 1, "name 2", 2, 1100).
			AddRow(3, 1, "name 3", 3, 1200)

		mock.ExpectQuery("SELECT (.+) FROM products").WillReturnRows(rows)

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetProductByBrandID(context.Background(), 1)
		if err != nil {
			t.Error("error shouldnt be occurs")
			t.FailNow()
		}
	})
}

func TestGetTransactionByTransactionID(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(ioutil.Discard)

	t.Run("error-exec-query-row-context", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM transactions").WillReturnError(fmt.Errorf("Error DB"))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetTransactionByTransactionID(context.Background(), 1)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("error-not-found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM transactions").WillReturnError(sql.ErrNoRows)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetTransactionByTransactionID(context.Background(), 1)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		rows := sqlmock.NewRows([]string{"id", "user_id", "date", "grand_total"}).
			AddRow(1, 1, time.Now(), 1000)

		mock.ExpectQuery("SELECT (.+) FROM transactions").WillReturnRows(rows)

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetTransactionByTransactionID(context.Background(), 1)
		if err != nil {
			t.Error("error shouldnt be occurs")
			t.FailNow()
		}
	})
}

func TestGetTransactionDetailByTransactionID(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(ioutil.Discard)

	t.Run("error-exec-query-row-context", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM transactions").WillReturnError(fmt.Errorf("Error DB"))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetTransactionByTransactionID(context.Background(), 1)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("error-not-found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM transactions").WillReturnError(sql.ErrNoRows)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetTransactionByTransactionID(context.Background(), 1)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		rows := sqlmock.NewRows([]string{"id", "user_id", "date", "grand_total"}).
			AddRow(1, 1, time.Now(), 1000)

		mock.ExpectQuery("SELECT (.+) FROM transactions").WillReturnRows(rows)

		rows = sqlmock.NewRows([]string{"transaction_id", "product_id", "price", "qty", "sub_total"}).
			AddRow(1, 1, 100, 1, 1000).
			AddRow(1, 2, 100, 1, 1000).
			AddRow(1, 2, 100, 1, 1000)

		mock.ExpectQuery("SELECT (.+) FROM transaction_detail").WillReturnRows(rows)

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetTransactionByTransactionID(context.Background(), 1)
		if err != nil {
			t.Error("error shouldnt be occurs")
			t.FailNow()
		}
	})
}

func TestCreateTransaction(t *testing.T) {
	t.Run("error-begin-trx", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectBegin().WillReturnError(fmt.Errorf("Transaction Error"))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}
		_, err = mySQL.CreateTransaction(context.Background(), &TransactionRecord{})
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("error-create-transaction-record", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO transactions").WillReturnError(fmt.Errorf("Create Transaction Error"))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}
		_, err = mySQL.CreateTransaction(context.Background(), &TransactionRecord{})
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("error-last-inserted-id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO transactions").WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("error LastInsertedID")))
		mock.ExpectRollback()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.CreateTransaction(context.Background(), &TransactionRecord{})
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO transactions").WillReturnResult(sqlmock.NewResult(12, 1))

		rows := sqlmock.NewRows([]string{"id", "product_id", "name", "qty", "price"}).AddRow(1, 1, "name", 1, 1000)
		mock.ExpectQuery("SELECT (.+) FROM products").WillReturnRows(rows)

		mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(12, 1))

		mock.ExpectExec("INSERT INTO transaction_detail").WillReturnResult(sqlmock.NewResult(12, 1))

		mock.ExpectExec("UPDATE transactions").WillReturnResult(sqlmock.NewResult(12, 1))

		mock.ExpectCommit()

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		var det []*TransactionDetailRecord

		det = append(det, &TransactionDetailRecord{
			ProductID: 1,
			Qty:       1,
		})

		rec := &TransactionRecord{
			UserID:            1,
			Date:              time.Now(),
			TransactionDetail: det,
		}

		_, err = mySQL.CreateTransaction(context.Background(), rec)
		if err != nil {
			t.Error("error shouldnt be occurs")
			t.FailNow()
		}
	})

}

func TestGetUserByID(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(ioutil.Discard)

	t.Run("error-exec-query-row-context", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM users").WillReturnError(fmt.Errorf("Error DB"))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetUserByID(context.Background(), 1)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("error-not-found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM users").WillReturnError(sql.ErrNoRows)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetUserByID(context.Background(), 1)
		if err == nil {
			t.Error("error should be occurs")
			t.FailNow()
		}
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		rows := sqlmock.NewRows([]string{"id", "email", "name", "address"}).AddRow(1, "donny@arieffian.com", "donny", "surabaya")

		mock.ExpectQuery("SELECT (.+) FROM users").WillReturnRows(rows)

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()
		// inject sqlmock.DB into MySQLDB
		mySQL := MySQLDB{
			instance: db,
		}

		_, err = mySQL.GetUserByID(context.Background(), 1)
		if err != nil {
			t.Error("error shouldnt be occurs")
			t.FailNow()
		}
	})
}
