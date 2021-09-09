package connectors

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/arieffian/kuncie-backend-test/internal/config"
)

var (
	mysqlLog        = log.WithField("file", "mysql_db_connector.go")
	mySQLDbInstance *MySQLDB
)

// GetMySQLDBInstance initializes the MySQL.DB instance
func GetMySQLDBInstance() *MySQLDB {
	if mySQLDbInstance == nil {
		host := config.Get("db.host")
		port := config.GetInt("db.port")
		user := config.Get("db.user")
		password := config.Get("db.password")
		database := config.Get("db.database")
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", user, password, host, port, database))
		if err != nil {
			mysqlLog.WithField("func", "GetMySQLDBInstance").Fatalf("sql.Open got %s", err.Error())
		}

		mySQLDbInstance = &MySQLDB{
			instance: db,
		}
	}
	return mySQLDbInstance
}

// MySQLDB db instance
type MySQLDB struct {
	instance *sql.DB
}

// GetBrandByID retrieves an BrandRecord from database where the brand id is specified.
func (db *MySQLDB) GetBrandByID(ctx context.Context, brandID int) (*BrandRecord, error) {
	fLog := mysqlLog.WithField("func", "GetBrandByID")
	brand := &BrandRecord{}

	row := db.instance.QueryRowContext(ctx, "SELECT id, name FROM brands WHERE id = ?", brandID)
	err := row.Scan(&brand.ID, &brand.Name)
	if err != nil {
		fLog.Errorf("row.Scan got %s", err.Error())
		return nil, err
	}

	return brand, nil
}

// CreateBrand insert an entity record of brand into database.
func (db *MySQLDB) CreateBrand(ctx context.Context, rec *BrandRecord) (string, error) {
	fLog := mysqlLog.WithField("func", "CreateBrand")

	_, err := db.instance.ExecContext(ctx, "INSERT INTO brands(name) VALUES(?)", rec.Name)
	if err != nil {
		fLog.Errorf("db.instance.ExecContext got %s", err.Error())
		return "", err
	}

	return "brand created successfully", nil
}

// CreateProduct insert an entity record of product into database.
func (db *MySQLDB) CreateProduct(ctx context.Context, rec *ProductRecord) (string, error) {
	fLog := mysqlLog.WithField("func", "CreateProduct")

	_, err := db.instance.ExecContext(ctx, "INSERT INTO products(brand_id, name, qty, price, sku) VALUES(?,?,?,?,?)", rec.BrandID, rec.Name, rec.Qty, rec.Price, rec.SKU)
	if err != nil {
		fLog.Errorf("db.instance.ExecContext got %s", err.Error())
		return "", err
	}

	return "product created successfully", nil
}

// GetProductByID retrieves an ProductRecord from database where the product id is specified.
func (db *MySQLDB) GetProductByID(ctx context.Context, productID int) (*ProductRecord, error) {
	fLog := mysqlLog.WithField("func", "GetProductByID")
	product := &ProductRecord{}

	row := db.instance.QueryRowContext(ctx, "SELECT id, brand_id, name, price, qty FROM products WHERE id = ?", productID)
	err := row.Scan(&product.ID, &product.BrandID, &product.Name, &product.Price, &product.Qty)
	if err != nil {
		fLog.Errorf("row.Scan got %s", err.Error())
		return nil, err
	}

	return product, nil
}

// GetProductByBrandID retrieves an array of ProductRecord from database where the brand id is specified.
func (db *MySQLDB) GetProductByBrandID(ctx context.Context, brandID int) ([]*ProductRecord, error) {
	fLog := mysqlLog.WithField("func", "GetProductByBrandID")

	q := fmt.Sprintf("SELECT id, brand_id, name, price, qty FROM products WHERE brand_id = %v", brandID)
	rows, err := db.instance.QueryContext(ctx, q)
	if err != nil {
		fLog.Errorf("db.instance.QueryContext got %s", err.Error())
		return nil, err
	}
	productList := make([]*ProductRecord, 0)
	for rows.Next() {
		product := &ProductRecord{}
		err := rows.Scan(&product.ID, &product.BrandID, &product.Name, &product.Price, &product.Qty)
		if err != nil {
			fLog.Errorf("rows.Scan got %s", err.Error())
		} else {
			productList = append(productList, product)
		}
	}
	return productList, nil
}

// GetTransactionByTransactionID retrieves the detail of a transaction from database where the transaction id is specified.
func (db *MySQLDB) GetTransactionByTransactionID(ctx context.Context, transactionID int) (*TransactionRecord, error) {
	fLog := mysqlLog.WithField("func", "GetTransactionByTransactionID")
	transaction := &TransactionRecord{}

	row := db.instance.QueryRowContext(ctx, "SELECT id, user_id, date, grand_total, discount, reason FROM transactions WHERE id = ?", transactionID)
	err := row.Scan(&transaction.ID, &transaction.UserID, &transaction.Date, &transaction.GrandTotal, &transaction.Discount, &transaction.Reason)
	if err != nil {
		fLog.Errorf("row.Scan got %s", err.Error())
		return nil, err
	}

	q := fmt.Sprintf("SELECT td.transaction_id, td.product_id, td.qty, td.sub_total, p.sku FROM transaction_detail td INNER JOIN products p ON td.product_id = p.id WHERE transaction_id = %v", transactionID)
	rows, err := db.instance.QueryContext(ctx, q)
	if err != nil {
		fLog.Errorf("db.instance.QueryContext got %s", err.Error())
		return nil, err
	}

	tDetail := make([]*TransactionDetailRecord, 0)
	for rows.Next() {
		tD := &TransactionDetailRecord{}
		err := rows.Scan(&tD.TransactionID, &tD.ProductID, &tD.Qty, &tD.SubTotal, &tD.SKU)
		if err != nil {
			fLog.Errorf("rows.Scan got %s", err.Error())
		} else {
			tDetail = append(tDetail, tD)
		}
	}

	transaction.TransactionDetail = tDetail

	return transaction, nil
}

// GetTransactionDetailByTransactionID retrieves the detail of a transaction from database where the transaction id is specified.
func (db *MySQLDB) GetTransactionDetailByTransactionID(ctx context.Context, transactionID int) ([]*TransactionDetailRecord, error) {
	fLog := mysqlLog.WithField("func", "GetTransactionByTransactionID")
	transaction := &TransactionRecord{}

	row := db.instance.QueryRowContext(ctx, "SELECT id FROM transactions WHERE id = ?", transactionID)
	err := row.Scan(&transaction.ID)
	if err != nil {
		fLog.Errorf("row.Scan got %s", err.Error())
		return nil, err
	}

	q := fmt.Sprintf("SELECT td.transaction_id, td.product_id, td.qty, td.sub_total, p.sku FROM transaction_detail td INNER JOIN products p ON td.product_id = p.id WHERE transaction_id = %v", transactionID)
	rows, err := db.instance.QueryContext(ctx, q)
	if err != nil {
		fLog.Errorf("db.instance.QueryContext got %s", err.Error())
		return nil, err
	}

	tDetail := make([]*TransactionDetailRecord, 0)
	for rows.Next() {
		tD := &TransactionDetailRecord{}
		err := rows.Scan(&tD.TransactionID, &tD.ProductID, &tD.Qty, &tD.SubTotal, tD.SKU)
		if err != nil {
			fLog.Errorf("rows.Scan got %s", err.Error())
		} else {
			tDetail = append(tDetail, tD)
		}
	}

	return tDetail, nil
}

// CreateTransaction insert an entity record of transaction into database.
func (db *MySQLDB) CreateTransaction(ctx context.Context, rec *TransactionRecord) (int, error) {
	fLog := mysqlLog.WithField("func", "CreateTransaction")

	// start db transaction
	tx, err := db.instance.BeginTx(ctx, nil)
	if err != nil {
		fLog.Errorf("db.instance.BeginTx got %s", err.Error())
		return 0, err
	}

	// create transaction record
	trans, err := tx.ExecContext(ctx, "INSERT INTO transactions(user_id, date, grand_total, discount, reason) VALUES(?,?,?,?,?)", rec.UserID, rec.Date, 0, rec.Discount, rec.Reason)
	if err != nil {
		fLog.Errorf("db.tx.ExecContext got %s", err.Error())
		errRollback := tx.Rollback()
		if errRollback != nil {
			fLog.Errorf("error rollback, got %s", err.Error())
			return 0, errRollback
		}
		return 0, err
	}

	tID, err := trans.LastInsertId()
	if err != nil {
		fLog.Errorf("db.tx.ExecContext got %s", err.Error())
		errRollback := tx.Rollback()
		if errRollback != nil {
			fLog.Errorf("error rollback, got %s", err.Error())
			return 0, errRollback
		}
		return 0, err
	}

	grandTotal := 0

	//loop tx detail
	for i := 0; i < len(rec.TransactionDetail); i++ {
		detail := rec.TransactionDetail[i]

		//get product stock from products table
		p, err := db.GetProductByID(ctx, detail.ProductID)
		if err != nil {
			fLog.Errorf("db.tx.ExecContext got %s", err.Error())
			errRollback := tx.Rollback()
			if errRollback != nil {
				fLog.Errorf("error rollback, got %s", err.Error())
				return 0, errRollback
			}
			return 0, err
		}

		//check qty
		if p.Qty-detail.Qty < 0 {
			fLog.Errorf("product qty is not enough")
			errRollback := tx.Rollback()
			if errRollback != nil {
				fLog.Errorf("error rollback, got %s", err.Error())
				return 0, errRollback
			}
			return 0, fmt.Errorf("product qty is not enough")
		}

		qty := p.Qty - detail.Qty
		subTotal := p.Price * detail.Qty
		grandTotal = grandTotal + subTotal

		//update qty from products table
		_, err = tx.ExecContext(ctx, "UPDATE products SET qty=? WHERE id=?", qty, detail.ProductID)
		if err != nil {
			fLog.Errorf("db.tx.ExecContext got %s", err.Error())
			errRollback := tx.Rollback()
			if errRollback != nil {
				fLog.Errorf("error rollback, got %s", err.Error())
				return 0, errRollback
			}
			return 0, err
		}

		//insert transaction detail
		_, err = tx.ExecContext(ctx, "INSERT INTO transaction_detail(transaction_id, product_id, price, qty, sub_total) VALUES(?,?,?,?,?)", tID, detail.ProductID, p.Price, detail.Qty, subTotal)
		if err != nil {
			fLog.Errorf("db.tx.ExecContext got %s", err.Error())
			errRollback := tx.Rollback()
			if errRollback != nil {
				fLog.Errorf("error rollback, got %s", err.Error())
				return 0, errRollback
			}
			return 0, err
		}
	}

	grandTotal = grandTotal - rec.Discount

	// update transaction grand total
	_, err = tx.ExecContext(ctx, "UPDATE transactions SET grand_total=? WHERE id=?", grandTotal, tID)
	if err != nil {
		fLog.Errorf("db.tx.ExecContext got %s", err.Error())
		errRollback := tx.Rollback()
		if errRollback != nil {
			fLog.Errorf("error rollback, got %s", err.Error())
			return 0, errRollback
		}
		return 0, err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		fLog.Errorf("db.tx.ExecContext got %s", err.Error())
		errRollback := tx.Rollback()
		if errRollback != nil {
			fLog.Errorf("error rollback, got %s", err.Error())
			return 0, errRollback
		}
		return 0, err
	}

	return int(tID), nil
}

// GetUserByID retrieves an UserRecord from database where the user id is specified.
func (db *MySQLDB) GetUserByID(ctx context.Context, userID int) (*UserRecord, error) {
	fLog := mysqlLog.WithField("func", "GetUserByID")
	user := &UserRecord{}

	row := db.instance.QueryRowContext(ctx, "SELECT id, name, email, address FROM users WHERE id = ?", userID)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Address)
	if err != nil {
		fLog.Errorf("row.Scan got %s", err.Error())
		return nil, err
	}

	return user, nil
}
