package repositories

import (
	"database/sql"
	"fmt"
	models "kasir-api/model"
	"strings"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}
func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	var (
		res *models.Transaction
	)
	//db
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Inisialisasi subtotal -> jumlah total transaksi keseluruhan.
	totalAmount := 0
	// Inisialisasi modeling transactionDetails -> nanti akan diinsert ke database.
	details := make([]models.TransactionDetails, 0)
	// Loop setiap item:
	for _, item := range items {
		var productName string
		var productID, price, stock int
		// Get product untuk mendapatkan harga (pricing).
		err := tx.QueryRow("SELECT id, name, price, stock FROM products WHERE id=$1", item.ProductID).Scan(&productID, &productName, &price, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		// Hitung current total = quantity × pricing.
		// diTambahkan ke dalam subtotal.
		subtotal := item.Quantity * price
		totalAmount += subtotal

		// kurangi jumlah stock
		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, productID)
		if err != nil {
			return nil, err
		}

		// itemnya dimasukan ke transactionDetails.
		details = append(details, models.TransactionDetails{
			ProductID:   productID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	// Insert transaction.
	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING ID", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	// Insert transactionDetails.
	if len(details) == 0 {
		return nil, nil // Atau handle sesuai logika bisnis jika tidak ada detail
	}

	valueStrings := make([]string, 0, len(details))
	valueArgs := make([]interface{}, 0, len(details)*4)
	i := 1 // Counter untuk placeholder $1, $2, dst.

	for _, d := range details {
		valueStrings = append(
			valueStrings,
			fmt.Sprintf("($%d, $%d, $%d, $%d)", i, i+1, i+2, i+3),
		)

		valueArgs = append(valueArgs, transactionID)
		valueArgs = append(valueArgs, d.ProductID)
		valueArgs = append(valueArgs, d.Quantity)
		valueArgs = append(valueArgs, d.Subtotal)

		i += 4
	}

	stmt := fmt.Sprintf(
		"INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES %s",
		strings.Join(valueStrings, ","),
	)

	_, err = tx.Exec(stmt, valueArgs...)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	res = &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}

	return res, nil
}
